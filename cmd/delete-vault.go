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
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func containerExists(containerName string) bool {
	cmd := exec.Command("docker", "ps", "-a", "-q", "-f", "name="+containerName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to query Docker containers: %v", err)
	}
	return strings.TrimSpace(out.String()) != ""
}

func DeleteVault() {
	if !isDockerInstalled() {
		log.Fatal("\033[33mDocker\033[0m is not installed.")
	}

	if !isDockerRunning() {
		log.Fatal("\033[33mDocker\033[0m is not running. Please start the \033[33mDocker daemon\033[0m.")
	}

	const containerName = "vault-raft-backend"

	if containerExists(containerName) {
		if containerIsRunning(containerName) {
			stopCmd := exec.Command("docker", "stop", containerName)
			if err := stopCmd.Run(); err != nil {
				log.Fatalf("Failed to stop Vault Docker container: \033[33m%v\033[0m", err)
			}
		}

		rmCmd := exec.Command("docker", "rm", containerName)
		if err := rmCmd.Run(); err != nil {
			log.Fatalf("Failed to remove Vault Docker container: \033[33m%v\033[0m", err)
		}
	} else {
		log.Printf("\033[33mNo\033[0m Docker container with the name \033[33m'%s'\033[0m found. Skipping stop and remove commands.", containerName)
	}

	deleteVolume("docker_vault_data")
	deleteVolume("vault_data")

	log.Println("Vault Docker \033[33mcontainer\033[0m and \033[33mvolume\033[0m have been cleaned up.")
}
