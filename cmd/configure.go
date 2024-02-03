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
    "path/filepath"
)

// Read current configuration from file
func readConfiguration() (*Configuration, error) {
    homeDir, _ := os.UserHomeDir()
    configFile := filepath.Join(homeDir, ".vaultify", "settings.json")

    file, err := os.ReadFile(configFile)
    if err != nil {
        return nil, err
    }

    var config Configuration
    if err := json.Unmarshal(file, &config); err != nil {
        return nil, err
    }

    return &config, nil
}

// Write configuration to file
func writeConfiguration(config *Configuration) error {
    homeDir, _ := os.UserHomeDir()
    configFile := filepath.Join(homeDir, ".vaultify", "settings.json")

    data, err := json.MarshalIndent(config, "", "    ")
    if err != nil {
        return err
    }

    return os.WriteFile(configFile, data, 0644)
}

// Validate if the engine name is supported
func isValidEngineName(name string) bool {
    // Add logic to validate engine name
    return true // Placeholder: replace with actual validation
}

func Configure() {

    if err := checkVaultifySetup(); err != nil {
        fmt.Println(err)
        fmt.Println("Please run \033[33m'vaultify init'\033[0m to set up \033[33mVaultify\033[0m.")
        return
    }

    config, err := readConfiguration()
    if err != nil {
        fmt.Println("❌ Error reading configuration:", err)
        return
    }

    for {
        fmt.Println("\n\033[33mVaultify Settings\033[0m")
        fmt.Println("-------------------------------------------")
        fmt.Printf("%-20s \033[33m%s\033[0m\n", "\033[33m1\033[0m. Default Engine Name:", config.Settings.DefaultEngineName)
        fmt.Printf("%-20s \033[33m%t\033[0m\n", "\033[33m2\033[0m. Use Terraform Workspaces:", config.Settings.TerraformWorkspace)
        fmt.Printf("%-20s \033[33m%s\033[0m\n", "\033[33m3\033[0m. Default Secret Storage:", config.Settings.DefaultSecretStorage)
        fmt.Printf("%-20s \033[33m%s\033[0m\n", "\033[33m4\033[0m. View Secret Storage Settings", "\033[97m(\033[33mAmazon\033[0m \033[97mor\033[0m \033[33mAzure\033[0m\033[97m)\033[0m")
        fmt.Println("-------------------------------------------")
        fmt.Println("Enter the number of the option you want to change (or \033[33m0\033[0m to exit):")

        var choice int
        fmt.Print("\033[33mChoice\033[0m: ")
        _, err := fmt.Scanln(&choice)
        if err != nil {
            fmt.Println("Invalid input. Please enter a \033[33mnumber\033[0m.")
            continue
        }

        switch choice {
        case 0:
            fmt.Println("\nExiting configuration.")
            return
        case 1:
            fmt.Print("\n\033[33mEnter new Default Engine Name\033[0m: ")
            var engineName string
            fmt.Scanln(&engineName)
            if !isValidEngineName(engineName) {
                fmt.Println("Invalid engine name. Please enter a valid \033[33mengine name\033[0m.")
                continue
            }
            config.Settings.DefaultEngineName = engineName
        case 2:
            fmt.Println("\nEnable \033[33mTerraform\033[0m Workspaces:")
            fmt.Println("\033[33m1\033[0m. True")
            fmt.Println("\033[33m2\033[0m. False")
            var workspaceOption int
            fmt.Print("\033[97m------------------\033[0m\n")
            fmt.Print("\033[33mChoice\033[0m: ")
            fmt.Scanln(&workspaceOption)
            switch workspaceOption {
            case 1:
                config.Settings.TerraformWorkspace = true
            case 2:
                config.Settings.TerraformWorkspace = false
            default:
                fmt.Println("Invalid option. Please select \033[33m1\033[0m for True or \033[33m2\033[0m for False.")
                continue
            }
        case 3:
            fmt.Println("\n\033[33mSelect Default Secret Storage:\033[0m")
            fmt.Println("\033[33m1\033[0m. Hashicorp Vault")
            fmt.Println("\033[33m2\033[0m. AWS S3 Bucket")
            fmt.Println("\033[33m3\033[0m. Azure Storage")
            var storageOption int
			fmt.Print("\033[97m------------------\033[0m\n")
			fmt.Print("\033[33mChoice\033[0m: ")
            fmt.Scanln(&storageOption)
            switch storageOption {
            case 1:
                config.Settings.DefaultSecretStorage = "vault"
            case 2:
                config.Settings.DefaultSecretStorage = "s3"
            case 3:
                config.Settings.DefaultSecretStorage = "azure_storage"
            default:
                fmt.Println("Invalid option. Please select \033[33m1\033[0m, \033[33m2\033[0m, or \033[33m3\033[0m.")
                continue
            }
        case 4:
            fmt.Println("\n\033[33mSecret Storage Settings:\033[0m")
            fmt.Print("\033[97m----------------------\033[0m\n")
            fmt.Println("\033[33m1\033[0m. Amazon Web Services")
            fmt.Println("\033[33m2\033[0m. Azure")
            fmt.Println("")
            fmt.Print("\033[33mSelect the platform to view settings\033[0m: ")
            var platformChoice int
            fmt.Scanln(&platformChoice)
            switch platformChoice {
            case 1:
                fmt.Println("\n\033[33mAWS Settings:\033[0m")
                fmt.Print("\033[97m----------------------\033[0m\n")
                awsSettings, _ := json.MarshalIndent(config.Settings.AWS, "", "  ")
                fmt.Println(string(awsSettings))
            case 2:
                fmt.Println("\n\033[33mAzure Settings:\033[0m")
                fmt.Print("\033[97m----------------------\033[0m\n")
                azureSettings, _ := json.MarshalIndent(config.Settings.Azure, "", "  ")
                fmt.Println(string(azureSettings))
            default:
                fmt.Println("\033[33mInvalid option\033[0m. Please select a valid platform.")
            }
        default:
            fmt.Println("\033[33mInvalid option\033[0m. Please enter a valid \033[33mnumber\033[0m.")
        }

        // Save updated configuration
        if err := writeConfiguration(config); err != nil {
            fmt.Println("❌ Error saving \033[33mconfiguration\033[0m:", err)
            return
        }

        fmt.Println("\nConfiguration in \033[33m`settings.json`\033[0m updated/viewed successfully.")
    }
}
