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
	passphrase := os.Getenv("VAULTIFY_PASSPHRASE")
	if passphrase == "" {
		fmt.Println("❌ Error: \033[33mVAULTIFY_PASSPHRASE\033[0m environemnt variable not set.")
		os.Exit(1)
	}
	if _, err := os.Stat("terraform.tfstate.gz.enc.b64"); os.IsNotExist(err) {
		fmt.Println("❌ Error: \033[33mterraform.tfstate.gz.enc.b64\033[0m file not found in the current directory.")
		fmt.Println("⚠️  Please run vaultify pull to get this file from your vault, if it exists.")
		os.Exit(1)
	}

	if _, err := os.Stat("terraform.tfstate"); err == nil {
		fmt.Println("✅ \033[33mterraform.tfstate\033[0m file already exists in the current directory.")
		fmt.Print("Do you want to overwrite it (yes/no/rename): ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "no" {
			fmt.Println("Exiting without making changes.")
			return
		} else if response == "rename" {
			fmt.Println("Saving as terraform_remote_pull.tfstate instead.")
			if err := unwrapAndSaveAs("terraform_remote_pull.tfstate", passphrase); err != nil {
				fmt.Println("❌ Error:", err)
				os.Exit(1)
			}
			return
		} else if response != "yes" {
			fmt.Println("Invalid response. Exiting.")
			return
		}
	}

	if err := unwrapAndSaveAs("terraform.tfstate", passphrase); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}

func unwrapAndSaveAs(outputFileName string, passphrase string) error {
	err := decodeBase64("terraform.tfstate.gz.enc.b64", "terraform.tfstate.gz.enc")
	if err != nil {
		return fmt.Errorf("base64 decoding failed: %w", err)
	}
	// decrypt
	_, err = decryptFile("terraform.tfstate.gz.enc", passphrase)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = gunzipFile("terraform.tfstate.gz", outputFileName)
	if err != nil {
		return fmt.Errorf("decompression failed: %w", err)
	}

	if err = os.Remove("terraform.tfstate.gz.enc"); err != nil {
		return fmt.Errorf("failed to delete terraform.tfstate.gz.enc: %w", err)
	}

	if err = os.Remove("terraform.tfstate.gz.enc.b64"); err != nil {
		return fmt.Errorf("failed to delete terraform.tfstate.gz.enc.b64: %w", err)
	}

	fmt.Printf("✅ Unwrapped state file saved as \033[33m%s\033[0m\n", outputFileName)
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
