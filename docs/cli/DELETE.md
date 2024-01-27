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

# Vaultify CLI Command - Delete

## Overview
The `delete` command in the Vaultify CLI allows users to delete a secret stored in Vault. This command is useful for removing sensitive information from the Vault.

## Functionality
- **Configure CURL Command:**
  The command configures the `curl` command to interact with Vault using the `VAULT_ADDR` environment variable.

- **Retrieve Workspace and Working Directory Information:**
  It retrieves the current Terraform workspace and working directory information.

- **Compose Secret Path:**
  The secret path is composed using the `dataPath`, Terraform workspace name, and the base name of the working directory.

- **Confirmation Prompt:**
  Before deleting the secret, the user is prompted for confirmation. The command will not proceed without user consent.

- **Execute Deletion:**
  If the user confirms, the secret is deleted from Vault using the `curl` command with appropriate headers and request parameters.

## Usage
To use the `delete` command in the Vaultify CLI, run the following command:
This command allows you to delete a secret from Vault.

```bash
vaultify delete
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
