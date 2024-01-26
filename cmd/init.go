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

// File: cmd/init.go
package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

// Init command implementation
func Init() {
	// Retrieve environment variables
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddr := os.Getenv("VAULT_ADDR")

	// Check if environment variables are set
	if vaultToken == "" || vaultAddr == "" {
		fmt.Println("Error: VAULT_TOKEN and VAULT_ADDR environment variables must be set.")
		os.Exit(1)
	}

	// Add your init command logic here
	if err := initializeVaultify(vaultToken, vaultAddr); err != nil {
		fmt.Printf("Error initializing Vaultify: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Vaultify initialized successfully.")
	}
}

func initializeVaultify(token, addr string) error {
	// Construct the curl command to authenticate against HashiCorp Vault
	cmd := exec.Command("curl", "--header", "X-Vault-Token:"+token, addr+"/v1/sys/health")

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to authenticate against Vault: %v", err)
	}

	// If the command runs successfully, it means authentication was successful
	return nil
}
