package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Unwrap() {
	if _, err := os.Stat("terraform.tfstate.gz.b64"); os.IsNotExist(err) {
		fmt.Println("❌ Error: terraform.tfstate.gz.b64 file not found in the current directory.")
		fmt.Println("⚠️  Please run vaultify pull to get this file from your vault, if it exists.")
		os.Exit(1)
	}

	if _, err := os.Stat("terraform.tfstate"); err == nil {
		fmt.Println("✅ terraform.tfstate file already exists in the current directory.")
		fmt.Print("Do you want to overwrite it (yes/no/rename): ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "no" {
			fmt.Println("Exiting without making changes.")
			return
		} else if response == "rename" {
			fmt.Println("Saving as terraform_remote_pull.tfstate instead.")
			if err := unwrapAndSaveAs("terraform_remote_pull.tfstate"); err != nil {
				fmt.Println("❌ Error:", err)
				os.Exit(1)
			}
			return
		} else if response != "yes" {
			fmt.Println("Invalid response. Exiting.")
			return
		}
	}

	if err := unwrapAndSaveAs("terraform.tfstate"); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}

func unwrapAndSaveAs(outputFileName string) error {
	err := decodeBase64("terraform.tfstate.gz.b64", "terraform.tfstate.gz")
	if err != nil {
		return fmt.Errorf("base64 decoding failed: %w", err)
	}

	err = gunzipFile("terraform.tfstate.gz", outputFileName)
	if err != nil {
		return fmt.Errorf("decompression failed: %w", err)
	}

	if err = os.Remove("terraform.tfstate.gz"); err != nil {
		return fmt.Errorf("failed to delete terraform.tfstate.gz: %w", err)
	}

	if err = os.Remove("terraform.tfstate.gz.b64"); err != nil {
		return fmt.Errorf("failed to delete terraform.tfstate.gz.b64: %w", err)
	}

	fmt.Printf("✅ Unwrapped state file saved as %s\n", outputFileName)
	return nil
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
