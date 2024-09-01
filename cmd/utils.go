package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

/////////////
// General //
/////////////

func parseDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", date)
}

func isValidBase64(input string) bool {
	_, err := base64.StdEncoding.DecodeString(input)
	return err == nil
}

////////////////////////////
// Azure Storage Account //
///////////////////////////

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
