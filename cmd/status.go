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
	"encoding/json"
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func Status() {

	config, err := readConfiguration()
	if err != nil {
		fmt.Println("❌ \033[33mError\033[0m loading configuration:", err)
		return
	}

	switch config.Settings.DefaultSecretStorage {
	case "vault":
		checkVaultStatus()
	case "s3":
		fmt.Println("⚠️ \033[33m AWS S3 Bucket\033[0m is currently under development.")
	case "azure_storage":
		err = CheckAzureEnvVars()
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}
		fmt.Println("✅ \033[33mAzure Storage\033[0m environment variables are set correctly.")

		_, err = AuthenticateWithAzureAD()
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}
		fmt.Println("✅ \033[33mAuthenticated\033[0m to Azure.")
		exists, err := checkAzureStorageAccountExists()
		if err != nil {
			fmt.Println("❌ Error checking \033[33mAzure storage account\033[0m:\033[33m", err)
			fmt.Println("⚠️  \033[0mPlease validate your \033[33mresource group name\033[0m and \033[33mstorage account name\033[0m in your \033[33m~/.vaultify/settings.json\033[0m file.")
			fmt.Println("\033[33m--------------------------\033[0m")
			azureSettings, _ := json.MarshalIndent(config.Settings.Azure, "", "  ")
			fmt.Println(string(azureSettings))
			fmt.Println("\033[33m--------------------------\033[0m")
		} else if exists {
			fmt.Println("✅ \033[0mAzure storage account\033[33m " + config.Settings.Azure.StorageAccountName + "\033[0m exists.")
		} else {
			fmt.Println("❌ \033[0mAzure storage account\033[33m " + config.Settings.Azure.StorageAccountName + "\033[0m does not exist.") // TODO: If storage account give input prompt for vaultify to create it for you automatically and stores its values dynamically for you in the settings.json file.
		}
	default:
		fmt.Println("Unknown DefaultSecretStorage setting:", config.Settings.DefaultSecretStorage)
	}
}

func checkVaultStatus() {
	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		fmt.Println("❌ Error: \033[33mVAULT_TOKEN\033[0m environment variable is not set. Please authenticate to Vault.")
		return
	}

	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		fmt.Println("❌ Error: \033[33mVAULT_ADDR\033[0m environment variable is not set. Please specify the Vault address.")
		return
	}

	vaultConfig := vault.DefaultConfig()
	vaultConfig.Address = vaultAddr

	vaultClient, err := vault.NewClient(vaultConfig)
	if err != nil {
		fmt.Printf("❌ Error: Unable to connect to %s: %s", vaultAddr, err.Error())
	}

	vaultClient.SetToken(vaultToken)

	init, err := vaultClient.Sys().InitStatus()
	if err != nil {
		fmt.Printf("❌ Error: Unable to authenticate with Vault: %s", err.Error())
	}

	if !init {
		fmt.Println("❌ Error: Vault is not initialized!")
	} else {
		fmt.Println("✅ \033[33mVaultify\033[0m is authenticated and connected to Vault at:\033[33m", vaultAddr)
	}
}

func initVaultClientWithStatus() (vaultClient *vault.Client, initStat bool) {
	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		fmt.Println("❌ Error: \033[33mVAULT_TOKEN\033[0m environment variable is not set. Please authenticate to Vault.")
		return
	}

	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		fmt.Println("❌ Error: \033[33mVAULT_ADDR\033[0m environment variable is not set. Please specify the Vault address.")
		return
	}

	vaultConfig := vault.DefaultConfig()
	vaultConfig.Address = vaultAddr

	vaultClient, err := vault.NewClient(vaultConfig)

	if err != nil {
		fmt.Printf("❌ Error: Unable to connect to %s: %s", vaultAddr, err.Error())
	}

	vaultClient.SetToken(vaultToken)

	initStat, err = vaultClient.Sys().InitStatus()

	if err != nil {
		fmt.Printf("❌ Error: Unable to authenticate with Vault: %s", err.Error())
	}

	return

}
