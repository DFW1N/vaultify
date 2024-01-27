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
	"os/exec"
	"strings"
)

// Status command implementation
func Status() {
	// Check if the VAULT_TOKEN environment variable is set
	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		fmt.Println("❌ Error: VAULT_TOKEN environment variable is not set. Please authenticate to Vault.")
		return
	}

	// Check if the VAULT_ADDR environment variable is set
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		fmt.Println("❌ Error: VAULT_ADDR environment variable is not set. Please specify the Vault address.")
		return
	}

	// Use the 'curl' command to perform an authenticated operation
	curlCommand := "curl"

	// Replace the following line with your actual Vault API endpoint or operation
	vaultAPIEndpoint := vaultAddr + "/v1/sys/init"

	// Execute the 'curl' command to check if Vaultify is authenticated
	curlCmd := exec.Command(
		curlCommand,
		"--header", "X-Vault-Token: "+vaultToken,
		"--request", "GET",
		vaultAPIEndpoint,
	)

	curlOutput, err := curlCmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Error executing 'curl' command:", err)
		return
	}

	// Check the response from the 'curl' command to determine authentication status
	if strings.Contains(string(curlOutput), "initialized\":true") {
		fmt.Println("✅ Vaultify is authenticated and connected to Vault at:", vaultAddr)
	} else {
		fmt.Println("❌ Error: Vaultify is not authenticated or unable to connect to Vault.")
	}
}
