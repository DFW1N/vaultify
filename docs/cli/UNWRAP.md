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

# Vaultify CLI Command - Unwrap

## Overview
The `unwrap` command in the Vaultify CLI is used to unwrap a secret from base64 encoding and save it as the `terraform.tfstate` file. It allows you to retrieve and decode the Terraform statefile from HashiCorp Vault.

## Functionality
- **Check File Existence:**
  The `unwrap` command first checks if the `terraform.tfstate.gz.b64` file exists in the current directory. If it does not exist, an error message is displayed, and the command exits.

- **Prompt for Overwrite:**
  If the `terraform.tfstate` file already exists in the current directory, the command prompts the user to choose whether to overwrite it. You can respond with "yes," "no," or "rename."

- **Unwrap and Save:**
  - If you choose "no," the command exits without making changes.
  - If you choose "rename," the unwrapped state file is saved as `terraform_remote_pull.tfstate` instead.
  - If you choose "yes" or do not specify any other response, the `terraform.tfstate` file is overwritten with the unwrapped state.

## Usage
To use the `unwrap` command in the Vaultify CLI, run the following command:
This command unwraps and saves the decoded Terraform statefile as `terraform.tfstate` in the current directory.

```bash
vaultify unwrap
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
