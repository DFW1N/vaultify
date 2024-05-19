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
	"fmt"
	"os"
)

// TODO: Add a case switch statement depending on, the default secret storage type.
// TODO: Ensure, to check azure role assignment permissions also.

func TokenPermissions() {

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

	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultToken == "" || vaultAddr == "" {
		fmt.Println("❌ Error: \033[33mVAULT_TOKEN\033[0m and \033[33mVAULT_ADDR\033[0m environment variables must be set.")
		os.Exit(1)
	}

	vaultAuthLookupSelf, err := vaultClient.Auth().Token().LookupSelf()
	if err != nil {
		fmt.Printf("❌ Error checking token permissions: \033[33m%v\033[0m\n", err)
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("\033[33mToken Permissions\033[0m:")
	fmt.Println("-----------------------------")
	fmt.Printf("Policies: \033[33m%v\033[0m\n", vaultAuthLookupSelf.Data["policies"])

	testPath := engineName + "/data/vaultify_test_permission"

	_, err = vaultClient.Logical().Write(testPath, map[string]interface{}{
		"data": map[string]interface{}{
			"data": "test",
		},
	})
	if err != nil {
		fmt.Println("❌ Error testing write \033[33mpermissions\033[0m"+engineName+" engine: ", err)
		return
	}

	fmt.Println("✅ Token has \033[33mpermission\033[0m to create secrets in " + engineName + " engine.")

	_, err = vaultClient.Logical().Delete(testPath)
	if err != nil {
		fmt.Println("❌ Failed to \033[33mclean up\033[0m test secret.")
	}
	fmt.Println("✅ \033[33mclean up\033[0m of test secrets complete.")
}
