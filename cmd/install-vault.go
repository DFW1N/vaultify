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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type DockerTag struct {
	Name string `json:"name"`
}

type DockerTagsResponse struct {
	Results []DockerTag `json:"results"`
}

func getLatestTag(repo string) (string, error) {
	url := "https://registry.hub.docker.com/v2/repositories/" + repo + "/tags"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tagsResponse DockerTagsResponse
	err = json.Unmarshal(body, &tagsResponse)
	if err != nil {
		return "", err
	}

	if len(tagsResponse.Results) > 0 {
		return tagsResponse.Results[0].Name, nil
	}

	return "", nil
}

func isDockerInstalled() bool {
	cmd := exec.Command("docker", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func isDockerRunning() bool {
	cmd := exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func InstallVault() {

	const networkName = "vault_network"
	const volumeName = "docker_vault_data"
	const containerName = "vault-raft-backend"

	var out bytes.Buffer
	var stderr bytes.Buffer

	latestTag, err := getLatestTag("buungroup/vault-raft")
	if err != nil || latestTag == "" {
		log.Fatalf("Failed to fetch the latest tag: \033[33m%v\033[0m", err)
	}

	imageName := "buungroup/vault-raft:" + latestTag

	if !isDockerInstalled() {
		log.Fatal("\033[33mDocker\033[0m is not installed.")
	}

	if !isDockerRunning() {
		log.Fatal("\033[97mDocker is not running. Please start the \033[33mDocker\033[97m daemon.\033[0m")
	}

	if !networkExists(networkName) {
		createNetworkCmd := exec.Command("docker", "network", "create", networkName)
		if err := createNetworkCmd.Run(); err != nil {
			log.Fatalf("Failed to create network: %v", err)
		} else {
			log.Println("\033[97mNetwork\033[33m " + networkName + "\033[0m \033[97created successfully.\033[0m")
		}
	}

	if !volumeExists(volumeName) {
		createVolumeCmd := exec.Command("docker", "volume", "create", volumeName)
		if err := createVolumeCmd.Run(); err != nil {
			log.Fatalf("Failed to create volume: %v", err)
		} else {
			log.Println("\033[97mVolume\033[33m " + volumeName + "\033[0m \033[97mcreated successfully.\033[0m")
		}
	}

	if !containerIsRunning(containerName) {
		runCmd := exec.Command("docker", "run", "-d",
			"--name", containerName,
			"--network", "vault_network",
			"--cap-add", "IPC_LOCK",
			"--restart", "always",
			"-e", "VAULT_API_ADDR=http://0.0.0.0:8200",
			"-p", "8200:8200",
			"-v", "vault_data:/vault/data",
			imageName)

		runCmd.Stdout = &out
		if err := runCmd.Run(); err != nil {
			log.Fatalf("Failed to run \033[33mVault\033[0m Docker container: %v", err)
		}
		log.Println("\033[33mHashiCorp Vault\033[0m has been installed and is running in a \033[33mDocker\033[0m container.")
	} else {
		log.Println("\033[33mVault Docker container\033[0m\033[33m " + containerName + " \033[0mis already running.")
	}

	execCmd := exec.Command("docker", "exec", "vault-raft-backend", "/bin/bash", "-c", "export VAULT_ADDR=$(grep VAULT_ADDR /root/.bashrc | cut -d'=' -f2 | tr -d \"'\") && /vault/config/initialize-vault.sh")
	execCmd.Stdout = &out
	execCmd.Stderr = &stderr
	if err := execCmd.Run(); err != nil {
		log.Fatalf("Failed to execute the commands inside the container: \033[33m%v\033[0m, Output: \033[33m%s\033[0m, Error: \033[33m%s\033[0m", err, out.String(), stderr.String())
	}

	log.Printf("Output: \033[33m%s\033[0m", out.String())

	out.Reset()
	stderr.Reset()

	cmdStr := `export VAULT_ADDR=$(grep VAULT_ADDR /root/.bashrc | cut -d'=' -f2 | tr -d "'") && ` +
		`export VAULT_TOKEN=$(grep VAULT_TOKEN /root/.bashrc | cut -d'=' -f2 | tr -d "'") && ` +
		`vault secrets enable kv && ` +
		`echo -e "\033[32mVault is set up login with:\033[0m" && ` + // Green color
		`echo -e "\033[97mAddress: $VAULT_ADDR" && ` + // Blue color
		`echo -e "\033[97mToken: $VAULT_TOKEN" && ` + // Magenta color
		`vaultify init && vaultify status`

	execCmd = exec.Command("docker", "exec", "vault-raft-backend", "/bin/bash", "-c", cmdStr)
	execCmd.Stdout = &out
	execCmd.Stderr = &stderr
	var vaultAddr, vaultToken string
	if err := execCmd.Run(); err != nil {
		log.Fatalf("Failed to execute the commands inside the container: \033[33m%v\033[0m, Output: \033[33m%s\033[0m, Error: \033[33m%s\033[0m", err, out.String(), stderr.String())
	} else {
		for _, line := range strings.Split(out.String(), "\n") {
			if strings.Contains(line, "Address:") { // Adjusted to match the echo output
				vaultAddr = strings.TrimSpace(strings.Split(line, "Address:")[1])
			} else if strings.Contains(line, "Token:") { // Adjusted to match the echo output
				vaultToken = strings.TrimSpace(strings.Split(line, "Token:")[1])
			}
		}
	}

	log.Printf("\033[33m%s\033[0m", out.String())
	// log.Printf("Vault Address: \033[33m%s\033[0m", vaultAddr)
	// log.Printf("Vault Token: \033[33m%s\033[0m", vaultToken)

	if promptUserForVaultEnvVariables() {
		exportEnvVariables(vaultAddr, vaultToken)
	}
}

func exportEnvVariables(vaultAddr, vaultToken string) {
	shellConfigFile := determineShellConfigFile()

	if shellConfigFile == "" {
		fmt.Println("Could not determine your shell configuration file (\033[33m.bashrc\033[0m or \033[33m.zshrc\033[0m).")
		return
	}

	appendCmds := fmt.Sprintf("echo 'export VAULT_ADDR=\"%s\"' >> %s && echo 'export VAULT_TOKEN=\"%s\"' >> %s",
		vaultAddr, shellConfigFile, vaultToken, shellConfigFile)

	execCmd := exec.Command("/bin/sh", "-c", appendCmds)
	if err := execCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to append environment variables to %s: %v\n", shellConfigFile, err)
		return
	}

	fmt.Printf("Appended \033[33mVAULT_ADDR\033[0m and \033[33mVAULT_TOKEN\033[0m to %s. Please run \033[33m'source %s'\033[0m to apply changes.\n", shellConfigFile, shellConfigFile)
}

func determineShellConfigFile() string {
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "zsh") {
		return os.Getenv("HOME") + "/.zshrc"
	} else if strings.Contains(shell, "bash") {
		return os.Getenv("HOME") + "/.bashrc"
	}
	return ""
}

func promptUserForVaultEnvVariables() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to export \033[33mVAULT_ADDR\033[0m and \033[33mVAULT_TOKEN\033[0m to your host's environment variables? This will modify your \033[33m.bashrc\033[0m or \033[33m.zshrc\033[0m file. [y/n]")
	fmt.Print("\033[33mChoice\033[0m: ")
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input: \033[33m%v\033[0m\n", err)
		return false
	}

	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes"
}
