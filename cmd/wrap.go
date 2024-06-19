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

var encryptedStateFile string
var encodedStateFile string

func Wrap() {
	passphrase := os.Getenv("VAULTIFY_PASSPHRASE")
	if passphrase == "" {
		fmt.Println("❌ Error: \033[33mVAULTIFY_PASSPHRASE\033[0m environemnt variable not set.")
		os.Exit(1)
	}
	files, err := filepath.Glob("*.tfstate")
	if err != nil {
		fmt.Println("❌ Error searching for .tfstate files:", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("❌ Error: No .tfstate files found in the current directory.")
		fmt.Println("⚠️  Please run vaultify pull and vaultify unwrap to get this file from your vault, if it doesn't exist locally.")
		os.Exit(1)
	}

	stateFilePath := files[0]

	stateFileContents, err := os.ReadFile(stateFilePath)
	if err != nil {
		fmt.Println("❌ Error reading state file:", err)
		os.Exit(1)
	}

	compressedStateFile, err := gzipStateFile(stateFileContents)
	if err != nil {
		fmt.Println("❌ Error compressing state file:", err)
		os.Exit(1)
	}

	encryptedFile, err := encryptContents(compressedStateFile, passphrase)
	if err != nil {
		fmt.Println("❌ Error encrypting contents of compressed state file:", err)
		os.Exit(1)
	}

	encodedStateFile = base64.StdEncoding.EncodeToString(encryptedFile)

	// Set the environment variable
	os.Setenv("TERRAFORM_STATE_BASE64", encodedStateFile)

	if value := os.Getenv("TERRAFORM_STATE_BASE64"); value != encodedStateFile {
		fmt.Println("❌ Error: Failed to set the environment variable TERRAFORM_STATE_BASE64 correctly.")
		os.Exit(1)
	}

	tempFilePath := "/tmp/.encoded_wrap"
	err = saveEncodedStateToFile(encodedStateFile, tempFilePath)
	if err != nil {
		fmt.Println("❌ Error saving encoded state file:", err)
		os.Exit(1)
	}

	if err := os.Remove(stateFilePath); err != nil {
		fmt.Println("❌ Error: Failed to delete the original .tfstate file.", err)
		os.Exit(1)
	}
	//fmt.Printf("✅ Deleted the original .tfstate file: %s\n", stateFilePath)

	fmt.Println("✅", stateFilePath, "wrapped successfully.")
	fmt.Println("⚠️  Please run 'vaultify push' to store your encoded state into your Hashicorp Vault.")
}

func saveEncodedStateToFile(encodedState string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

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
