// ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Get your secret from vault easily with a single command.                #
// ########################################################################################

package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	vault "github.com/hashicorp/vault/api"
	"gopkg.in/yaml.v2"
)

func Get(args []string) {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	path := getCmd.String("path", "", "Path to retrieve the secret from")
	key := getCmd.String("key", "", "Name of the secret (will be used as the path if no path is provided)")
	jsonOutput := getCmd.Bool("json", false, "Output in JSON format")
	yamlOutput := getCmd.Bool("yaml", false, "Output in YAML format")
	err := getCmd.Parse(args)
	if err != nil {
		fmt.Println("❌ Error parsing get flags:", err)
		return
	}

	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
		return
	}

	config, err := readConfiguration()
	if err != nil {
		fmt.Println("❌ Error reading configuration:", err)
		return
	}

	defaultSecretStorage := config.Settings.DefaultSecretStorage

	var secretValue string
	var metadata map[string]string
	var fullPath string

	switch defaultSecretStorage {
	case "vault":
		secretValue, metadata, fullPath, err = getFromVault(*path, *key)
	case "azure_storage":
		secretValue, metadata, fullPath, err = getFromAzureStorage(*path, *key)
	default:
		fmt.Println("❌ Unsupported secret storage specified.")
		return
	}

	if err != nil {
		fmt.Printf("❌ Error retrieving secret: %v\n", err)
		return
	}

	if err := LogHistory("get", fullPath); err != nil {
		fmt.Printf("❌ Error logging history: %v\n", err)
	}
	if err := LogUserAction("get", fullPath); err != nil {
		fmt.Printf("❌ Error logging user action: %v\n", err)
	}

	if *jsonOutput {
		outputJSON(fullPath, filepath.Base(*key), secretValue, metadata)
	} else if *yamlOutput {
		outputYAML(fullPath, filepath.Base(*key), secretValue, metadata)
	} else {
		fmt.Printf("%v\n", secretValue)
	}
}

func getFromVault(path, key string) (string, map[string]string, string, error) {
	vaultClient, err := vault.NewClient(nil)
	if err != nil {
		return "", nil, "", fmt.Errorf("error creating Vault client: %v", err)
	}

	config, err := readConfiguration()
	if err != nil {
		return "", nil, "", fmt.Errorf("error reading configuration: %v", err)
	}

	engineName := config.Settings.DefaultEngineName
	secretPath := filepath.Join("secrets", path)
	if key != "" {
		secretPath = filepath.Join(secretPath, key)
	}

	fullPath := fmt.Sprintf("%s/data/%s", engineName, secretPath)

	secret, err := vaultClient.Logical().Read(fullPath)
	if err != nil {
		return "", nil, "", fmt.Errorf("error reading secret from Vault: %v", err)
	}

	if secret == nil || secret.Data == nil {
		return "", nil, "", fmt.Errorf("no secret found at path: %s", fullPath)
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", nil, "", fmt.Errorf("unexpected secret data format")
	}

	encryptedValue, exists := data["value"].(string)
	if !exists {
		return "", nil, "", fmt.Errorf("no value found in secret at path %s", fullPath)
	}

	// Decrypt the value
	decryptedValue, err := decryptSecret(encryptedValue, os.Getenv("VAULTIFY_PASSPHRASE"))
	if err != nil {
		return "", nil, "", fmt.Errorf("error decrypting secret: %v", err)
	}

	metadata := make(map[string]string)
	for k, v := range data {
		if k != "value" {
			metadata[k] = fmt.Sprintf("%v", v)
		}
	}

	return decryptedValue, metadata, fmt.Sprintf("vault:%s", fullPath), nil
}

func getFromAzureStorage(path, key string) (string, map[string]string, string, error) {
	config, err := readConfiguration()
	if err != nil {
		return "", nil, "", fmt.Errorf("error loading configuration: %v", err)
	}

	accountName := config.Settings.Azure.StorageAccountName

	storageAccountKey, err := listStorageAccountKeys()
	if err != nil {
		return "", nil, "", fmt.Errorf("error getting storage account key: %v", err)
	}

	// Construct the full path
	fullPath := path
	if key != "" {
		fullPath = filepath.Join(fullPath, key)
	}

	// Split the path into container name and blob path
	pathParts := strings.SplitN(fullPath, "/", 2)
	containerName := pathParts[0]
	blobName := ""
	if len(pathParts) > 1 {
		blobName = pathParts[1]
	}

	method := "GET"
	date := time.Now().UTC().Format(http.TimeFormat)
	url := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, blobName)

	authHeader, err := generateSignature(accountName, storageAccountKey, method, "", "", date, "", containerName, blobName)
	if err != nil {
		return "", nil, "", fmt.Errorf("error generating authorization signature: %v", err)
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", nil, "", fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("x-ms-date", date)
	req.Header.Set("x-ms-version", "2019-12-12")
	req.Header.Set("Authorization", authHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil, "", fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", nil, "", fmt.Errorf("failed to get blob, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	encryptedValue, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, "", fmt.Errorf("error reading response body: %v", err)
	}

	// Decrypt the value
	decryptedValue, err := decryptSecret(string(encryptedValue), os.Getenv("VAULTIFY_PASSPHRASE"))
	if err != nil {
		return "", nil, "", fmt.Errorf("error decrypting secret: %v", err)
	}

	// Extract metadata from headers
	metadata := make(map[string]string)
	for k, v := range resp.Header {
		if strings.HasPrefix(k, "X-Ms-Meta-") {
			metaKey := strings.TrimPrefix(k, "X-Ms-Meta-")
			metadata[metaKey] = v[0]
		}
	}

	storageLocation := fmt.Sprintf("azure_storage:%s/%s/%s", accountName, containerName, blobName)
	return decryptedValue, metadata, storageLocation, nil
}

func outputJSON(fullPath, secretName, value string, metadata map[string]string) {
	output := map[string]interface{}{
		"path":        fullPath,
		"secret_name": secretName,
		"value":       value,
		"metadata":    metadata,
	}
	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("❌ Error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func outputYAML(fullPath, secretName, value string, metadata map[string]string) {
	output := map[string]interface{}{
		"path":        fullPath,
		"secret_name": secretName,
		"value":       value,
		"metadata":    metadata,
	}
	yamlData, err := yaml.Marshal(output)
	if err != nil {
		fmt.Println("❌ Error marshaling to YAML:", err)
		return
	}
	fmt.Println(string(yamlData))
}
