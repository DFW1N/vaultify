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
	passphrase := os.Getenv("VAULTIFY_PASSPHRASE")
	if passphrase == "" {
		fmt.Println("❌ Error: \033[33mVAULTIFY_PASSPHRASE\033[0m environemnt variable not set.")
		os.Exit(1)
	}
	if _, err := os.Stat("terraform.tfstate"); os.IsNotExist(err) {
		fmt.Println("❌ Error: \033[33mLocal terraform.tfstate\033[0m file not found.")
		return
	}

	fmt.Println("Pulling state file from \033[33mVault\033[0m...")
	Pull()

	if err := unwrapAndSaveAs("terraform_remote_compare_pull.tfstate", passphrase); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}

	if err := compareStateFiles("terraform.tfstate", "terraform_remote_compare_pull.tfstate"); err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("✅ \033[33mComparison\033[0m completed successfully.")
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
		fmt.Println("⚠️ State files \033[33mdiffer\033[0m.")
	} else {
		fmt.Println("✅ State files are \033[33midentical\033[0m.")
	}

	return nil
}
