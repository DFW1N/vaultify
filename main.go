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
			fmt.Println("Usage: \033[33mvaultify\033[0m [command]")
			fmt.Println("Use \033[33m'vaultify -h'\033[0m for help.")
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
		case "retrieve":
			cmd.Retrieve()
		case "publish":
			cmd.Publish()
		case "permissions":
			cmd.TokenPermissions()
		case "install-vault":
			cmd.InstallVault()
		case "delete-vault":
			if len(os.Args) > 1 && os.Args[1] == "delete-vault" {
				handleDeleteVaultCommand(os.Args[2:])
				return
			}
		case "pwgen":
			cmd.GenPassphrase()
		default:
			fmt.Printf("Unknown command: \033[33m%s\033[0m\n", os.Args[1])
			fmt.Println("Use \033[33m'vaultify -h'\033[0m for help.")
		}
	}
}

func handleDeleteVaultCommand(args []string) {
	deleteVaultFlagSet := flag.NewFlagSet("delete-vault", flag.ContinueOnError)
	autoConfirm := deleteVaultFlagSet.Bool("y", false, "Automatically confirm deletion without prompting")
	err := deleteVaultFlagSet.Parse(args)
	if err != nil {
		fmt.Println("Error parsing delete-vault flags:", err)
		return
	}

	cmd.DeleteVault(*autoConfirm)
}
