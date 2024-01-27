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

# Vaultify CLI Command - Pull

## Overview
The `pull` command in the Vaultify CLI is used to pull the Terraform state file from a remote HashiCorp Vault server. This command retrieves the state file stored in Vault, decodes it, and saves it locally for further use.

## Functionality
- **Check Secret Path:**
  The `pull` command constructs the secret path in HashiCorp Vault based on the current Terraform workspace and working directory. It then checks if the secret path exists in Vault.

- **Retrieve Secret:**
  If the secret path exists, the command retrieves the secret from Vault, decodes it, and saves it locally as a base64-encoded string.

- **Save to File:**
  The base64-encoded string is saved to a file named `terraform.tfstate.gz.b64` in the current working directory.

## Usage
To use the `pull` command in the Vaultify CLI, run the following command:
This command retrieves the Terraform state file from the remote HashiCorp Vault server.

```bash
vaultify pull
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
