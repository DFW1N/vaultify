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

import "fmt"

func Help() {
	fmt.Println("\033[33mVaultify\033[0m - A CLI tool for managing statefiles.")
	fmt.Println("\nCommands:")
	fmt.Println("  \033[33minit\033[0m      			Initialize Vaultify in your operating system")
	fmt.Println("  \033[33mvalidate\033[0m      		Vaultify will validate your terraform.tfstate file json")
	fmt.Println("  \033[33mcompare\033[0m      			Vaultify will compare your local terraform.tfstate file json to your remote vault terraform.tfstate file")
	fmt.Println("  \033[33mupdate\033[0m    			Update Vaultify")
	fmt.Println("  \033[33mwrap\033[0m      			Wrap a secret in base64")
	fmt.Println("  \033[33munwrap\033[0m    			Unwrap a secret from base64")
	fmt.Println("  \033[33mdelete\033[0m    			Delete the Hashicorp secret from Vault")
	fmt.Println("  \033[33mpath\033[0m     			Display the Hashicorp secret path used to store statefile")
	fmt.Println("  \033[33mpull\033[0m      			Pull state from remote Hashicorp Vault server")
	fmt.Println("  \033[33minstall-vault\033[0m      			Deploy a local developer Hashicorp Vault server")
	fmt.Println("  \033[33mdelete-vault\033[0m      			Delete a local developer Hashicorp Vault server")
	fmt.Println("  \033[33mpermissions\033[0m      			Validate your roles and permissions on your used Hashicorp Vault token")
	fmt.Println("  \033[33mpush\033[0m      			Push state to remote Hashicorp Vault server afer you have wrapped your statefile")
	fmt.Println("  \033[33mstatus\033[0m      			Checks if Vaultify is still authenticated to Hashicorp Vault.")
	fmt.Println("  \033[33mconfigure\033[0m      		Configures the Vaultify project, allowing customization of settings such as the Vault address, authentication method, and data paths")
	fmt.Println("  \033[33m-v\033[0m, \033[33m--version\033[0m     		Show the Vaultify version")
	fmt.Println("  \033[33m-h\033[0m, \033[33m--help\033[0m        		Show this help message")
}
