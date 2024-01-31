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
	"os/exec"
	"path/filepath"
	"strings"
)

// ###############################
// # checkVaultifySetup Function #
// ###############################

func checkVaultifySetup() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("❌ Error getting user home directory: \033[33m%v\033[0m", err)
	}

	vaultifyDir := filepath.Join(homeDir, ".vaultify")
	settingsFilePath := filepath.Join(vaultifyDir, "settings.json")

	// Check if .vaultify directory exists
	if _, err := os.Stat(vaultifyDir); os.IsNotExist(err) {
		return fmt.Errorf("❌ Error: \033[33m.vaultify\033[0m directory not found")
	}

	// Check if settings.json exists
	if _, err := os.Stat(settingsFilePath); os.IsNotExist(err) {
		return fmt.Errorf("❌ Error: \033[33msettings.json\033[0m file not found in .vaultify directory")
	}

	return nil
}

// ################################
// # getCurrentWorkspace Function #
// ################################

func getCurrentWorkspace() (string, error) {
	cmd := exec.Command("terraform", "workspace", "show")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// #########################
// # readSettings Function #
// #########################

// readSettings reads the settings from the settings.json file
func readSettings() (*Settings, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("❌ Error getting user home directory: %v", err)
	}

	settingsFilePath := filepath.Join(homeDir, ".vaultify", "settings.json")
	content, err := os.ReadFile(settingsFilePath)
	if err != nil {
		return nil, fmt.Errorf("❌ Error reading settings file: \033[33m%v\033[0m", err)
	}

	var settings Settings
	err = json.Unmarshal(content, &settings)
	if err != nil {
		return nil, fmt.Errorf("❌ Error unmarshalling settings JSON: \033[33m%v\033[0m", err)
	}

	return &settings, nil
}

// #####################################
// # logic to be used within functions #
// #####################################

// // Detecting the shell configuration file
// shellRC, shellType, err := detectShellRC()
// if err != nil {
// 	fmt.Printf("❌Error detecting shell: %v\n", err)
// 	os.Exit(1)
// }

// fmt.Printf("Detected shell type: %s\n", shellType)

// // Setting environment variable in the detected shell configuration file
// if err := setEnvInShellRC("VAULTIFY_TFWORKSPACE", "true", shellRC); err != nil {
// 	fmt.Printf("❌ Error setting \033[33mVAULTIFY_TFWORKSPACE\033[0m: %v\n", err)
// 	os.Exit(1)
// }

// // Inform the user
// fmt.Printf("Please source your shell configuration file to apply the changes:\n")
// fmt.Printf("source \033[33m%s\033[0m\n", shellRC)

// // Setting environment variable for the current process
// os.Setenv("VAULTIFY_TFWORKSPACE", "true")

// ##########################
// # detectShellRC Function #
// ##########################

// func detectShellRC() (string, string, error) {
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		return "", "", err
// 	}

// 	shell := os.Getenv("SHELL")
// 	shellType := "bash"
// 	shellRC := filepath.Join(homeDir, ".bashrc")
// 	if strings.Contains(shell, "zsh") {
// 		shellType = "zsh"
// 		shellRC = filepath.Join(homeDir, ".zshrc")
// 	}

// 	fmt.Println("Detected shell configuration file:", shellRC)
// 	return shellRC, shellType, nil
// }

// ############################
// # setEnvInShellRC Function #
// ############################

// func setEnvInShellRC(variable, value, shellRCPath string) error {
// 	// Check if the variable already exists in the file
// 	if err := checkVariableInFile(shellRCPath, variable); err != nil {
// 		fmt.Println("Environment variable already set in", shellRCPath)
// 		return nil
// 	}

// 	return appendToFile(shellRCPath, fmt.Sprintf("\nexport %s=\"%s\"\n", variable, value))
// }

// #########################
// # appendToFile Function #
// #########################

// func appendToFile(filePath, content string) error {
// 	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
// 	if err != nil {
// 		fmt.Println("❌ Error opening file:", err)
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(content)
// 	if err != nil {
// 		fmt.Println("❌ Error writing to file:", err)
// 	}
// 	return err
// }

// ################################
// # checkVariableInFile Function #
// ################################

// func checkVariableInFile(filePath, variable string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		if strings.Contains(scanner.Text(), variable) {
// 			return fmt.Errorf("variable \033[33m%s\033[0m already set", variable)
// 		}
// 	}
// 	return nil
// }
