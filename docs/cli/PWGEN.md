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
The `pwgen` command in the Vaultify CLI is designed to provide you a randomly generated token for your environment variable `VAULTIFY_PASSPHRASE` thats used to encrypt your state files or secrets.

## Functionality
- **Generates Random Passphrase:**
  The function generates a random passphrase

> This is an optional CLI command, that provides a easy way to generate a randomised value, vaultify is not responsible for storing this secret and if lost is at risk of the user this is simply to make things easier.


## Usage
To use the `pwgen` command in the Vaultify CLI, run one of the following commands:

```bash
vaultify pwgen
```

### Example Output

```bash
go run main.go pwgen                    
Passphrase: A*jaPz=BT@Va@gTWn9N#Jzx0PCoJTfl=
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
| **Sharath Nair** |
| *Maintainer & Contributor* |