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

	switch defaultSecretStorage {
	case "vault":
		pullFromVault()
	case "azure_storage":
		key, err := listStorageAccountKeys()
		if err != nil {
			log.Fatalf("Failed to list storage account keys: \033[33m%v\033[0m", err)
		}
		if err := pullBlobFromAzureStorage(accountName, key); err != nil {
			fmt.Printf("Error pulling blob from Azure Storage: %v\n", err)
		}
	case "s3":
		log.Println("AWS S3 pulling is currently under development.")
	default:
		log.Println("Unsupported secret storage specified.")
	}
}

func pullFromVault() {

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

	engineName := settings.Settings.DefaultEngineName
	dataPath := "vaultify"

	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		fmt.Println("❌ Error getting current \033[33mTerraform\033[0m workspace:", err)
		return
	}

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Error getting current working directory:", err)
		return
	}

	workingDirName := filepath.Base(workingDir)

	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workingDirName, workspaceName, "terraform.tfstate")

	secretValue, err := vaultClient.KVv2(engineName).Get(context.Background(), secretPath)
	if err != nil {
		fmt.Printf("❌ Error: %v, %s\n", vault.ErrSecretNotFound, err)
		return
	}

	if secretValue == nil {
		fmt.Printf("❌ Error: No secrets at %s, %s\n", secretPath, err)
		return
	}

	fmt.Println("✅ Secret exists in Vault. Retrieving...")

	base64String, ok := secretValue.Data[secretPath].(string)
	if !ok {
		fmt.Println("❌ Error: Specific \033[33mkey\033[0m not found in the data")
	}

	targetFilePath := "terraform.tfstate.gz.b64"
	if _, err := os.Stat(targetFilePath); err == nil {
		fmt.Println("❌ Error: File \033[33mterraform.tfstate.gz.b64\033[0m already exists in the directory.")
		return
	} else if !os.IsNotExist(err) {
		fmt.Printf("❌ Error checking if file exists: %v\n", err)
		return
	}

	if err := saveStateToFile([]byte(base64String), targetFilePath); err != nil {
		fmt.Println("❌ Error saving base64 string to file:\033[33m", err)
		return
	}

	fmt.Println("✅ Secret retrieved and saved as \033[33mterraform.tfstate.gz.b64\033[0m")
}

func pullBlobFromAzureStorage(accountName, key string) error {
	containerName := "vaultify"

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

	method := "GET"
	date := time.Now().UTC().Format(http.TimeFormat)
	url := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, blobName)

	authHeader, err := generateSignature(accountName, key, method, "0", "", date, "", containerName, blobName)
	if err != nil {
		return fmt.Errorf("❌ Error generating authorization signature for download: \033[33m%v\033[0m", err)
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("❌ Error creating HTTP request for download: \033[33m%v\033[0m", err)
	}

	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2019-12-12")
	req.Header.Set("Authorization", authHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("❌ Error making HTTP request for download: \033[33m%v\033[0m", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("❌ Failed to download blob, status code: \033[33m%d\033[0m, response: \033[33m%s\033[0m", resp.StatusCode, string(responseBody))
	}

	outputFile, err := os.Create("terraform.tfstate.gz.b64")
	if err != nil {
		return fmt.Errorf("❌ Error creating file to save downloaded blob: \033[33m%v\033[0m", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, resp.Body)
	if err != nil {
		return fmt.Errorf("❌ Error writing downloaded blob to file: \033[33m%v\033[0m", err)
	}

	fmt.Println("✅ Blob downloaded successfully and saved as \033[33mterraform.tfstate.gz.b64\033[0m")
	return nil
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
