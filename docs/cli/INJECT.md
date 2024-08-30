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

# Vaultify CLI Command - Inject

## Overview
The `inject` command in the Vaultify CLI is designed to securely store secrets in HashiCorp Vault. It allows users to specify a custom path, key, and value for the secret, or use default settings based on the current workspace.

## Functionality
- **Parse Command-Line Flags:**
  The function parses flags for the Vault path and secret key.

- **Check Vaultify Setup:**
  Ensures that Vaultify is properly set up and Vault is initialized.

- **Determine Secret Path:**
  Uses the provided path or generates a default path based on the current workspace and directory.

- **Ensure KV Path Exists:**
  Checks if the specified path exists in Vault and creates it if necessary.

- **Write Secret to Vault:**
  Writes the provided secret to Vault at the specified path and key.

## Usage
To use the `inject` command in the Vaultify CLI, run one of the following commands:

### Custom Path
```bash
vaultify inject -path "foo/bar/tfc_token" "my_secret_token"
```

### Key as Path
```bash
vaultify inject -key "my_secret" "secretvalue"
```

### Use default path
```bash
vaultify inject "secretvalue"
```

- `-path`: Optional. Specifies the Vault path to store the secret. If not provided, a default path based on the current workspace is used.
- `-key`: Optional. Specifies the key for the secret. Defaults to "default" if not provided this will be your secret name.
- `<secret_value>`: Required. The secret value to be stored in the secret name in Vault.

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |