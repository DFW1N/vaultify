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

// File: cmd/init.go
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Settings structure for the JSON content
type Settings struct {
	Settings struct {
		TerraformWorkspace bool   `json:"terraform_workspace"`
		DefaultEngineName  string `json:"default_engine_name"`
	} `json:"settings"`
}

// Init command implementation
func Init() {
	// Create .vaultify folder in the user's home directory

	vaultifyDir, err := ensureVaultifyFolder()
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	if err := createSettingsFile(vaultifyDir); err != nil {
		fmt.Printf("❌ Error creating settings.json: %v\n", err)
		os.Exit(1)
	}

	// Retrieve environment variables
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddr := os.Getenv("VAULT_ADDR")

	// Check if environment variables are set
	if vaultToken == "" || vaultAddr == "" {
		fmt.Println("❌ Error: \033[33mVAULT_TOKEN\033[0m and \033[33mVAULT_ADDR\033[0m environment variables must be set.")
		os.Exit(1)
	}

	// createVaultifyFolder creates a .vaultify folder in the user's home directory
	if err := initializeVaultify(vaultToken, vaultAddr); err != nil {
		fmt.Printf("❌ Error initializing \033[33mVaultify\033[0m: %v\n", err)
		os.Exit(1)
	} else {

		fmt.Println("✅ \033[33mVaultify\033[0m initialized successfully.")
	}
}

func ensureVaultifyFolder() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	vaultifyDir := filepath.Join(homeDir, ".vaultify")

	if _, err := os.Stat(vaultifyDir); os.IsNotExist(err) {
		if err := os.Mkdir(vaultifyDir, 0700); err != nil {
			return "", fmt.Errorf("failed to create \033[33m.vaultify\033[0m directory: %w", err)
		}
		fmt.Println("✅ Created \033[33m.vaultify\033[0m folder")
	} else if err != nil {
		return "", fmt.Errorf("error checking \033[33m.vaultify\033[0m directory: %w", err)
	}

	return vaultifyDir, nil
}

func createSettingsFile(vaultifyDir string) error {
	settingsFilePath := filepath.Join(vaultifyDir, "settings.json")

	// Check if settings.json already exists
	if _, err := os.Stat(settingsFilePath); err == nil {
		fmt.Println("✅ \033[33msettings.json\033[0m already exists")
		return nil // File exists, no need to create
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking \033[33msettings.json\033[0m: %w", err)
	}

	// Create settings.json with specified content
	settings := Settings{}
	settings.Settings.TerraformWorkspace = true
	settings.Settings.DefaultEngineName = "kv"

	jsonData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings to JSON: %w", err)
	}

	if err := os.WriteFile(settingsFilePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write \033[33msettings.json\033[0m: %w", err)
	}

	fmt.Println("✅ Generated \033[33msettings.json\033[0m")
	return nil
}

func initializeVaultify(token, addr string) error {
	return nil
}
