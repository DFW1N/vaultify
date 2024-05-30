// ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Easily push, pull and encrypt tofu and terraform statefiles from Vault. #
// ########################################################################################

package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Add a case switch statement depending on, the default secret storage type.

func Delete() {
	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
		return
	}

	vaultClient, initStat := initVaultClientWithStatus()
	if !initStat {
		fmt.Println("❌ Error: Vault is not initialized!")
		return
	}

	settings, err := readSettings()
	if err != nil {
		fmt.Println("❌ Error reading settings:", err)
		return
	}

	engineName := settings.Settings.DefaultEngineName
	dataPath := "vaultify"

	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		fmt.Println("❌ Error getting current Terraform workspace:", err)
		return
	}
	workspaceName = strings.TrimSpace(workspaceName)

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Error getting current working directory:", err)
		return
	}

	workingDirName := filepath.Base(workingDir)
	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workingDirName, workspaceName, "terraform.tfstate")

	metadataValue, err := vaultClient.KVv2(engineName).GetMetadata(context.Background(), secretPath)
	if err == nil {
		if metadataValue == nil {
			fmt.Println("❌ Error retrieving metadata:", err)
			return
		}
		latestVersion := fmt.Sprintf("%d", metadataValue.CurrentVersion)
		if versionData, exists := metadataValue.Versions[latestVersion]; exists && versionData.DeletionTime.String() != "0001-01-01 00:00:00 +0000 UTC" {
			fmt.Println("✅ The latest version of the secret has already been deleted.")
			return
		}
	}

	fmt.Printf("Are you sure you want to delete the secret at '%s'? [y/N]: ", secretPath)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if strings.ToLower(input) != "y" && strings.ToLower(input) != "yes" {
		fmt.Println("Deletion canceled.")
		return
	}

	err = vaultClient.KVv2(engineName+"/").Delete(context.Background(), secretPath)
	if err != nil {
		fmt.Println("❌ Error deleting secret from Vault:", err)
		return
	}

	fmt.Println("✅ Secret successfully deleted from Vault.")
}
