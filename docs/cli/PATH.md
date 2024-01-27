<!-- // ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Easily push, pull and encrypt tofu and terraform statefiles from Vault. #
// ######################################################################################## -->

# Vaultify CLI Command - Path Function

## Overview
The `Path` function is a part of the Vaultify project, designed by Sacha Roussakis-Notter, for managing the paths of Terraform state files in Vault. This function is key for pushing, pulling, and encrypting Terraform and TOFU (Trusted Only First Use) state files with Vault.

## Functionality
- **Engine Name and Data Path Configuration:** 
  The function sets `engineName` as `"kv"` and `dataPath` as `"vaultify"`.

- **Workspace Name Retrieval:**
  Retrieves the current Terraform workspace. If an error occurs, it prints an error message and exits the function.

- **Working Directory Retrieval:**
  Obtains the current working directory. On failure, it prints an error message and returns early.

- **Path Construction:**
  Constructs the secret path for the Terraform state file using the engine name, data path, workspace name, and the base name of the working directory.

- **Output:**
  Prints the constructed Vault path for the Terraform state file.

## Usage
To use the `Path` function in the Vaultify CLI, run the following command:
This command will output the Vault path where the Terraform state file is stored.

```bash
vaultify path
```

## Author
- **Sacha Roussakis-Notter:** Maintainer and Creator of the Vaultify project.
