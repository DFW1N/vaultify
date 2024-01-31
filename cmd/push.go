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

	// Check for .vaultify directory and settings.json
	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
		return
	}

	encodedStateFilePath := "/tmp/.encoded_wrap"

	if _, err := os.Stat(encodedStateFilePath); os.IsNotExist(err) {
		fmt.Println("❌ Error: \033[33m.encoded_wrap\033[0m file not found in the \033[33m/tmp\033[0m directory.")
		fmt.Println("Please run \033[33m'vaultify wrap'\033[0m to create the \033[33m.encoded_wrap\033[0m file.")
		return
	}

	encodedStateFileContents, err := os.ReadFile(encodedStateFilePath)
	if err != nil {
		fmt.Println("❌ Error reading \033[33m.encoded_wrap\033[0m file:", err)
		return
	}

	encodedStateFile := string(encodedStateFileContents)
	if encodedStateFile == "" {
		fmt.Println("❌ Error: \033[33m.encoded_wrap\033[0m file is empty.")
		return
	}

	if !isValidBase64(encodedStateFile) {
		fmt.Println("❌ Error: \033[33m.encoded_wrap\033[0m file does not contain valid base64 data.")
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
	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workingDirName, workspaceName, "terraform.tfstate")

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
		fmt.Printf("✅ Secret written to HashiCorp Vault under: \033[33m%s\033[0m\n", secretPath)
		fmt.Printf("💠 The file size uploaded to Hashicorp Vault: \033[33m%.2f\033[0m KB\n", float64(len(encodedStateFile))/1024)

		// Delete the terraform.tfstate file only if it exists
		if _, err := os.Stat("terraform.tfstate"); err == nil {
			if err := os.Remove("terraform.tfstate"); err != nil {
				fmt.Println("❌ Error: Failed to delete the \033[33mterraform.tfstate\033[0m file.", err)
				return
			}
			//fmt.Println("✅ Deleted the terraform.tfstate file.")
		}

		if err := os.Remove(encodedStateFilePath); err != nil {
			fmt.Println("❌ Error: Failed to delete the \033[33m/tmp/.encoded_wrap\033[0m file.", err)
			return
		}
		//fmt.Println("✅ Deleted the /tmp/.encoded_wrap file.")
	} else {
		fmt.Println("❌ Failed to write secret to Hashicorp Vault.")
		fmt.Printf("Response code: \033[33m%s\033[0m\n", httpStatus)
		// ... [additional error handling code if required]
	}
}

// isValidBase64 checks if a given string is in a valid base64 format
func isValidBase64(input string) bool {
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}
