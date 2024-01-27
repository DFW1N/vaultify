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
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Delete command implementation
func Delete() {
	curlCommand := "curl"
	vaultURL := os.Getenv("VAULT_ADDR")
	engineName := "kv"
	dataPath := "vaultify"

	workspaceName, err := getCurrentWorkspace()
	if err != nil {
		fmt.Println("❌ Error getting current Terraform workspace:", err)
		return
	}
	workspaceName = strings.TrimSpace(workspaceName) // Ensure to trim any whitespace

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Error getting current working directory:", err)
		return
	}

	workingDirName := filepath.Base(workingDir)
	secretPath := fmt.Sprintf("%s/%s/%s_%s", dataPath, workspaceName, workingDirName, "terraform.tfstate")

	// Ask for confirmation
	fmt.Printf("Are you sure you want to delete the secret at '%s'? [y/N]: ", secretPath)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if strings.ToLower(input) != "y" && strings.ToLower(input) != "yes" {
		fmt.Println("Deletion canceled.")
		return
	}

	deleteCmd := exec.Command(
		curlCommand,
		"--silent", "--show-error",
		"--header", "X-Vault-Token: "+os.Getenv("VAULT_TOKEN"),
		"--request", "DELETE",
		vaultURL+"/v1/"+engineName+"/data/"+secretPath,
	)

	_, err = deleteCmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Error deleting secret from Vault:", err)
		return
	}

	fmt.Println("✅ Secret successfully deleted from Vault.")
}
