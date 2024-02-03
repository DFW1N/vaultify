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
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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

	settings, err := readSettings()
	if err != nil {
		fmt.Println("❌ Error reading settings:", err)
		return
	}

	curlCommand := "curl"
	vaultURL := os.Getenv("VAULT_ADDR")
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

	metadataCmd := exec.Command(
		curlCommand,
		"--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"),
		"--request", "GET",
		"--silent",
		vaultURL+"/v1/"+engineName+"/metadata/"+secretPath,
	)

	metadataOutput, err := metadataCmd.Output()
	if err != nil {
		fmt.Println("❌ Error retrieving metadata:", err)
		return
	}

	var metadata struct {
		Data struct {
			CurrentVersion int `json:"current_version"`
			Versions       map[string]struct {
				DeletionTime string `json:"deletion_time"`
			} `json:"versions"`
		} `json:"data"`
	}

	if err := json.Unmarshal(metadataOutput, &metadata); err != nil {
		fmt.Println("❌ Error parsing metadata JSON:", err)
		return
	}

	latestVersion := fmt.Sprintf("%d", metadata.Data.CurrentVersion)
	if versionData, exists := metadata.Data.Versions[latestVersion]; exists && versionData.DeletionTime != "" {
		fmt.Println("✅ The latest version of the secret has already been deleted.")
		return
	}

	fmt.Printf("Are you sure you want to delete the secret at '%s'? [y/N]: ", secretPath)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if strings.ToLower(input) != "y" && strings.ToLower(input) != "yes" {
		fmt.Println("Deletion canceled.")
		return
	}

	deleteCmd := exec.Command(
		curlCommand,
		"--silent", "--show-error",
		"--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"),
		"--request", "DELETE",
		vaultURL+"/v1/"+engineName+"/data/"+secretPath,
	)

	_, err = deleteCmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Error deleting secret from Vault:", err)
		return
	}

	fmt.Println("✅ Secret successfully deleted from Vault.")
}
