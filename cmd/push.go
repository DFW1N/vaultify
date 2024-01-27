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
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Push command implementation
func Push() {
	// Look for a .encoded_wrap file in the current directory
	encodedStateFilePath := ".encoded_wrap"

	// Check if the .encoded_wrap file exists
	if _, err := os.Stat(encodedStateFilePath); os.IsNotExist(err) {
		fmt.Println("❌ Error: .encoded_wrap file not found in the current directory.")
		fmt.Println("Please run 'vaultify wrap' to create the .encoded_wrap file.")
		return
	}

	// Read the contents of the .encoded_wrap file
	encodedStateFileContents, err := os.ReadFile(encodedStateFilePath) // Use os.ReadFile for Go 1.16+
	if err != nil {
		fmt.Println("❌ Error reading .encoded_wrap file:", err)
		return
	}

	// Convert the contents to a string (assuming it contains the encoded state)
	encodedStateFile := string(encodedStateFileContents)

	// Check if the read content is empty
	if encodedStateFile == "" {
		fmt.Println("❌ Error: .encoded_wrap file is empty.")
		return
	}

	// Validate if the content is in a valid base64 format
	if !isValidBase64(encodedStateFile) {
		fmt.Println("❌ Error: .encoded_wrap file does not contain valid base64 data.")
		return
	}

	// Set the environment variable
	os.Setenv("TERRAFORM_STATE_BASE64", encodedStateFile)

	// Simulate the Vault update (replace with your actual implementation)
	// Example: Your code to push the encoded state to Vault goes here

	// Construct the curl command
	curlCommand := "curl"

	// Prepare the encoded payload
	encodedPayload := encodedStateFile

	// Construct the Vault URL using the VAULT_ADDR environment variable
	vaultURL := os.Getenv("VAULT_ADDR")

	// Define the Vault engine name (replace 'kv' with your desired engine name)
	engineName := "kv"

	// Construct the secret path
	dataPath := "vaultify"

	// Get the current Terraform workspace name
	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		fmt.Println("❌ Error getting current Terraform workspace:", err)
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
	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workspaceName, workingDirName, "terraform.tfstate")

	// Check if the secret path exists in Vault
	checkPathCmd := exec.Command(
		curlCommand,
		"--silent", "--show-error",
		"--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"),
		"--request", "GET",
		vaultURL+"/v1/"+engineName+"/data/"+dataPath,
	)

	checkPathOutput, err := checkPathCmd.CombinedOutput()

	// Check for errors
	if err != nil {
		fmt.Println("❌ Error checking if secret path exists:", err)
		return
	}

	// Check the HTTP response status code
	pathStatus := strings.TrimSpace(string(checkPathOutput))

	// If the path doesn't exist, create it
	if pathStatus == "404" {
		createPathCmd := exec.Command(
			curlCommand,
			"--silent", "--show-error",
			"--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"),
			"--request", "POST",
			"--data-raw", `{"type": "kv"}`,
			vaultURL+"/v1/"+engineName+"/data/"+dataPath,
		)

		createPathOutput, err := createPathCmd.CombinedOutput()

		// Check for errors
		if err != nil {
			fmt.Println("❌ Error creating secret path:", err)
			return
		}

		// Check the HTTP response status code after creating the path
		if !strings.Contains(string(createPathOutput), "success") {
			fmt.Println("❌ Failed to create secret path:", string(createPathOutput))
			return
		}
	}

	// At this point, the path exists (or has been created), so proceed to push the secret

	// Execute the curl command to push the secret to Vault with specified values
	pushCmd := exec.Command(
		curlCommand,
		"--silent", "--show-error",
		"--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"),
		"--request", "PUT",
		"--data-raw", "{\"data\": {\""+secretPath+"\": \""+encodedPayload+"\"}}",
		"--write-out", "%{http_code}",
		"--output", "response.json",
		vaultURL+"/v1/"+engineName+"/data/"+secretPath,
	)

	pushOutput, err := pushCmd.CombinedOutput()

	// Check for errors
	if err != nil {
		fmt.Println("❌ Error pushing secret to Vault:", err)
		return
	}

	// Check the HTTP response status code after pushing the secret
	httpStatus := strings.TrimSpace(string(pushOutput))
	if httpStatus == "200" || httpStatus == "204" {
		fmt.Printf("✅ Secret written to HashiCorp Vault under: %s\n", secretPath)
		fmt.Printf("💠 The file size uploaded to Hashicorp Vault: %.2f KB\n", float64(len(encodedStateFile))/1024)

		// Simulate terraformApplyStatus (replace with your actual logic)
		if false {
			fmt.Println("❗ Vault updated, but apply failed. Exiting...")
			return
		}
	} else {
		fmt.Println("❌ Failed to write secret to Hashicorp Vault.")
		fmt.Printf("Response code: %s\n", httpStatus)
		fmt.Println("dataPath:", dataPath)
		fmt.Println("engineName:", engineName)
		fmt.Println("workspaceName:", workspaceName)
		fmt.Println("workingDir:", workingDir)
		secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workspaceName, workingDirName, "terraform.tfstate")
		fmt.Println("secretPath:", secretPath)
		fmt.Println("Please review the values being used to validate that has been printed above.")
		return
	}

	fmt.Println("✅ Pushed terraform.tfstate to HashiCorp Vault successfully.")
}

// isValidBase64 checks if a given string is in a valid base64 format
func isValidBase64(input string) bool {
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}

// getCurrentWorkspace gets the current Terraform workspace name
func getCurrentWorkspace() (string, error) {
	// Run the 'terraform workspace show' command to get the current workspace
	cmd := exec.Command("terraform", "workspace", "show")

	// Capture the command output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// Trim any leading/trailing whitespace and return the workspace name
	return strings.TrimSpace(string(output)), nil
}
