// ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Inject a value into your vault, by running a simple command.            #
// ########################################################################################

package cmd

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func Inject(args []string) {
	injectCmd := flag.NewFlagSet("inject", flag.ExitOnError)
	path := injectCmd.String("path", "", "Path to store the secret")
	key := injectCmd.String("key", "", "Name of the secret (will be used as the path if no path is provided)")
	metadataFlag := injectCmd.String("metadata", "", "Metadata in JSON format or path to JSON file")
	secretFileFlag := injectCmd.String("secret-file", "", "Path to file containing the secret value")
	err := injectCmd.Parse(args)
	if err != nil {
		fmt.Println("❌ Error parsing inject flags:", err)
		return
	}

	if *secretFileFlag == "" && injectCmd.NArg() < 1 {
		fmt.Println("Usage: vaultify inject [-path <path>] [-key <secret_name>] [-metadata <json_metadata_or_file>] (-secret-file <file_path> | <secret_value>)")
		return
	}

	var secretValue string
	if *secretFileFlag != "" {
		content, err := os.ReadFile(*secretFileFlag)
		if err != nil {
			fmt.Printf("❌ Error reading secret file: %v\n", err)
			return
		}
		secretValue = string(content)
	} else {
		secretValue = injectCmd.Arg(0)
	}

	var metadata map[string]interface{}
	if *metadataFlag != "" {
		metadata, err = parseMetadata(*metadataFlag)
		if err != nil {
			fmt.Println("❌ Error parsing metadata:", err)
			return
		}
	}

	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
		return
	}

	config, err := readConfiguration()
	if err != nil {
		fmt.Println("❌ \033[33mError\033[0m loading configuration:", err)
		return
	}

	defaultSecretStorage := config.Settings.DefaultSecretStorage

	switch defaultSecretStorage {
	case "vault":
		if err := injectToVault(*path, *key, secretValue, metadata); err != nil {
			fmt.Println("❌ Error injecting secret to Vault:", err)
		}
	case "azure_storage":
		fullPath := *path
		if *key != "" {
			fullPath = filepath.Join(fullPath, *key)
		}
		if err := injectToAzureStorage(fullPath, "", secretValue, metadata); err != nil {
			fmt.Println("❌ Error injecting secret to Azure Storage:", err)
		}
	default:
		fmt.Println("❌ Unsupported secret storage specified.")
	}
}

func parseMetadata(metadataFlag string) (map[string]interface{}, error) {
	var metadata map[string]interface{}

	if _, err := os.Stat(metadataFlag); err == nil {
		metadataBytes, err := os.ReadFile(metadataFlag)
		if err != nil {
			return nil, fmt.Errorf("error reading metadata file: %v", err)
		}
		if err := json.Unmarshal(metadataBytes, &metadata); err != nil {
			return nil, fmt.Errorf("error parsing metadata JSON from file: %v", err)
		}
	} else {
		if err := json.Unmarshal([]byte(metadataFlag), &metadata); err != nil {
			return nil, fmt.Errorf("error parsing inline metadata JSON: %v", err)
		}
	}

	return metadata, nil
}

func injectToVault(path, key, secretValue string, metadata map[string]interface{}) error {
	vaultClient, initStat := initVaultClientWithStatus()
	if !initStat {
		return fmt.Errorf("❌ Error: Vault is not initialized")
	}

	settings, err := readSettings()
	if err != nil {
		return fmt.Errorf("❌ Error reading settings: %v", err)
	}

	engineName := settings.Settings.DefaultEngineName

	secretPath := "secrets"
	if path != "" {
		secretPath = filepath.Join(secretPath, strings.Trim(path, "/"))
	} else if key != "" {
		secretPath = filepath.Join(secretPath, key)
	} else {
		secretPath = filepath.Join(secretPath, "default")
	}

	encryptedValue, err := encryptContents([]byte(secretValue), os.Getenv("VAULTIFY_PASSPHRASE"))
	if err != nil {
		return fmt.Errorf("error encrypting secret: %v", err)
	}

	secretData := map[string]interface{}{
		"data": map[string]interface{}{
			"value": string(encryptedValue),
		},
	}

	fullPath := fmt.Sprintf("%s/data/%s", engineName, secretPath)

	// Write the secret data
	_, err = vaultClient.Logical().Write(fullPath, secretData)
	if err != nil {
		return fmt.Errorf("error writing secret: %v", err)
	}

	if len(metadata) > 0 {
		metadataPath := fmt.Sprintf("%s/metadata/%s", engineName, secretPath)
		metadataUpdate := map[string]interface{}{
			"custom_metadata": metadata,
		}
		_, err = vaultClient.Logical().Write(metadataPath, metadataUpdate)
		if err != nil {
			return fmt.Errorf("error updating metadata: %v", err)
		}
	}

	fmt.Printf("✅ Secret injected to HashiCorp Vault under: \033[33m%s\033[0m\n", fullPath)
	fmt.Printf("💠 Secret Name: \033[33m%s\033[0m\n", filepath.Base(secretPath))
	if len(metadata) > 0 {
		fmt.Println("📋 Custom metadata added:")
		for k, v := range metadata {
			fmt.Printf("   \033[33m%s\033[0m: %v\n", k, v)
		}
	}

	if err := LogHistory("inject", fmt.Sprintf("vault:%s", fullPath)); err != nil {
		fmt.Printf("❌ Error logging history: %v\n", err)
	}

	return nil
}

func injectToAzureStorage(path, key, secretValue string, metadata map[string]interface{}) error {
	config, err := readConfiguration()
	if err != nil {
		return fmt.Errorf("❌ Error loading configuration: %v", err)
	}

	accountName := config.Settings.Azure.StorageAccountName

	storageAccountKey, err := listStorageAccountKeys()
	if err != nil {
		return fmt.Errorf("❌ Error getting storage account key: %v", err)
	}

	// Split the path into container name and blob path
	pathParts := strings.SplitN(path, "/", 2)
	containerName := pathParts[0]
	blobPath := ""
	if len(pathParts) > 1 {
		blobPath = pathParts[1]
	}

	// Create the container if it doesn't exist
	if err := createContainerIfNotExists(accountName, containerName); err != nil {
		return fmt.Errorf("❌ Error creating container: %v", err)
	}

	// Construct the blob name
	blobName := key
	if blobPath != "" {
		blobName = filepath.Join(blobPath, key)
	}

	// Encrypt the secret value
	encryptedValue, err := encryptContents([]byte(secretValue), os.Getenv("VAULTIFY_PASSPHRASE"))
	if err != nil {
		return fmt.Errorf("error encrypting secret: %v", err)
	}

	method := "PUT"
	contentType := "application/octet-stream"
	contentLength := fmt.Sprintf("%d", len(encryptedValue))
	blobType := "BlockBlob"
	date := time.Now().UTC().Format(http.TimeFormat)
	url := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, blobName)

	// Prepare metadata headers
	metadataHeaders := make([]string, 0, len(metadata))
	for k, v := range metadata {
		metadataHeaders = append(metadataHeaders, fmt.Sprintf("x-ms-meta-%s:%v", strings.ToLower(k), v))
	}
	sort.Strings(metadataHeaders)

	authHeader, err := generateSignature(accountName, storageAccountKey, method, contentLength, contentType, date, blobType, containerName, blobName, metadataHeaders)
	if err != nil {
		return fmt.Errorf("error generating authorization signature: %v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(encryptedValue))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", contentLength)
	req.Header.Set("x-ms-blob-type", blobType)
	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2019-12-12")
	req.Header.Set("Authorization", authHeader)

	for _, header := range metadataHeaders {
		parts := strings.SplitN(header, ":", 2)
		if len(parts) == 2 {
			req.Header.Set(parts[0], parts[1])
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: \033[33m%v\033[0m", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to upload blob, status code: \033[33m%d\033[0m, body: \033[33m%s\033[0m", resp.StatusCode, string(body))
	}

	fmt.Printf("✅ Secret injected successfully to \033[33m%s\033[0m, container: \033[33m%s\033[0m, blob name: \033[33m%s\033[0m.\n", accountName, containerName, blobName)
	fmt.Printf("💠 The secret size uploaded to Azure Storage: \033[33m%.2f\033[0m KB\n", float64(len(encryptedValue))/1024)

	storageLocation := fmt.Sprintf("azure_storage:%s/%s/%s", accountName, containerName, blobName)
	if err := LogHistory("inject", storageLocation); err != nil {
		fmt.Printf("❌ Error logging history: %v\n", err)
	}

	return nil
}

func createContainerIfNotExists(accountName, containerName string) error {
	accessToken, err := AuthenticateWithAzureAD()
	if err != nil {
		return fmt.Errorf("error getting Azure AD token: %v", err)
	}

	config, err := readConfiguration()
	if err != nil {
		return fmt.Errorf("error loading configuration: %v", err)
	}

	resourceGroupName := config.Settings.Azure.StorageAccountResourceGroupName
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")

	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/default/containers/%s?api-version=2021-04-01", subscriptionId, resourceGroupName, accountName, containerName)

	req, err := http.NewRequest("PUT", url, strings.NewReader("{}"))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create container, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}
