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

# Vaultify CLI Command - Get

## Overview
The `get` command in the Vaultify CLI is designed to retrieve secrets from HashiCorp Vault. It allows users to fetch either all secrets at a given path or a specific secret by key, providing flexibility in secret management.

## Functionality
- **Parse Command-Line Flags:**
  The function parses flags for the Vault path and secret key.

- **Check Vaultify Setup:**
  Ensures that Vaultify is properly set up and Vault is initialized.

- **Determine Secret Path:**
  Uses the provided path or generates a default path based on the current workspace and directory.

- **Read Secret from Vault:**
  Retrieves the secret(s) from the specified path in Vault.

- **Display Secret(s):**
  Prints either all secrets at the given path or a specific secret if a key is provided.

## Usage
To use the `get` command in the Vaultify CLI, run one of the following commands:

```bash
vaultify get [-path <vault_path>] [-key <secret_key>]
vaultify get -path "my/secret/path"
vaultify get -path "my/secret/path" -key "my_secret_key"
vaultify get
vaultify get -key "my_secret_key"
```

- `-path`: Optional. Specifies the Vault path to retrieve the secret from. If not provided, a default path based on the current workspace is used.
- `-key`: Optional. Specifies the key of the secret to retrieve. If not provided, all secrets at the specified path are displayed.

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |