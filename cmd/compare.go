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

func Compare() {
	if _, err := os.Stat("terraform.tfstate"); os.IsNotExist(err) {
		fmt.Println("❌ Error: Local terraform.tfstate file not found.")
		return
	}

	fmt.Println("Executing Compare command...")

	fmt.Println("Pulling state file from Vault...")
	Pull()

	fmt.Println("Unwrapping state file...")
	if err := unwrapAndSaveAs("terraform_remote_compare_pull.tfstate"); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}

	fmt.Println("Comparing state files...")
	if err := compareStateFiles("terraform.tfstate", "terraform_remote_compare_pull.tfstate"); err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("✅ Comparison completed successfully.")
}

func compareStateFiles(localFile, remoteFile string) error {
	localContent, err := os.ReadFile(localFile)
	if err != nil {
		return fmt.Errorf("error reading local file: %w", err)
	}

	remoteContent, err := os.ReadFile(remoteFile)
	if err != nil {
		return fmt.Errorf("error reading remote file: %w", err)
	}

	if string(localContent) != string(remoteContent) {
		fmt.Println("⚠️ State files differ.")
	} else {
		fmt.Println("✅ State files are identical.")
	}

	return nil
}
