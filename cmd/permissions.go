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
	"os/exec"
)

// TODO: Add a case switch statement depending on, the default secret storage type.
// TODO: Ensure, to check azure role assignment permissions also.

func TokenPermissions() {

	if err := checkVaultifySetup(); err != nil {
		fmt.Println(err)
		fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
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

	checkCmd := exec.Command("curl", "-s", "--header", "X-Vault-Token: "+vaultToken, vaultAddr+"/v1/auth/token/lookup-self")

	output, err := checkCmd.Output()
	if err != nil {
		fmt.Printf("❌ Error checking token permissions: \033[33m%v\033[0m\n", err)
		os.Exit(1)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(output, &response); err != nil {
		fmt.Printf("❌ Error parsing JSON response: \033[33m%v\033[0m\n", err)
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("\033[33mToken Permissions\033[0m:")
	fmt.Println("-----------------------------")
	fmt.Printf("Policies: \033[33m%v\033[0m\n", response["data"].(map[string]interface{})["policies"])

	testPath := vaultAddr + "/v1/" + engineName + "/data/vaultify_test_permission"
	testCmd := exec.Command("curl", "-s", "--header", "X-Vault-Token: "+vaultToken, "--request", "POST", "--data", "{\"data\": {\"test\": \"value\"}}", testPath)

	testOutput, err := testCmd.Output()
	if err != nil {
		fmt.Println("❌ Error testing write \033[33mpermissions\033[0m"+engineName+" engine:", err)
		return
	}

	var testResponse map[string]interface{}
	if err := json.Unmarshal(testOutput, &testResponse); err != nil {
		fmt.Printf("❌ Error parsing test response: %v\n", err)
		return
	}

	if testResponse["errors"] != nil {
		fmt.Println("❌ Token does \033[33mnot\033[0m have permission to create secrets in" + engineName + "engine.")
	} else {
		fmt.Println("✅ Token has \033[33mpermission\033[0m to create secrets in " + engineName + " engine.")
	}

	deleteCmd := exec.Command("curl", "-s", "--header", "X-Vault-Token: "+vaultToken, "--request", "DELETE", testPath)
	if _, err := deleteCmd.Output(); err != nil {
		fmt.Println("❌ Failed to \033[33mclean up\033[0m test secret.")
	}
}
