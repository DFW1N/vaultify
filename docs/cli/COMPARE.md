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

# Vaultify CLI Command - Compare

## Overview
The `compare` command in the Vaultify CLI is designed to compare the local `terraform.tfstate` file with the remote state file pulled from Vault. This function is essential for ensuring the consistency of your Terraform state files.

## Functionality
- **Check Local State File:**
  The function checks if the `terraform.tfstate` file exists locally. If it does not, it prints an error message and exits the function.

- **Execute Pull Command:**
  The `Pull` function is executed to retrieve the remote state file from Vault.

- **Unwrap and Save Remote State:**
  The remote state file is unwrapped and saved as `terraform_remote_compare_pull.tfstate`. If any error occurs during this process, it prints an error message and exits the function.

- **Compare State Files:**
  The function compares the contents of the local and remote state files. If the files differ, it notifies the user. Otherwise, it confirms that the state files are identical.

## Usage
To use the `compare` command in the Vaultify CLI, run the following command:
This command will compare the local `terraform.tfstate` file with the remote state file from Vault.

```bash
vaultify compare
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
