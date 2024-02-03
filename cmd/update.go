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

	osName := runtime.GOOS
	archName := runtime.GOARCH

	if archName == "amd64" {
		archName = "x86_64"
	}

	assetName := fmt.Sprintf("vaultify_%s_%s.tar.gz", osName, archName)
	if osName == "linux" && archName == "arm64" {
		assetName = "vaultify_Linux_arm64.tar.gz"
	}

	downloadURL := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", repositoryOwner, repositoryName, latestVersion, assetName)

	downloadCmd := exec.Command("wget", "-q", downloadURL)
	downloadCmd.Stdout = os.Stdout
	downloadCmd.Stderr = os.Stderr
	if err := downloadCmd.Run(); err != nil {
		fmt.Println("Error downloading binary:", err)
		return
	}

	tarCmd := exec.Command("tar", "xzf", fmt.Sprintf("vaultify_%s_%s.tar.gz", osName, archName))
	tarCmd.Stdout = os.Stdout
	tarCmd.Stderr = os.Stderr
	if err := tarCmd.Run(); err != nil {
		fmt.Println("Error extracting binary:", err)
		return
	}

	chmodCmd := exec.Command("chmod", "+x", "vaultify")
	chmodCmd.Stdout = os.Stdout
	chmodCmd.Stderr = os.Stderr
	if err := chmodCmd.Run(); err != nil {
		fmt.Println("Error making binary executable:", err)
		return
	}

	moveCmd := exec.Command("sudo", "mv", "vaultify", "/usr/local/bin/")
	moveCmd.Stdout = os.Stdout
	moveCmd.Stderr = os.Stderr
	if err := moveCmd.Run(); err != nil {
		fmt.Println("Error moving binary to \033[33m/usr/local/bin\033[0m:", err)
		return
	}

	deleteCmd := exec.Command("rm", fmt.Sprintf("vaultify_%s_%s.tar.gz", osName, archName))
	deleteCmd.Stdout = os.Stdout
	deleteCmd.Stderr = os.Stderr
	if err := deleteCmd.Run(); err != nil {
		fmt.Println("Error deleting downloaded .tar.gz file:", err)
		return
	}

	fmt.Println("\033[33mVaultify\033[0m has been updated to version \033[33m" + latestVersion + "\033[0m")
}

func getInstalledVersion() (string, error) {
    return GetVersion(), nil
}

func getLatestReleaseTag(owner, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	output, err := exec.Command("curl", "-s", url).Output()
	if err != nil {
		return "", err
	}

	tagStart := strings.Index(string(output), `"tag_name": "`)
	if tagStart == -1 {
		return "", fmt.Errorf("tag_name not found in GitHub API response")
	}

	tagStart += len(`"tag_name": "`)
	tagEnd := strings.Index(string(output)[tagStart:], `"`) + tagStart

	return string(output)[tagStart:tagEnd], nil
}