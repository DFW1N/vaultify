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

```bash
██╗   ██╗ █████╗ ██╗   ██╗██╗  ████████╗██╗███████╗██╗   ██╗
██║   ██║██╔══██╗██║   ██║██║  ╚══██╔══╝██║██╔════╝╚██╗ ██╔╝
██║   ██║███████║██║   ██║██║     ██║   ██║█████╗   ╚████╔╝ 
╚██╗ ██╔╝██╔══██║██║   ██║██║     ██║   ██║██╔══╝    ╚██╔╝  
 ╚████╔╝ ██║  ██║╚██████╔╝███████╗██║   ██║██║        ██║   
  ╚═══╝  ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝   ╚═╝╚═╝        ╚═╝   
                                                            
```

# Introduction

The Vaultify Command-Line Interface (CLI) provides a set of commands to simplify the management of sensitive data and Terraform statefiles with HashiCorp Vault. Each command serves a specific logical purpose, facilitating tasks related to encryption, storage, and retrieval of data in Vault.

This document provides an overview of the available CLI commands, explaining their logical functions and how they can be used to enhance your data management workflow. Whether you need to encrypt Terraform statefiles, push data into Vault, or retrieve stored secrets, Vaultify CLI commands offer the tools you need to securely interact with HashiCorp Vault.

Here's a brief overview of the main Vaultify CLI commands and their logical purposes:

| Command                | Description                                                                                                      |
|------------------------|------------------------------------------------------------------------------------------------------------------|
| [`vaultify init`](https://vaultify.buungroup.com/#/cli/INIT.md) | Initialize Vaultify in your operating system.                                                                    |
| [`vaultify validate`](https://vaultify.buungroup.com/#/cli/VALIDATE.md) | Vaultify will validate your `terraform.tfstate` file JSON.                                                         |
| [`vaultify compare`](https://vaultify.buungroup.com/#/cli/COMPARE.md) | Vaultify will compare your local `terraform.tfstate` file JSON to your remote Vault `terraform.tfstate` file.  |
| [`vaultify update`](https://vaultify.buungroup.com/#/cli/UPDATE.md) | Update Vaultify to the latest version.                                                                           |
| [`vaultify wrap`](https://vaultify.buungroup.com/#/cli/WRAP.md) | Encrypts and encodes Terraform statefiles for secure storage in HashiCorp Vault.                                |
| [`vaultify unwrap`](https://vaultify.buungroup.com/#/cli/UNWRAP.md) | Decrypts and decodes Terraform statefiles, retrieving them from HashiCorp Vault for use.                       |
| [`vaultify delete`](https://vaultify.buungroup.com/#/cli/DELETE.md) | Delete the HashiCorp secret from Vault.                                                                         |
| [`vaultify path`](https://vaultify.buungroup.com/#/cli/PATH.md) | Display the HashiCorp secret path used to store statefiles.  
| [`vaultify retrieve`](https://vaultify.buungroup.com/#/cli/RETRIEVE.md) | Combines pull and unwrap together to speed up state existence.                                                   |
| [`vaultify publish`](https://vaultify.buungroup.com/#/cli/PUBLISH.md) | Combines wrap and push together to speed up pushing your state to Hashicorp Vault.                                                   |
| [`vaultify permissions`](https://vaultify.buungroup.com/#/cli/PERMISSIONS.md) | It will validate the policies on your token, then attempt to create a test secret on your engine type you have as default it will either suceed or fail.                                                   |
| [`vaultify pull`](https://vaultify.buungroup.com/#/cli/PULL.md) | Pulls encrypted data from HashiCorp Vault and decodes it, making it accessible for local use.                    |
| [`vaultify push`](https://vaultify.buungroup.com/#/cli/PUSH.md) | Pushes encrypted data, such as Terraform statefiles, into HashiCorp Vault, allowing for centralized and secure storage. |
| [`vaultify status`](https://vaultify.buungroup.com/#/cli/STATUS.md) | Checks if Vaultify is still authenticated to HashiCorp Vault.                                                     |
| [`vaultify configure`](https://vaultify.buungroup.com/#/cli/CONFIGURE.md) | Configures the Vaultify project, allowing customization of settings such as the Vault address, authentication method, and data paths. |
| [`vaultify -v, --version`](https://vaultify.buungroup.com/#/cli/VERSION.md) | Show the Vaultify version.                                                                                      |
| [`vaultify -h, --help`](https://vaultify.buungroup.com/#/cli/HELP.md)    | Show this help message.                                                                                         |

Each command has its specific use case, enabling you to seamlessly integrate Vaultify into your workflow for secure data management and collaboration. Detailed documentation for each command can be found in their respective sections linked above.

Let's dive into the details of each command and explore how they can be employed to enhance your interaction with HashiCorp Vault and Terraform statefiles.

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
