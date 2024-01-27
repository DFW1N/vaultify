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

# Vaultify CLI Command - Validate

## Overview
The `validate` command in the Vaultify CLI allows you to validate the `terraform.tfstate` file located in the current working directory. This command ensures that the state file is valid JSON and correctly formatted for Terraform.

## Functionality
- **Check State File Existence:**
  The `validate` command first checks if the `terraform.tfstate` file exists in the current working directory. If the file is not found, it reports an error and exits.

- **JSON Validation:**
  It then attempts to open and read the `terraform.tfstate` file. If successful, it decodes the JSON content to validate its structure.
  
- **Validation Result:**
  If the JSON content is correctly formatted, the command reports that the validation passed. Otherwise, it reports that the validation failed and specifies that the Terraform state file is not valid JSON.

## Usage
To use the `validate` command in the Vaultify CLI, run the following command:

```bash
vaultify validate
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
