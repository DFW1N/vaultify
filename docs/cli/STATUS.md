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

# Vaultify CLI Command - Status

## Overview
The `status` command in the Vaultify CLI is used to check the authentication status of Vaultify with the HashiCorp Vault server. It verifies if Vaultify is properly authenticated and connected to Vault.

## Functionality
- **Check Authentication:**
  The `status` command uses the `curl` command to perform an authenticated operation against the HashiCorp Vault server.

- **Verify Status:**
  It checks the response from the `curl` command to determine whether Vaultify is authenticated and connected to Vault.

## Usage
To use the `status` command in the Vaultify CLI, run the following command:
This command checks the authentication status of Vaultify with the HashiCorp Vault server.

```bash
vaultify status
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
