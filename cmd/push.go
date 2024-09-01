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
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	vault "github.com/hashicorp/vault/api"
)

func Push() {

	config, err := readConfiguration()
	if err != nil {
		fmt.Println("❌ \033[33mError\033[0m loading configuration:", err)
		return
	}

	defaultSecretStorage := config.Settings.DefaultSecretStorage
	accountName := config.Settings.Azure.StorageAccountName

	var pushLocation string
	var pushErr error

	switch defaultSecretStorage {
	case "vault":
		pushLocation, pushErr = pushToVault()
	case "azure_storage":
		key, err := listStorageAccountKeys()
		if err != nil {
			fmt.Printf("Failed to list storage account keys: \033[33m%v\033[0m\n", err)
			return
		}
		createContainer(accountName, key)
		encodedStateFilePath := "/tmp/.encoded_wrap"
		pushLocation, pushErr = uploadBlobWithAccessKey(accountName, key, encodedStateFilePath)
	case "s3":
		fmt.Println("⚠️ \033[33m AWS S3 Bucket\033[0m is currently under development.")
	default:
		fmt.Println("Unsupported secret storage specified.")
	}

	if pushErr != nil {
		fmt.Printf("❌ Error pushing to %s: %v\n", defaultSecretStorage, pushErr)
		return
	}

	if pushLocation != "" {
		if err := LogHistory("push", pushLocation); err != nil {
			fmt.Printf("❌ Error logging history: %v\n", err)
		}
		if err := LogUserAction("push", pushLocation); err != nil {
			fmt.Printf("❌ Error logging user action: %v\n", err)
		}
		fmt.Printf("✅ Successfully pushed state to %s\n", pushLocation)
	}
}

func pushToVault() (string, error) {
	if err := checkVaultifySetup(); err != nil {
		return "", fmt.Errorf("vaultify setup error: %w", err)
	}

	vaultClient, initStat := initVaultClientWithStatus()
	if !initStat {
		return "", fmt.Errorf("vault is not initialized")
	}

	settings, err := readSettings()
	if err != nil {
		return "", fmt.Errorf("error reading settings: %w", err)
	}

	encodedStateFilePath := "/tmp/.encoded_wrap"

	if _, err := os.Stat(encodedStateFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf(".encoded_wrap file not found in the /tmp directory")
	}

	encodedStateFileContents, err := os.ReadFile(encodedStateFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading .encoded_wrap file: %w", err)
	}

	encodedStateFile := string(encodedStateFileContents)
	if encodedStateFile == "" {
		return "", fmt.Errorf(".encoded_wrap file is empty")
	}

	if !isValidBase64(encodedStateFile) {
		return "", fmt.Errorf(".encoded_wrap file does not contain valid base64 data")
	}

	os.Setenv("TERRAFORM_STATE_BASE64", encodedStateFile)
	encodedPayload := encodedStateFile

	engineName := settings.Settings.DefaultEngineName
	dataPath := "vaultify"

	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		return "", fmt.Errorf("error getting current Terraform workspace: %w", err)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %w", err)
	}

	workingDirName := filepath.Base(workingDir)
	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workingDirName, workspaceName, "terraform.tfstate")
	fullPath := fmt.Sprintf("%s/data/%s", engineName, secretPath)

	err = ensureKVPathExists(vaultClient, engineName, dataPath)
	if err != nil {
		return "", fmt.Errorf("unable to perform operation: %w", err)
	}

	secretData := map[string]interface{}{
		"data": map[string]interface{}{
			secretPath: encodedPayload,
		},
	}

	_, err = vaultClient.Logical().Write(fullPath, secretData)
	if err != nil {
		return "", fmt.Errorf("error pushing secret to Vault: %w", err)
	}

	fmt.Printf("✅ Secret written to HashiCorp Vault under: \033[33m%s\033[0m\n", fullPath)
	fmt.Printf("💠 The file size uploaded to Hashicorp Vault: \033[33m%.2f\033[0m KB\n", float64(len(encodedStateFile))/1024)

	if _, err := os.Stat("terraform.tfstate"); err == nil {
		if err := os.Remove("terraform.tfstate"); err != nil {
			return "", fmt.Errorf("failed to delete the terraform.tfstate file: %w", err)
		}
	}

	if err := os.Remove(encodedStateFilePath); err != nil {
		return "", fmt.Errorf("failed to delete the /tmp/.encoded_wrap file: %w", err)
	}

	return fmt.Sprintf("vault:%s", fullPath), nil
}

func uploadBlobWithAccessKey(accountName, key, encodedStateFilePath string) (string, error) {
	containerName := "vaultify"

	if err := checkVaultifySetup(); err != nil {
		errMsg := fmt.Sprintf("%v\nPlease run 'vaultify init' to set up Vaultify.", err)
		return "", fmt.Errorf(errMsg)
	}

	if _, err := os.Stat(encodedStateFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("❌ Error: .encoded_wrap file not found in the /tmp directory. Please run 'vaultify wrap' to create the .encoded_wrap file")
	} else if err != nil {
		return "", fmt.Errorf("❌ Error checking .encoded_wrap file: %v", err)
	}

	encodedStateFileContents, err := os.ReadFile(encodedStateFilePath)
	if err != nil {
		return "", fmt.Errorf("❌ Error reading .encoded_wrap file: %v", err)
	}

	if len(encodedStateFileContents) == 0 {
		return "", fmt.Errorf("❌ Error: .encoded_wrap file is empty")
	}

	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		return "", fmt.Errorf("❌ Error getting current Terraform workspace: %v", err)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("❌ Error getting current working directory: %v", err)
	}

	workingDirName := filepath.Base(workingDir)
	blobName := fmt.Sprintf("%s/%s_%s", workingDirName, workspaceName, "terraform.tfstate")

	file, err := os.Open(encodedStateFilePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file contents: %v", err)
	}

	method := "PUT"
	contentType := "application/octet-stream"
	contentLength := fmt.Sprintf("%d", len(fileContents))
	blobType := "BlockBlob"
	date := time.Now().UTC().Format(http.TimeFormat)
	url := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, blobName)

	authHeader, err := generateSignature(accountName, key, method, contentLength, contentType, date, blobType, containerName, blobName)
	if err != nil {
		return "", fmt.Errorf("error generating authorization signature: %v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(fileContents))
	if err != nil {
		return "", fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", contentLength)
	req.Header.Set("x-ms-blob-type", blobType)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2019-12-12")
	req.Header.Set("Authorization", authHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making HTTP request: \033[33m%v\033[0m", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to upload blob, status code: \033[33m%d\033[0m, body: \033[33m%s\033[0m", resp.StatusCode, string(body))
	}

	fmt.Println("✅ Blob uploaded successfully to \033[33m" + accountName + "\033[0m, blob name \033[33m" + blobName + "\033[0m.")
	fmt.Printf("💠 The file size uploaded to Azure Storage: \033[33m%.2f\033[0m KB\n", float64(len(fileContents))/1024)

	if _, err := os.Stat("terraform.tfstate"); err == nil {
		if err := os.Remove("terraform.tfstate"); err != nil {
			return "", fmt.Errorf("❌ Error: Failed to delete the terraform.tfstate file: %v", err)
		}
	}

	if err := os.Remove(encodedStateFilePath); err != nil {
		return "", fmt.Errorf("❌ Error: Failed to delete the /tmp/.encoded_wrap file: %v", err)
	}

	storageLocation := fmt.Sprintf("azure_storage:%s/%s/%s", accountName, containerName, blobName)
	return storageLocation, nil
}

func ensureKVPathExists(client *vault.Client, mountPath string, path string) error {
	var secret *vault.Secret
	checkKVVersion, err := client.Logical().Read("sys/mounts/" + mountPath)
	if err != nil {
		return fmt.Errorf("❌ Error:  Unable to determine KV version of secrets engine at: %s", mountPath)
	}

	if checkKVVersion.Data["options"] != nil {
		// KV v2
		secret, err = client.Logical().List(mountPath + "/metadata/" + path)
	} else {
		// KV v1
		secret, err = client.Logical().Read(mountPath + "/data/" + path)
	}
	if err != nil {
		if respErr, ok := err.(*vault.ResponseError); ok && respErr.StatusCode == 403 {
			return fmt.Errorf("❌ Error: Permission denied for path: %s, %w", path, err)
		}
		return fmt.Errorf("❌ Error reading path: %s, %w", path, err)
	}

	if secret != nil {
		fmt.Printf("✅ Writing to path %s\n", path)
		return nil
	}

	_, err = client.Logical().Write(mountPath+"/data/"+path, map[string]interface{}{
		"data": map[string]interface{}{
			"type": "kv",
		},
	})
	if err != nil {
		return fmt.Errorf("❌ Error creating path: %s, %w", path, err)
	}

	fmt.Printf("✅ Path %s created successfully\n", path)

	return nil
}
