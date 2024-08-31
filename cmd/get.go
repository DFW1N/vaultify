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
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func Get(args []string) {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	path := getCmd.String("path", "", "Vault path to retrieve the secret from")
	key := getCmd.String("key", "", "Name of the secret (will be used as the path if no path is provided)")
	jsonOutput := getCmd.Bool("json", false, "Output in JSON format")
	yamlOutput := getCmd.Bool("yaml", false, "Output in YAML format")
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

	secretPath := "secrets"

	if *path != "" {
		secretPath = filepath.Join(secretPath, strings.Trim(*path, "/"))
	} else if *key != "" {
		secretPath = filepath.Join(secretPath, *key)
	} else {
		secretPath = filepath.Join(secretPath, "default")
	}

	// Read the secret from Vault
	fullPath := fmt.Sprintf("%s/data/%s", engineName, secretPath)

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

	value, exists := data["value"]
	if !exists {
		fmt.Printf("❌ No value found in secret at path \033[33m%s\033[0m\n", fullPath)
		return
	}

	if *jsonOutput {
		outputJSON(fullPath, secretPath, value)
	} else if *yamlOutput {
		outputYAML(fullPath, secretPath, value)
	} else {
		fmt.Printf("%v\n", value)
	}
}

func outputJSON(fullPath, secretPath string, value interface{}) {
	output := map[string]interface{}{
		"path":        fullPath,
		"secret_name": filepath.Base(secretPath),
		"value":       value,
	}
	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("❌ Error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func outputYAML(fullPath, secretPath string, value interface{}) {
	output := map[string]interface{}{
		"path":        fullPath,
		"secret_name": filepath.Base(secretPath),
		"value":       value,
	}
	yamlData, err := yaml.Marshal(output)
	if err != nil {
		fmt.Println("❌ Error marshaling to YAML:", err)
		return
	}
	fmt.Println(string(yamlData))
}
