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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	switch defaultSecretStorage {
	case "vault":
		pushToVault()
	case "azure_storage":
		key, err := listStorageAccountKeys()
		if err != nil {
			log.Fatalf("Failed to list storage account keys: \033[33m%v\033[0m", err)
		}
		accountName := accountName
		// fmt.Printf("Account Name: %s\n", accountName)
		// fmt.Printf("Storage Account Key: %s\n", key)
		createContainer(accountName, key)
		encodedStateFilePath := "/tmp/.encoded_wrap"
		if err := uploadBlobWithAccessKey(accountName, key, encodedStateFilePath); err != nil {
			log.Fatalf("Error uploading blob: \033[33m%v\033[0m", err)
		}

	case "s3":
		fmt.Println("⚠️ \033[33m AWS S3 Bucket\033[0m is currently under development.")
	default:
		fmt.Println("Unsupported secret storage specified.")
	}
}

func pushToVault() {

	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
		return
	}

	vaultClient, initStat := initVaultClientWithStatus()
	if !initStat {
		fmt.Println("❌ Error: Vault is not initialized!")
		return
	}

	settings, err := readSettings()
	if err != nil {
		fmt.Println("❌ Error reading settings:", err)
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
	encodedPayload := encodedStateFile

	engineName := settings.Settings.DefaultEngineName
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

	err = ensureKVPathExists(vaultClient, engineName, dataPath)
	if err != nil {
		fmt.Println("❌ Error: Unable to perform operation", err)
	}

	secretData := map[string]interface{}{
		"data": map[string]interface{}{
			secretPath: encodedPayload,
		},
	}

	_, err = vaultClient.Logical().Write(engineName+"/data/"+secretPath, secretData)
	if err != nil {
		fmt.Println("❌ Error pushing secret to Vault:", err)
		return
	}

	fmt.Printf("✅ Secret written to HashiCorp Vault under: \033[33m%s\033[0m\n", secretPath)
	fmt.Printf("💠 The file size uploaded to Hashicorp Vault: \033[33m%.2f\033[0m KB\n", float64(len(encodedStateFile))/1024)

	if _, err := os.Stat("terraform.tfstate"); err == nil {
		if err := os.Remove("terraform.tfstate"); err != nil {
			fmt.Println("❌ Error: Failed to delete the \033[33mterraform.tfstate\033[0m file.", err)
			return
		}
	}

	if err := os.Remove(encodedStateFilePath); err != nil {
		fmt.Println("❌ Error: Failed to delete the \033[33m/tmp/.encoded_wrap\033[0m file.", err)
		return
	}
}

func isValidBase64(input string) bool {
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}

func listStorageAccountKeys() (string, error) {
	accessToken, err := AuthenticateWithAzureAD()
	if err != nil {
		return "", err
	}

	config, err := readConfiguration()
	if err != nil {
		return "", fmt.Errorf("error loading configuration: %v", err)
	}

	accountName := config.Settings.Azure.StorageAccountName
	resourceGroupName := config.Settings.Azure.StorageAccountResourceGroupName
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")

	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/listKeys?api-version=2019-06-01", subscriptionId, resourceGroupName, accountName)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to list storage account keys, status code: %d", resp.StatusCode)
	}

	var result struct {
		Keys []struct {
			Value string `json:"value"`
		} `json:"keys"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Keys) > 0 {
		return result.Keys[0].Value, nil
	}

	return "", fmt.Errorf("no keys found for the storage account")
}

func uploadBlobWithAccessKey(accountName, key, encodedStateFilePath string) error {
	containerName := "vaultify"

	if err := checkVaultifySetup(); err != nil {
		errMsg := fmt.Sprintf("%v\nPlease run 'vaultify init' to set up Vaultify.", err)
		return fmt.Errorf(errMsg)
	}

	if _, err := os.Stat(encodedStateFilePath); os.IsNotExist(err) {
		return fmt.Errorf("❌ Error: .encoded_wrap file not found in the /tmp directory. Please run 'vaultify wrap' to create the .encoded_wrap file")
	} else if err != nil {
		return fmt.Errorf("❌ Error checking .encoded_wrap file: %v", err)
	}

	encodedStateFileContents, err := os.ReadFile(encodedStateFilePath)
	if err != nil {
		return fmt.Errorf("❌ Error reading .encoded_wrap file: %v", err)
	}

	if len(encodedStateFileContents) == 0 {
		return fmt.Errorf("❌ Error: .encoded_wrap file is empty")
	}

	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		return fmt.Errorf("❌ Error getting current Terraform workspace: %v", err)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("❌ Error getting current working directory: %v", err)
	}

	workingDirName := filepath.Base(workingDir)
	blobName := fmt.Sprintf("%s/%s_%s", workingDirName, workspaceName, "terraform.tfstate")

	file, err := os.Open(encodedStateFilePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file contents: %v", err)
	}

	method := "PUT"
	contentType := "application/octet-stream"
	contentLength := fmt.Sprintf("%d", len(fileContents))
	blobType := "BlockBlob"
	date := time.Now().UTC().Format(http.TimeFormat)
	url := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, blobName)

	authHeader, err := generateSignature(accountName, key, method, contentLength, contentType, date, blobType, containerName, blobName)
	if err != nil {
		return fmt.Errorf("error generating authorization signature: %v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(fileContents))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", contentLength)
	req.Header.Set("x-ms-blob-type", blobType)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2019-12-12")
	req.Header.Set("Authorization", authHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: \033[33m%v\033[0m", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to upload blob, status code: \033[33m%d\033[0m, body: \033[33m%s\033[0m", resp.StatusCode, string(body))
	}

	fmt.Println("✅ Blob uploaded successfully to \033[33m" + accountName + "\033[0m, blob name \033[33m" + blobName + "\033[0m.")
	fmt.Printf("💠 The file size uploaded to Azure Storage: \033[33m%.2f\033[0m KB\n", float64(len(fileContents))/1024)

	if _, err := os.Stat("terraform.tfstate"); err == nil {
		if err := os.Remove("terraform.tfstate"); err != nil {
			return fmt.Errorf("❌ Error: Failed to delete the terraform.tfstate file: %v", err)
		}
	}

	if err := os.Remove(encodedStateFilePath); err != nil {
		return fmt.Errorf("❌ Error: Failed to delete the /tmp/.encoded_wrap file: %v", err)
	}
	return nil
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
