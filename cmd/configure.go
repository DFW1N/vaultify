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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Configuration structure to hold settings
type Configuration struct {
	Settings struct {
		DefaultEngineName   string `json:"default_engine_name"`
		TerraformWorkspaces bool   `json:"terraform_workspace"`
	} `json:"settings"`
}

// Read current configuration from file
func readConfiguration() (*Configuration, error) {
	homeDir, _ := os.UserHomeDir()
	configFile := filepath.Join(homeDir, ".vaultify", "settings.json")

	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Write configuration to file
func writeConfiguration(config *Configuration) error {
	homeDir, _ := os.UserHomeDir()
	configFile := filepath.Join(homeDir, ".vaultify", "settings.json")

	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(configFile, data, 0644)
}

// Validate if the engine name is supported
func isValidEngineName(name string) bool {
	// Add logic to validate engine name
	return true // Placeholder: replace with actual validation
}

// Configure command implementation
func Configure() {

	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
		return
	}

	config, err := readConfiguration()
	if err != nil {
		fmt.Println("❌ Error reading configuration:", err)
		return
	}

	for {
		// Present the configuration options
		fmt.Println("\n\033[33mVaultify\033[0m Settings")
		fmt.Println("")
		fmt.Printf("%-20s %s\n", "\033[33mOption\033[0m", "\033[33mValue\033[0m")
		fmt.Println("-------------------------------------------")
		fmt.Printf("%-20s \033[33m%s\033[0m\n", "\033[33m1\033[0m. Default Engine Name:", config.Settings.DefaultEngineName)
		fmt.Printf("%-20s \033[33m%t\033[0m\n", "\033[33m2\033[0m. Use Terraform Workspaces:", config.Settings.TerraformWorkspaces)
		fmt.Println()

		fmt.Println("Enter the number of the option you want to change (or \033[33m0\033[0m to exit):")

		// Read user input
		var choice int
		fmt.Print("\033[33mChoice\033[0m: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a \033[33mnumber\033[0m.")
			continue
		}

		switch choice {
		case 0:
			fmt.Println("\nExiting configuration.")
			return
		case 1:
			fmt.Print("\n\033[33mEnter new Default Engine Name\033[0m: ")
			fmt.Scanln(&config.Settings.DefaultEngineName)
			if !isValidEngineName(config.Settings.DefaultEngineName) {
				fmt.Println("Invalid engine name. Please enter a valid \033[33mengine name\033[0m.")
				continue
			}
		case 2:
			var workspaceInput string
			fmt.Print("\nEnable Terraform Workspaces (\033[33mtrue\033[0m/\033[33mfalse\033[0m): ")
			fmt.Scanln(&workspaceInput)
			workspaceInput = strings.ToLower(workspaceInput)
			if workspaceInput == "true" {
				config.Settings.TerraformWorkspaces = true
			} else if workspaceInput == "false" {
				config.Settings.TerraformWorkspaces = false
			} else {
				fmt.Println("Invalid input. Please enter \033[33m'true'\033[0m or \033[33m'false'\033[0m.")
				continue
			}
		default:
			fmt.Println("Invalid option. Please enter a \033[33mvalid number\033[0m.")
			continue
		}

		// Save updated configuration
		if err := writeConfiguration(config); err != nil {
			fmt.Println("❌ Error saving \033[33mconfiguration\033[0m:", err)
			return
		}

		fmt.Println("\nConfiguration in \033[33m`settings.json`\033[0m updated successfully.")
	}
}
