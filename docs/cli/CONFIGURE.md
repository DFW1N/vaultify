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

# Vaultify CLI Command - Configure

## Overview
The `configure` command in the Vaultify CLI allows users to configure various settings for interacting with Vault. This command ensures that the required environment variables are set and provides an interface to change configuration options.

## Functionality
- **Check VAULT_TOKEN Environment Variable:**
  The function checks if the `VAULT_TOKEN` environment variable is set. If it is not set, it displays an error message and exits.

- **Check VAULT_ADDR Environment Variable:**
  The function checks if the `VAULT_ADDR` environment variable is set. If it is not set, it displays an error message and exits.

- **Initialize Configuration:**
  The command initializes the configuration with values from the environment variables `VAULT_TOKEN` and `VAULT_ADDR`.

- **Present Configuration Options:**
  The current configuration options are presented in a table format, including the Vault Address.

- **User Interaction:**
  The user is prompted to select an option to change the configuration. The user can enter a new Vault Address if needed.

- **Exit Configuration:**
  The user can exit the configuration process by entering '0'.

## Usage
To use the `configure` command in the Vaultify CLI, run the following command:
This command allows you to configure Vaultify options interactively.

```bash
vaultify configure
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
