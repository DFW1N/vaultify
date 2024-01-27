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
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
)

// Unwrap command implementation
func Unwrap() {
	// Check if terraform.tfstate.gz.b64 file exists in the working directory
	if _, err := os.Stat("terraform.tfstate.gz.b64"); os.IsNotExist(err) {
		fmt.Println("❌ Error: terraform.tfstate.gz.b64 file not found in the current directory.")
		os.Exit(1)
	}

	// Decode base64 to a temporary file
	err := decodeBase64("terraform.tfstate.gz.b64", "terraform.tfstate.gz")
	if err != nil {
		fmt.Println("❌ Error: Base64 decoding failed.", err)
		os.Exit(1)
	}
	fmt.Println("✅ Base64 decoding successful.")

	// Check if the file is created and not empty
	if fileInfo, err := os.Stat("terraform.tfstate.gz"); os.IsNotExist(err) || fileInfo.Size() == 0 {
		fmt.Println("❌ Error: terraform.tfstate.gz is empty or not created.")
		os.Exit(1)
	}
	fmt.Println("✅ File terraform.tfstate.gz created successfully.")

	// Gunzip the file
	err = gunzipFile("terraform.tfstate.gz", "terraform-test.tfstate")
	if err != nil {
		fmt.Println("❌ Error: Decompression failed.", err)
		os.Exit(1)
	}
	fmt.Println("✅ Decompression successful.")

	// Delete the terraform.tfstate.gz file
	if err := os.Remove("terraform.tfstate.gz"); err != nil {
		fmt.Println("❌ Error: Failed to delete terraform.tfstate.gz file.", err)
		os.Exit(1)
	}
	fmt.Println("✅ Deleted terraform.tfstate.gz file.")

	// Output the terraform.tfstate file
	fmt.Println("✅ Vaultify successfully unwrapped terraform.tfstate.")
	fmt.Println("The unwrapped state file is named terraform-test.tfstate and can be found in the working directory.")
}

func decodeBase64(inputFile, outputFile string) error {
	inputData, err := os.ReadFile(inputFile) // Use os.ReadFile (Go 1.16+)
	if err != nil {
		return err
	}

	decodedData, err := base64.StdEncoding.DecodeString(string(inputData))
	if err != nil {
		return err
	}

	err = os.WriteFile(outputFile, decodedData, 0644) // Use os.WriteFile (Go 1.16+)
	if err != nil {
		return err
	}

	return nil
}

func gunzipFile(inputFile, outputFile string) error {
	cmd := exec.Command("gunzip", "-c", inputFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	err = os.WriteFile(outputFile, out, 0644) // Use os.WriteFile (Go 1.16+)
	if err != nil {
		return err
	}

	return nil
}
