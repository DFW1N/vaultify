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

# Vaultify CLI Command - Init

## Overview
The `init` command in the Vaultify CLI is used to initialize Vaultify for your operating system. This command ensures that the required environment variables, such as `VAULT_TOKEN` and `VAULT_ADDR`, are correctly set and that Vaultify can authenticate with HashiCorp Vault.

## Functionality
- **Check Environment Variables:**
  The `init` command checks if the `VAULT_TOKEN` and `VAULT_ADDR` environment variables are set. If any of these variables is not set, it displays an error message and exits.

- **Initialize Vaultify:**
  If the environment variables are correctly set, the command proceeds to initialize Vaultify by attempting to authenticate against HashiCorp Vault using the provided token and address.

## Usage
To use the `init` command in the Vaultify CLI, run the following command:
This command initializes Vaultify for your operating system.

```bash
vaultify init
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
