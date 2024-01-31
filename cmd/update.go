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
	"os/exec"
	"runtime"
	"strings"
)

const repositoryOwner = "DFW1N"
const repositoryName = "vaultify"

// Update command implementation
func Update() {
	installedVersion, err := getInstalledVersion()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	latestVersion, err := getLatestReleaseTag(repositoryOwner, repositoryName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if installedVersion == latestVersion {
		fmt.Println("\033[33mVaultify\033[0m is already up to date.")
		return
	}

	fmt.Println("Updating \033[33mVaultify\033[0m...")

	// Build the download URL for the latest release binary
	downloadURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/vaultify", repositoryOwner, repositoryName, latestVersion)

	// Use wget to download the binary and hide output
	downloadCmd := exec.Command("wget", downloadURL)
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	if runtime.GOOS == "linux" {
		downloadCmd.Stdout = nil
		downloadCmd.Stderr = nil
	}
	if err := downloadCmd.Run(); err != nil {
		fmt.Println("Error downloading binary:", err)
		return
	}

	// Make the downloaded binary executable
	chmodCmd := exec.Command("chmod", "+x", "vaultify")
	chmodCmd.Stdout = os.Stdout
	chmodCmd.Stderr = os.Stderr
	if err := chmodCmd.Run(); err != nil {
		fmt.Println("Error making binary executable:", err)
		return
	}

	// Move the binary to /usr/local/bin (requires sudo)
	moveCmd := exec.Command("sudo", "mv", "vaultify", "/usr/local/bin/")
	moveCmd.Stdout = os.Stdout
	moveCmd.Stderr = os.Stderr
	if err := moveCmd.Run(); err != nil {
		fmt.Println("Error moving binary to \033[33m/usr/local/bin\033[0m:", err)
		return
	}

	fmt.Println("\033[33mVaultify\033[0m has been updated to version \033[33m" + latestVersion + "\033[0m")
}

// Function to get the version of the currently installed binary
func getInstalledVersion() (string, error) {
	cmd := exec.Command("vaultify", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

// Function to get the latest release tag from a GitHub repository
func getLatestReleaseTag(owner, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	output, err := exec.Command("curl", url).Output()
	if err != nil {
		return "", err
	}

	// Parse the JSON response to extract the tag name
	tagStart := strings.Index(string(output), `"tag_name": "`)
	if tagStart == -1 {
		return "", fmt.Errorf("tag_name not found in GitHub API response")
	}

	tagStart += len(`"tag_name": "`)
	tagEnd := strings.Index(string(output)[tagStart:], `"`) + tagStart

	return string(output)[tagStart:tagEnd], nil
}
