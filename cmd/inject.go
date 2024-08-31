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
	"path/filepath"
	"strings"
)

func Inject(args []string) {
	injectCmd := flag.NewFlagSet("inject", flag.ExitOnError)
	path := injectCmd.String("path", "", "Vault path to store the secret")
	key := injectCmd.String("key", "", "Name of the secret (will be used as the path if no path is provided)")
	err := injectCmd.Parse(args)
	if err != nil {
		fmt.Println("❌ Error parsing inject flags:", err)
		return
	}

	if injectCmd.NArg() < 1 {
		fmt.Println("Usage: vaultify inject [-path <vault_path>] [-key <secret_name>] <secret_value>")
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

	// Always start with secrets/
	secretPath := "secrets"

	// If path is provided, use it. Otherwise, use the key as the path.
	if *path != "" {
		secretPath = filepath.Join(secretPath, strings.Trim(*path, "/"))
	} else if *key != "" {
		secretPath = filepath.Join(secretPath, *key)
	} else {
		// If neither path nor key is provided, use a default
		secretPath = filepath.Join(secretPath, "default")
	}

	// Prepare the secret data
	secretData := map[string]interface{}{
		"data": map[string]interface{}{
			"value": secretValue,
		},
	}

	// Write the secret to Vault
	fullPath := fmt.Sprintf("%s/data/%s", engineName, secretPath)

	_, err = vaultClient.Logical().Write(fullPath, secretData)
	if err != nil {
		fmt.Println("❌ Error injecting secret to Vault:", err)
		return
	}

	fmt.Printf("✅ Secret injected to HashiCorp Vault under: \033[33m%s\033[0m\n", fullPath)
	fmt.Printf("💠 Secret Name: \033[33m%s\033[0m\n", filepath.Base(secretPath))

	if err := LogHistory("inject", fullPath); err != nil {
		fmt.Printf("❌ Error logging history: %v\n", err)
	}
}
