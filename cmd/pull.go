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
	"os/exec"
	"path/filepath"
	"strings"
)

// VaultResponse represents the structure of the JSON response
type VaultResponse struct {
	Data struct {
		Data map[string]string `json:"data"`
	} `json:"data"`
}

// Pull command implementation
func Pull() {

	// Check for .vaultify directory and settings.json
	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
		return
	}

	// Read settings from settings.json
	settings, err := readSettings()
	if err != nil {
		fmt.Println("❌ Error reading settings:", err)
		return
	}

	// Construct the curl command
	curlCommand := "curl"

	// Construct the Vault URL using the VAULT_ADDR environment variable
	vaultURL := os.Getenv("VAULT_ADDR")

	// Use the engine name from settings
	engineName := settings.Settings.DefaultEngineName

	// Construct the secret path
	dataPath := "vaultify"

	// Get the current Terraform workspace name
	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		fmt.Println("❌ Error getting current \033[33mTerraform\033[0m workspace:", err)
		return
	}

	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Error getting current working directory:", err)
		return
	}

	// Extract the folder name from the full path
	workingDirName := filepath.Base(workingDir)

	// Construct the complete secret path including the workspace name and working directory
	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workingDirName, workspaceName, "terraform.tfstate")

	// Check if the secret path exists in Vault
	checkPathCmd := exec.Command(
		curlCommand,
		"--silent", "--show-error",
		"--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"),
		"--request", "GET",
		"--silent", "--write-out", "%{http_code}",
		"--output", "/dev/null",
		vaultURL+"/v1/"+engineName+"/data/"+secretPath,
	)

	checkPathOutput, err := checkPathCmd.CombinedOutput()

	// Check for errors
	if err != nil {
		fmt.Println("❌ Error checking if secret path exists:", err)
		return
	}

	// Check the HTTP response status code
	pathStatus := strings.TrimSpace(string(checkPathOutput))

	// If the path doesn't exist, exit with an error message
	if pathStatus == "404" {
		fmt.Printf("❌ Error: Secret path not found in HashiCorp Vault. Path: \033[33m%s\033[0m\n", vaultURL+"/v1/"+engineName+"/data/"+secretPath)
		return
	}

	fmt.Println("✅ Secret exists in Vault. Retrieving...")

	// Retrieve the secret from Vault
	pullCmd := exec.Command(
		curlCommand,
		"--silent", "--show-error",
		"--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"),
		"--request", "GET",
		vaultURL+"/v1/"+engineName+"/data/"+secretPath,
	)

	// Capture the command's output
	pullOutput, err := pullCmd.Output()
	if err != nil {
		fmt.Println("❌ Error retrieving \033[33msecret\033[0m from Vault:", err)
		return
	}

	// Unmarshal the JSON response into the VaultResponse struct
	var response VaultResponse
	err = json.Unmarshal(pullOutput, &response)
	if err != nil {
		fmt.Println("❌ Error unmarshalling \033[33mJSON\033[0m:", err)
		return
	}

	// Construct the key dynamically
	dynamicKey := fmt.Sprintf("%s/%s/%s_%s", dataPath, workingDirName, workspaceName, "terraform.tfstate")

	// Extract the base64 encoded string using the dynamic key
	base64String, ok := response.Data.Data[dynamicKey]
	if !ok {
		fmt.Println("❌ Error: Specific \033[33mkey\033[0m not found in the data")
		return
	}

	targetFilePath := "terraform.tfstate.gz.b64"
	if _, err := os.Stat(targetFilePath); err == nil {
		fmt.Println("❌ Error: File \033[33mterraform.tfstate.gz.b64\033[0m already exists in the directory.")
		return
	} else if !os.IsNotExist(err) {
		fmt.Printf("❌ Error checking if file exists: %v\n", err)
		return
	}

	// Save the base64 encoded string to the file
	if err := saveStateToFile([]byte(base64String), targetFilePath); err != nil {
		fmt.Println("❌ Error saving base64 string to file:", err)
		return
	}

	fmt.Println("✅ Secret retrieved and saved as \033[33mterraform.tfstate.gz.b64\033[0m")
}

// saveStateToFile saves the state data to the specified file
func saveStateToFile(data []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
