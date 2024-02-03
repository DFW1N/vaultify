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
	"encoding/json"
	"fmt"
	"os"
)

func Validate() {
	if _, err := os.Stat("terraform.tfstate"); os.IsNotExist(err) {
		fmt.Println("❌ Error: \033[33mterraform.tfstate\033[0m file not found in the current directory nothing to validate.")
		os.Exit(1)
	}

	file, err := os.Open("terraform.tfstate")
	if err != nil {
		fmt.Printf("❌ Error opening \033[33mterraform.tfstate\033[0m file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var state map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&state); err != nil {
		fmt.Printf("❌ Error decoding JSON: %v\n", err)
		fmt.Println("❗ Validation failed: Terraform state file is not valid JSON.")
		os.Exit(1)
	}

	fmt.Println("✅ Validation passed: Terraform state file is correctly formatted.")
}
