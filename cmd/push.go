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
	encodedStateFilePath := "/tmp/.encoded_wrap"

	if _, err := os.Stat(encodedStateFilePath); os.IsNotExist(err) {
		fmt.Println("❌ Error: .encoded_wrap file not found in the /tmp directory.")
		fmt.Println("Please run 'vaultify wrap' to create the .encoded_wrap file.")
		return
	}

	encodedStateFileContents, err := os.ReadFile(encodedStateFilePath)
	if err != nil {
		fmt.Println("❌ Error reading .encoded_wrap file:", err)
		return
	}

	encodedStateFile := string(encodedStateFileContents)
	if encodedStateFile == "" {
		fmt.Println("❌ Error: .encoded_wrap file is empty.")
		return
	}

	if !isValidBase64(encodedStateFile) {
		fmt.Println("❌ Error: .encoded_wrap file does not contain valid base64 data.")
		return
	}

	os.Setenv("TERRAFORM_STATE_BASE64", encodedStateFile)
	curlCommand := "curl"
	encodedPayload := encodedStateFile
	vaultURL := os.Getenv("VAULT_ADDR")
	engineName := "kv"
	dataPath := "vaultify"
	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		fmt.Println("❌ Error getting current Terraform workspace:", err)
		return
	}

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Error getting current working directory:", err)
		return
	}

	workingDirName := filepath.Base(workingDir)
	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workspaceName, workingDirName, "terraform.tfstate")

	checkPathCmd := exec.Command(curlCommand, "--silent", "--show-error", "--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"), "--request", "GET", vaultURL+"/v1/"+engineName+"/data/"+dataPath)
	checkPathOutput, err := checkPathCmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Error checking if secret path exists:", err)
		return
	}

	pathStatus := strings.TrimSpace(string(checkPathOutput))
	if pathStatus == "404" {
		createPathCmd := exec.Command(curlCommand, "--silent", "--show-error", "--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"), "--request", "POST", "--data-raw", `{"type": "kv"}`, vaultURL+"/v1/"+engineName+"/data/"+dataPath)
		createPathOutput, err := createPathCmd.CombinedOutput()
		if err != nil {
			fmt.Println("❌ Error creating secret path:", err)
			return
		}
		if !strings.Contains(string(createPathOutput), "success") {
			fmt.Println("❌ Failed to create secret path:", string(createPathOutput))
			return
		}
	}

	pushCmd := exec.Command(curlCommand, "--silent", "--show-error", "--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"), "--request", "PUT", "--data-raw", "{\"data\": {\""+secretPath+"\": \""+encodedPayload+"\"}}", "--write-out", "%{http_code}", "--output", "response.json", vaultURL+"/v1/"+engineName+"/data/"+secretPath)
	pushOutput, err := pushCmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Error pushing secret to Vault:", err)
		return
	}

	httpStatus := strings.TrimSpace(string(pushOutput))
	if httpStatus == "200" || httpStatus == "204" {
		fmt.Printf("✅ Secret written to HashiCorp Vault under: %s\n", secretPath)

		// Delete the terraform.tfstate file only if it exists
		if _, err := os.Stat("terraform.tfstate"); err == nil {
			if err := os.Remove("terraform.tfstate"); err != nil {
				fmt.Println("❌ Error: Failed to delete the terraform.tfstate file.", err)
				return
			}
			//fmt.Println("✅ Deleted the terraform.tfstate file.")
		}

		if err := os.Remove(encodedStateFilePath); err != nil {
			fmt.Println("❌ Error: Failed to delete the /tmp/.encoded_wrap file.", err)
			return
		}
		fmt.Println("✅ Deleted the /tmp/.encoded_wrap file.")
	} else {
		fmt.Println("❌ Failed to write secret to Hashicorp Vault.")
		fmt.Printf("Response code: %s\n", httpStatus)
		// ... [additional error handling code if required]
	}
}

// isValidBase64 checks if a given string is in a valid base64 format
func isValidBase64(input string) bool {
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}

// getCurrentWorkspace gets the current Terraform workspace name
func getCurrentWorkspace() (string, error) {
	cmd := exec.Command("terraform", "workspace", "show")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
