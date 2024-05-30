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
)

type AzureSettings struct {
	StorageAccountName              string `json:"storage_account_name"`
	StorageAccountResourceGroupName string `json:"storage_account_resource_group_name"`
}

type AwsSettings struct {
	S3BucketName string `json:"s3_bucket_name"`
}

type Configuration struct {
	Settings struct {
		TerraformWorkspace   bool          `json:"terraform_workspace"`
		DefaultEngineName    string        `json:"default_engine_name"`
		DefaultSecretStorage string        `json:"default_secret_storage"`
		Azure                AzureSettings `json:"azure"`
		AWS                  AwsSettings   `json:"aws"`
	} `json:"settings"`
}

func Init() {
	vaultifyDir, err := ensureVaultifyFolder()
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	if err := createSettingsFile(vaultifyDir); err != nil {
		fmt.Printf("❌ Error creating \033[33msettings.json\033[0m: %v\n", err)
		os.Exit(1)
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultToken == "" || vaultAddr == "" {
		fmt.Println("❌ Error: \033[33mVAULT_TOKEN\033[0m and \033[33mVAULT_ADDR\033[0m environment variables must be set.\033[0m")
		os.Exit(1)
	}

	fmt.Println("✅ \033[33mVaultify\033[0m initialized successfully.\033[0m")
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
		fmt.Println("✅ Created \033[33m.vaultify\033[0m folder.\033[0m")
	} else if err != nil {
		return "", fmt.Errorf("error checking \033[33m.vaultify\033[0m directory: %w", err)
	}

	return vaultifyDir, nil
}

func createSettingsFile(vaultifyDir string) error {
	settingsFilePath := filepath.Join(vaultifyDir, "settings.json")

	if _, err := os.Stat(settingsFilePath); err == nil {
		fmt.Println("✅ \033[33msettings.json\033[0m already exists.\033[0m")
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking \033[33msettings.json\033[0m: \033[33m%w\033[0m", err)
	}

	settings := Configuration{
		Settings: struct {
			TerraformWorkspace   bool          `json:"terraform_workspace"`
			DefaultEngineName    string        `json:"default_engine_name"`
			DefaultSecretStorage string        `json:"default_secret_storage"`
			Azure                AzureSettings `json:"azure"`
			AWS                  AwsSettings   `json:"aws"`
		}{
			TerraformWorkspace:   true,
			DefaultEngineName:    "kv",
			DefaultSecretStorage: "vault",
			Azure: AzureSettings{
				StorageAccountName:              "",
				StorageAccountResourceGroupName: "",
			},
			AWS: AwsSettings{
				S3BucketName: "",
			},
		},
	}

	jsonData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings to JSON: \033[33m%w\033[0m", err)
	}

	if err := os.WriteFile(settingsFilePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write \033[33msettings.json\033[0m: %w", err)
	}

	fmt.Println("✅ Generated \033[33msettings.json\033[0m successfully.\033[0m")
	return nil
}
