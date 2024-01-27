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

// Configuration structure to hold settings
type Configuration struct {
	VaultAddress string
}

// Configure command implementation
func Configure() {

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

	// Initialize the configuration with values from environment variables
	config := Configuration{
		VaultAddress: vaultAddr,
	}

	// Present the configuration options in a table format
	fmt.Println("\nVaultify Configuration Options:\n")
	fmt.Printf("%-20s %s\n", "Option", "Value")
	fmt.Println("-------------------------------------------")
	fmt.Printf("%-20s %s\n", "1. Vault Address", config.VaultAddress)
	fmt.Println()

	fmt.Println("Enter the number of the option you want to change (or 0 to exit):\n")

	// Read user input
	var choice int
	for {
		fmt.Print("Choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Handle user choices
		switch choice {
		case 0:
			// Exit
			fmt.Println("\nExiting configuration.")
			return
		case 1:
			// Change Vault Address
			fmt.Print("\nEnter new Vault Address: ")
			fmt.Scanln(&config.VaultAddress)
			fmt.Println("\nVault Address updated successfully.")
		default:
			fmt.Println("Invalid option. Please enter a valid number.")
		}
	}
}

func main() {
	Configure()
}
