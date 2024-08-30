// ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Inject a value into your vault, by running a simple command.            #
// ########################################################################################

package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Inject(args []string) {
	injectCmd := flag.NewFlagSet("inject", flag.ExitOnError)
	path := injectCmd.String("path", "", "Vault path to store the secret")
	key := injectCmd.String("key", "value", "Key for the secret")
	err := injectCmd.Parse(args)
	if err != nil {
		fmt.Println("❌ Error parsing inject flags:", err)
		return
	}

	if injectCmd.NArg() < 1 {
		fmt.Println("Usage: vaultify inject [-path <vault_path>] [-key <secret_key>] <secret_value>")
		return
	}

	secretValue := injectCmd.Arg(0)

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

	var secretPath string
	if *path != "" {
		secretPath = *path
	} else {
		workspaceName, err := getCurrentWorkspace()
		if err != nil {
			fmt.Println("❌ Error getting current Terraform workspace:", err)
			return
		}

		workingDir, err := os.Getwd()
		if err != nil {
			fmt.Println("❌ Error getting current working directory:", err)
			return
		}

		workingDirName := filepath.Base(workingDir)
		secretPath = fmt.Sprintf("%s/%s/%s", dataPath, workingDirName, workspaceName)
	}

	// Ensure the path exists
	err = ensureKVPathExists(vaultClient, engineName, secretPath)
	if err != nil {
		fmt.Println("❌ Error: Unable to perform operation", err)
		return
	}

	// Prepare the secret data
	secretData := map[string]interface{}{
		"data": map[string]interface{}{
			*key: secretValue,
		},
	}

	// Write the secret to Vault
	fullPath := engineName + "/data/" + strings.TrimPrefix(secretPath, "/")
	_, err = vaultClient.Logical().Write(fullPath, secretData)
	if err != nil {
		fmt.Println("❌ Error injecting secret to Vault:", err)
		return
	}

	fmt.Printf("✅ Secret injected to HashiCorp Vault under: \033[33m%s\033[0m\n", fullPath)
	fmt.Printf("💠 Key: \033[33m%s\033[0m\n", *key)
}
