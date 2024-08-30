// ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Get your secret from vault easily with a single command.                #
// ########################################################################################

package cmd

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Get(args []string) {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	path := getCmd.String("path", "", "Vault path to retrieve the secret from")
	key := getCmd.String("key", "", "Key of the secret to retrieve (optional)")
	err := getCmd.Parse(args)
	if err != nil {
		fmt.Println("❌ Error parsing get flags:", err)
		return
	}

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

	// Read the secret from Vault
	fullPath := engineName + "/data/" + strings.TrimPrefix(secretPath, "/")
	secret, err := vaultClient.Logical().Read(fullPath)
	if err != nil {
		fmt.Println("❌ Error reading secret from Vault:", err)
		return
	}

	if secret == nil || secret.Data == nil {
		fmt.Printf("❌ No secret found at path: \033[33m%s\033[0m\n", fullPath)
		return
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		fmt.Println("❌ Error: Unexpected secret data format")
		return
	}

	if *key != "" {
		// If a key is specified, print only that key's value
		value, exists := data[*key]
		if !exists {
			fmt.Printf("❌ No value found for key \033[33m%s\033[0m at path \033[33m%s\033[0m\n", *key, fullPath)
			return
		}
		fmt.Printf("%v\n", value)
	} else {
		// If no key is specified, print all key-value pairs
		fmt.Printf("Secrets at path \033[33m%s\033[0m:\n", fullPath)
		for k, v := range data {
			fmt.Printf("\033[33m%s\033[0m: %v\n", k, v)
		}
	}
}
