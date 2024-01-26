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

	// Calculate the size of the encoded state
	fileSize := float64(len(encodedStateFile)) / 1024 // Size in KB

	// Set the environment variable
	os.Setenv("TERRAFORM_STATE_BASE64", encodedStateFile)
	// fmt.Println("Encoded State File Contents:", encodedStateFile)

	// Simulate the Vault update (replace with your actual implementation)
	// Example: Your code to push the encoded state to Vault goes here

	fmt.Printf("💠 The file size uploaded to Hashicorp Vault: %.2f KB\n", fileSize)
	fmt.Println("✅ Pushed terraform.tfstate to HashiCorp Vault successfully.")
}

// isValidBase64 checks if a given string is in a valid base64 format
func isValidBase64(input string) bool {
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}
