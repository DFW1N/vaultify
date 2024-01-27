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

func Unwrap() {
	// Check if terraform.tfstate already exists
	if _, err := os.Stat("terraform.tfstate"); err == nil {
		fmt.Println("✅ terraform.tfstate file already exists in the current directory.")
		fmt.Println("⚠️  No need to unwrap. Exiting.")
		return
	}

	if _, err := os.Stat("terraform.tfstate.gz.b64"); os.IsNotExist(err) {
		fmt.Println("❌ Error: terraform.tfstate.gz.b64 file not found in the current directory.")
		fmt.Println("⚠️  Please run vaultify pull to get this file from your vault, if it exists.")
		os.Exit(1)
	}

	err := decodeBase64("terraform.tfstate.gz.b64", "terraform.tfstate.gz")
	if err != nil {
		fmt.Println("❌ Error: Base64 decoding failed.", err)
		os.Exit(1)
	}
	//fmt.Println("✅ Base64 decoding successful.")

	// Check if the file is created and not empty
	if fileInfo, err := os.Stat("terraform.tfstate.gz"); os.IsNotExist(err) || fileInfo.Size() == 0 {
		fmt.Println("❌ Error: terraform.tfstate.gz is empty or not created.")
		os.Exit(1)
	}
	//fmt.Println("✅ File terraform.tfstate.gz created successfully.")

	// Gunzip the file
	err = gunzipFile("terraform.tfstate.gz", "terraform.tfstate")
	if err != nil {
		fmt.Println("❌ Error: Decompression failed.", err)
		os.Exit(1)
	}
	//fmt.Println("✅ Decompression successful.")

	// Delete the terraform.tfstate.gz file
	if err := os.Remove("terraform.tfstate.gz"); err != nil {
		fmt.Println("❌ Error: Failed to delete terraform.tfstate.gz file.", err)
		os.Exit(1)
	}
	//fmt.Println("✅ Deleted terraform.tfstate.gz file.")

	// Delete the terraform.tfstate.gz.b64 file
	if err := os.Remove("terraform.tfstate.gz.b64"); err != nil {
		fmt.Println("❌ Error: Failed to delete terraform.tfstate.gz.b64 file.", err)
		os.Exit(1)
	}
	//fmt.Println("✅ Deleted terraform.tfstate.gz.b64 file.")

	// Output the terraform.tfstate file
	fmt.Println("✅ Vaultify successfully unwrapped terraform.tfstate.")
	fmt.Println("⚠️  The unwrapped state file is named terraform.tfstate and can be found in the working directory.")
}

func decodeBase64(inputFile, outputFile string) error {
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	decodedData, err := base64.StdEncoding.DecodeString(string(inputData))
	if err != nil {
		return err
	}

	err = os.WriteFile(outputFile, decodedData, 0644)
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

	err = os.WriteFile(outputFile, out, 0644)
	if err != nil {
		return err
	}

	return nil
}
