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

package main

import (
	"flag"
	"fmt"
	"os"
	"vaultify/cmd"
)

func main() {
	// Retrieve environment variables
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddr := os.Getenv("VAULT_ADDR")

	// Check if environment variables are set
	if vaultToken == "" || vaultAddr == "" {
		fmt.Println("Error: VAULT_TOKEN and VAULT_ADDR environment variables must be set.")
		os.Exit(1)
	}

	// Define flags for version and help, allowing both -v and version, -h and help
	var (
		versionFlag     bool
		versionFlagLong bool
		helpFlag        bool
		helpFlagLong    bool
	)

	flag.BoolVar(&versionFlag, "v", false, "Prints the version of the program")
	flag.BoolVar(&versionFlagLong, "version", false, "Prints the version of the program")
	flag.BoolVar(&helpFlag, "h", false, "Prints the help information")
	flag.BoolVar(&helpFlagLong, "help", false, "Prints the help information")

	flag.Parse()

	switch {
	case versionFlag || versionFlagLong:
		cmd.Version()
	case helpFlag || helpFlagLong:
		cmd.Help()
	default:
		if len(os.Args) < 2 {
			fmt.Println("Usage: vaultify [command]")
			fmt.Println("Use 'vaultify -h' for help.")
			return
		}

		switch os.Args[1] {
		case "init":
			cmd.Init()
		case "validate":
			cmd.Validate()
		case "compare":
			cmd.Compare()
		case "update":
			cmd.Update()
		case "wrap":
			cmd.Wrap()
		case "unwrap":
			cmd.Unwrap()
		case "pull":
			cmd.Pull()
		case "push":
			cmd.Push()
		case "status":
			cmd.Status()
		case "configure":
			cmd.Configure()
		case "delete":
			cmd.Delete()
		case "path":
			cmd.Path()
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
		}
	}
}
