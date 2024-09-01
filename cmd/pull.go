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
	"context"
	//"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	//"os/exec"
	"path/filepath"
	//"strings"
	"time"

	vault "github.com/hashicorp/vault/api"
)

type VaultResponse struct {
	Data struct {
		Data map[string]string `json:"data"`
	} `json:"data"`
}

func Pull() {

	if err := checkVaultifySetup(); err != nil {
		log.Printf("%v\nPlease run 'vaultify init' to set up Vaultify.\n", err)
		return
	}

	config, err := readConfiguration()
	if err != nil {
		fmt.Println("❌ \033[33mError\033[0m loading configuration:", err)
		return
	}

	defaultSecretStorage := config.Settings.DefaultSecretStorage
	accountName := config.Settings.Azure.StorageAccountName

	var pullLocation string

	switch defaultSecretStorage {
	case "vault":
		var err error
		pullLocation, err = pullFromVault()
		if err != nil {
			fmt.Printf("Error pulling from Vault: %v\n", err)
			return
		}
	case "azure_storage":
		key, err := listStorageAccountKeys()
		if err != nil {
			fmt.Printf("Failed to list storage account keys: \033[33m%v\033[0m\n", err)
			return
		}
		pullLocation, err = pullBlobFromAzureStorage(accountName, key)
		if err != nil {
			fmt.Printf("Failed to pull blob from Azure Storage: \033[33m%v\033[0m\n", err)
			return
		}
	case "s3":
		fmt.Println("⚠️ \033[33m AWS S3 Bucket\033[0m is currently under development.")
	default:
		log.Println("Unsupported secret storage specified.")
	}

	if pullLocation != "" {
		if err := LogHistory("pull", pullLocation); err != nil {
			fmt.Printf("❌ Error logging history: %v\n", err)
		}
		if err := LogUserAction("pull", pullLocation); err != nil {
			fmt.Printf("❌ Error logging user action: %v\n", err)
		}
		fmt.Printf("✅ Successfully pulled state from %s\n", pullLocation)
	}
}

func pullFromVault() (string, error) {

	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
	}

	vaultClient, initStat := initVaultClientWithStatus()
	if !initStat {
		return "", fmt.Errorf("❌ Error: Vault is not initialized!")
	}

	settings, err := readSettings()
	if err != nil {
		return "", fmt.Errorf("❌ Error getting current Terraform workspace: %w", err)
	}

	engineName := settings.Settings.DefaultEngineName
	dataPath := "vaultify"

	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		return "", fmt.Errorf("❌ Error getting current Terraform workspace: %w", err)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("❌ Error getting current working directory: %w", err)
	}

	workingDirName := filepath.Base(workingDir)

	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workingDirName, workspaceName, "terraform.tfstate")
	fullPath := fmt.Sprintf("%s/data/%s", engineName, secretPath)

	secretValue, err := vaultClient.KVv2(engineName).Get(context.Background(), secretPath)
	if err != nil {
		return "", fmt.Errorf("❌ Error: %v, %w", vault.ErrSecretNotFound, err)
	}

	if secretValue == nil {
		return "", fmt.Errorf("❌ Error: No secrets at %s", secretPath)
	}

	fmt.Println("✅ Secret exists in Vault. Retrieving...")

	base64String, ok := secretValue.Data[secretPath].(string)
	if !ok {
		fmt.Println("❌ Error: Specific \033[33mkey\033[0m not found in the data")
	}

	targetFilePath := "terraform.tfstate.gz.enc.b64"
	if _, err := os.Stat(targetFilePath); err == nil {
		return "", fmt.Errorf("❌ Error: File terraform.tfstate.gz.enc.b64 already exists in the directory")
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("❌ Error checking if file exists: %w", err)
	}

	if err := saveStateToFile([]byte(base64String), targetFilePath); err != nil {
		return "", fmt.Errorf("❌ Error saving base64 string to file: %w", err)
	}

	fmt.Println("✅ Secret retrieved and saved as \033[33mterraform.tfstate.gz.enc.b64\033[0m")

	return fmt.Sprintf("vault:%s", fullPath), nil
}

func pullBlobFromAzureStorage(accountName, key string) (string, error) {
	containerName := "vaultify"

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

	method := "GET"
	date := time.Now().UTC().Format(http.TimeFormat)
	url := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, blobName)

	authHeader, err := generateSignature(accountName, key, method, "0", "", date, "", containerName, blobName)
	if err != nil {
		return "", fmt.Errorf("❌ Error generating authorization signature for download: %v", err)
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("❌ Error creating HTTP request for download: %v", err)
	}

	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2019-12-12")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("❌ Error making HTTP request for download: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("❌ Failed to download blob, status code: %d, response: %s", resp.StatusCode, string(responseBody))
	}

	outputFile, err := os.Create("terraform.tfstate.gz.enc.b64")
	if err != nil {
		return "", fmt.Errorf("❌ Error creating file to save downloaded blob: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("❌ Error writing downloaded blob to file: %v", err)
	}

	fmt.Println("✅ Blob downloaded successfully and saved as \033[33mterraform.tfstate.gz.enc.b64\033[0m")

	pullLocation := fmt.Sprintf("azure_storage:%s/%s/%s", accountName, containerName, blobName)
	return pullLocation, nil
}

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
