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
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// Define a global variable to store the encoded state file
var encodedStateFile string

// Wrap command implementation
func Wrap() {
	// Get a list of files in the current directory
	files, err := filepath.Glob("*.tfstate")
	if err != nil {
		fmt.Println("❌ Error searching for .tfstate files:", err)
		os.Exit(1)
	}

	// Check if no .tfstate files were found
	if len(files) == 0 {
		fmt.Println("❌ Error: No .tfstate files found in the current directory.")
		os.Exit(1)
	}

	// Use the first found .tfstate file
	stateFilePath := files[0]

	// Read the contents of the found .tfstate file using os.ReadFile (Go 1.16+)
	stateFileContents, err := os.ReadFile(stateFilePath)
	if err != nil {
		fmt.Println("❌ Error reading state file:", err)
		os.Exit(1)
	}

	// Compress the state file using gzip
	compressedStateFile, err := gzipStateFile(stateFileContents)
	if err != nil {
		fmt.Println("❌ Error compressing state file:", err)
		os.Exit(1)
	}

	// Encode the compressed state file using base64
	encodedStateFile = base64.StdEncoding.EncodeToString(compressedStateFile)

	// Set the environment variable
	os.Setenv("TERRAFORM_STATE_BASE64", encodedStateFile)

	// Check if the environment variable is set correctly
	if value := os.Getenv("TERRAFORM_STATE_BASE64"); value != encodedStateFile {
		fmt.Println("❌ Error: Failed to set the environment variable TERRAFORM_STATE_BASE64 correctly.")
		os.Exit(1)
	}

	tempFilePath := ".encoded_wrap"
	err = saveEncodedStateToFile(encodedStateFile, tempFilePath)
	if err != nil {
		fmt.Println("❌ Error saving encoded state file:", err)
		os.Exit(1)
	}

	fmt.Println("✅", stateFilePath, "wrapped successfully.")
	fmt.Println("Encoded state file saved to:", tempFilePath)
}

func saveEncodedStateToFile(encodedState string, filePath string) error {
	// Create or open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the encoded state to the file
	_, err = file.WriteString(encodedState)
	if err != nil {
		return err
	}

	return nil
}

func gzipStateFile(data []byte) ([]byte, error) {
	var compressedData bytes.Buffer
	writer := gzip.NewWriter(&compressedData)

	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return compressedData.Bytes(), nil
}
