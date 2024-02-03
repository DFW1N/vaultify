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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
    "net/http"
	"io"
	"github.com/Azure/azure-sdk-for-go/storage"
    "crypto/hmac"
    "crypto/sha256"
	"encoding/base64"
)

// ###############################
// # checkVaultifySetup Function #
// ###############################

func checkVaultifySetup() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("❌ Error getting user home directory: \033[33m%v\033[0m", err)
	}

	vaultifyDir := filepath.Join(homeDir, ".vaultify")
	settingsFilePath := filepath.Join(vaultifyDir, "settings.json")

	if _, err := os.Stat(vaultifyDir); os.IsNotExist(err) {
		return fmt.Errorf("❌ Error: \033[33m.vaultify\033[0m directory not found")
	}

	if _, err := os.Stat(settingsFilePath); os.IsNotExist(err) {
		return fmt.Errorf("❌ Error: \033[33msettings.json\033[0m file not found in .vaultify directory")
	}

	return nil
}

// ################################
// # getCurrentWorkspace Function #
// ################################

func getCurrentWorkspace() (string, error) {
	cmd := exec.Command("terraform", "workspace", "show")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// #########################
// # readSettings Function #
// #########################

func readSettings() (*Configuration, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return nil, fmt.Errorf("❌ Error getting user home directory: %v", err)
    }

    settingsFilePath := filepath.Join(homeDir, ".vaultify", "settings.json")
    content, err := os.ReadFile(settingsFilePath)
    if err != nil {
        return nil, fmt.Errorf("❌ Error reading settings file: \033[33m%v\033[0m", err)
    }

    var config Configuration
    err = json.Unmarshal(content, &config)
    if err != nil {
        return nil, fmt.Errorf("❌ Error unmarshalling settings JSON: \033[33m%v\033[0m", err)
    }

    return &config, nil
}

// #####################################
// # logic to be used within functions #
// #####################################

// // Detecting the shell configuration file
// shellRC, shellType, err := detectShellRC()
// if err != nil {
// 	fmt.Printf("❌Error detecting shell: %v\n", err)
// 	os.Exit(1)
// }

// fmt.Printf("Detected shell type: %s\n", shellType)

// // Setting environment variable in the detected shell configuration file
// if err := setEnvInShellRC("VAULTIFY_TFWORKSPACE", "true", shellRC); err != nil {
// 	fmt.Printf("❌ Error setting \033[33mVAULTIFY_TFWORKSPACE\033[0m: %v\n", err)
// 	os.Exit(1)
// }

// // Inform the user
// fmt.Printf("Please source your shell configuration file to apply the changes:\n")
// fmt.Printf("source \033[33m%s\033[0m\n", shellRC)

// // Setting environment variable for the current process
// os.Setenv("VAULTIFY_TFWORKSPACE", "true")

// ##########################
// # detectShellRC Function #
// ##########################

// func detectShellRC() (string, string, error) {
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		return "", "", err
// 	}

// 	shell := os.Getenv("SHELL")
// 	shellType := "bash"
// 	shellRC := filepath.Join(homeDir, ".bashrc")
// 	if strings.Contains(shell, "zsh") {
// 		shellType = "zsh"
// 		shellRC = filepath.Join(homeDir, ".zshrc")
// 	}

// 	fmt.Println("Detected shell configuration file:", shellRC)
// 	return shellRC, shellType, nil
// }

// ############################
// # setEnvInShellRC Function #
// ############################

// func setEnvInShellRC(variable, value, shellRCPath string) error {
// 	// Check if the variable already exists in the file
// 	if err := checkVariableInFile(shellRCPath, variable); err != nil {
// 		fmt.Println("Environment variable already set in", shellRCPath)
// 		return nil
// 	}

// 	return appendToFile(shellRCPath, fmt.Sprintf("\nexport %s=\"%s\"\n", variable, value))
// }

// #########################
// # appendToFile Function #
// #########################

// func appendToFile(filePath, content string) error {
// 	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
// 	if err != nil {
// 		fmt.Println("❌ Error opening file:", err)
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(content)
// 	if err != nil {
// 		fmt.Println("❌ Error writing to file:", err)
// 	}
// 	return err
// }

// ################################
// # checkVariableInFile Function #
// ################################

// func checkVariableInFile(filePath, variable string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		if strings.Contains(scanner.Text(), variable) {
// 			return fmt.Errorf("variable \033[33m%s\033[0m already set", variable)
// 		}
// 	}
// 	return nil
// }

// ###################
// # Docker Function #
// ###################

func volumeExists(volumeName string) bool {
	cmd := exec.Command("docker", "volume", "ls", "-q", "-f", "name="+volumeName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to query Docker volumes: %v", err)
	}
	return strings.TrimSpace(out.String()) == volumeName
}

func networkExists(networkName string) bool {
	cmd := exec.Command("docker", "network", "ls", "--format", "{{.Name}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to query Docker networks: %v", err)
	}

	networks := strings.Split(strings.TrimSpace(out.String()), "\n")
	for _, name := range networks {
		if name == networkName {
			return true
		}
	}
	return false
}

func containerIsRunning(containerName string) bool {
	cmd := exec.Command("docker", "ps", "-q", "-f", "name="+containerName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to query running Docker containers: \033[33m%v\033[0m", err)
	}
	return strings.TrimSpace(out.String()) != ""
}

func deleteVolume(volumeName string) {
	if volumeExists(volumeName) {
		// Remove the volume
		rmVolumeCmd := exec.Command("docker", "volume", "rm", volumeName)
		if err := rmVolumeCmd.Run(); err != nil {
			log.Fatalf("Failed to remove Docker volume \033[33m'%s'\033[0m: \033[33m%v\033[0m", volumeName, err)
		} else {
			log.Printf("Docker volume \033[33m'%s'\033[0m removed successfully.", volumeName)
		}
	} else {
		log.Printf("\033[33mNo\033[0m Docker volume with the name \033[33m'%s'\033[0m found. Skipping volume removal.", volumeName)
	}
}

// ####################
// # Azure  Functions #
// ####################

type OAuthResponse struct {
    AccessToken string `json:"access_token"`
}


type Subscription struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type SubscriptionsResponse struct {
	Subscriptions []Subscription `json:"value"`
}


func AuthenticateWithAzureAD() (string, error) {
	tenantID := os.Getenv("ARM_TENANT_ID")
	clientID := os.Getenv("ARM_CLIENT_ID")
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	resource := "https://management.azure.com/"

	data := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&resource=%s",
		clientID, clientSecret, resource)

	url := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", tenantID)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("authentication failed: status code %d", resp.StatusCode)
	}

	var oauthResp OAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&oauthResp); err != nil {
		return "", err
	}

	return oauthResp.AccessToken, nil
}

func checkAzureStorageAccountExists() (bool, error) {
    accessToken, err := AuthenticateWithAzureAD()
    if err != nil {
        return false, fmt.Errorf("error obtaining access token: %v", err)
    }

    config, err := readConfiguration()
    if err != nil {
        return false, fmt.Errorf("error reading configuration: %v", err)
    }

    accountName := config.Settings.Azure.StorageAccountName
    resourceGroup := config.Settings.Azure.StorageAccountResourceGroupName
    subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")

    if subscriptionID == "" {
        return false, fmt.Errorf("subscription ID is missing")
    }

    url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2019-06-01", subscriptionID, resourceGroup, accountName)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return false, err
    }

    req.Header.Set("Authorization", "Bearer "+accessToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        return true, nil
    } else {
        bodyBytes, _ := io.ReadAll(resp.Body)
        return false, fmt.Errorf("storage account check failed with status %d: %s", resp.StatusCode, string(bodyBytes))
    }
}

func CheckAzureEnvVars() error {
    config, err := readConfiguration()
    if err != nil {
        return fmt.Errorf("error reading configuration: %v", err)
    }

    if config.Settings.DefaultSecretStorage == "azure_storage" {
        requiredEnvVars := []string{
            "ARM_SUBSCRIPTION_ID",
            "ARM_CLIENT_ID",
            "ARM_CLIENT_SECRET",
            "ARM_TENANT_ID",
        }

        var missingVars []string
        for _, envVar := range requiredEnvVars {
            if os.Getenv(envVar) == "" {
                missingVars = append(missingVars, envVar)
            }
        }

        if len(missingVars) > 0 {
            return fmt.Errorf("missing required environment variables for Azure storage: \033[33m%v\033[0m", missingVars)
        }
    }

    return nil
}

// ####################################################
// # Create Vaultify Container inside Storage Account #
// ####################################################

func createContainer(accountName, key string) {
    containerName := "vaultify"

    client, err := storage.NewBasicClient(accountName, key)
    if err != nil {
        fmt.Println("Error creating storage client:\033[33m", err)
        return
    }

    blobClient := client.GetBlobService()

    container := blobClient.GetContainerReference(containerName)

    exists, err := container.Exists()
    if err != nil {
        fmt.Println("Error checking container existence:\033[33m", err)
        return
    }

    if exists {

    } else {
        err := container.Create(nil)
        if err != nil {
            fmt.Println("Failed to create container:\033[33m", err)
            return
        }

        fmt.Println("Container \033[33m'vaultify'\033[0m created successfully.")
    }
}

func generateSignature(accountName, accountKey, method, contentLength, contentType, date, blobType, containerName, blobName string) (string, error) {
    urlPath := fmt.Sprintf("/%s/%s/%s", accountName, containerName, blobName)

    stringToSign := method + "\n"

    if method == "PUT" {
        stringToSign += "\n\n" + contentLength + "\n\n" + contentType + "\n\n\n\n\n\n\n"
    } else {
        stringToSign += "\n\n\n\n\n\n\n\n\n\n\n"
    }

    if method == "PUT" { 
        stringToSign += "x-ms-blob-type:" + blobType + "\n"
    }
    stringToSign += "x-ms-date:" + date + "\n" + "x-ms-version:2019-12-12\n" + urlPath

    key, err := base64.StdEncoding.DecodeString(accountKey)
    if err != nil {
        return "", fmt.Errorf("error decoding storage account access key: %v", err)
    }

    hasher := hmac.New(sha256.New, key)
    hasher.Write([]byte(stringToSign))
    signature := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

    authHeader := fmt.Sprintf("SharedKey %s:%s", accountName, signature)
    return authHeader, nil
}