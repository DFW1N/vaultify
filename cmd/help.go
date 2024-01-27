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
	fmt.Println("Vaultify - A CLI tool for managing statefiles.")
	fmt.Println("\nCommands:")
	fmt.Println("  init      			Initialize Vaultify in your operating system")
	fmt.Println("  validate      		Vaultify will validate your terraform.tfstate file json")
	fmt.Println("  compare      			Vaultify will compare your local terraform.tfstate file json to your remote vault terraform.tfstate file")
	fmt.Println("  update    			Update Vaultify")
	fmt.Println("  wrap      			Wrap a secret in base64")
	fmt.Println("  unwrap    			Unwrap a secret from base64")
	fmt.Println("  pull      			Pull state from remote Hashicorp Vault server")
	fmt.Println("  push      			Push state to remote Hashicorp Vault server afer you have wrapped your statefile")
	fmt.Println("  status      			Checks if Vaultify is still authenticated to Hashicorp Vault.")
	fmt.Println("  configure      		Configures the Vaultify project, allowing customization of settings such as the Vault address, authentication method, and data paths")
	fmt.Println("  -v, --version     		Show the Vaultify version")
	fmt.Println("  -h, --help        		Show this help message")
}
