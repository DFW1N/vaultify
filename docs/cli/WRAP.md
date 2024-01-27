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

# Vaultify CLI Command - Wrap

## Overview
The `wrap` command in the Vaultify CLI allows you to wrap and encode a Terraform state file (.tfstate) to prepare it for storage in HashiCorp Vault. This command performs the following steps:
1. Compresses the Terraform state file using gzip.
2. Encodes the compressed state file using base64.
3. Saves the encoded state to a temporary file.
4. Deletes the original Terraform state file.
5. Tells you to use `vaultify push` to ensure your encoded state has been stored and saved securely.
6. Deletes the encoded state file in the temporary directory.

## Functionality
- **Wrap Terraform State:**
  The `wrap` command identifies the first found .tfstate file in the current directory, compresses it using gzip, encodes it with base64, and sets it as an environment variable.

- **Temporary File Storage:**
  The encoded state is saved to a temporary file (`/tmp/.encoded_wrap`) to ensure that it is not lost in case of errors.

- **Original State File Deletion:**
  After wrapping is complete, the original .tfstate file is deleted to prevent duplication.

## Usage
To wrap and encode a Terraform state file, run the following command:

```bash
vaultify wrap
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
