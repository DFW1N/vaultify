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

- `vaultify wrap`: Encrypts and encodes Terraform statefiles for secure storage in HashiCorp Vault.

- `vaultify unwrap`: Decrypts and decodes Terraform statefiles, retrieving them from HashiCorp Vault for use.

- `vaultify push`: Pushes encrypted data, such as Terraform statefiles, into HashiCorp Vault, allowing for centralized and secure storage.

- `vaultify pull`: Pulls encrypted data from HashiCorp Vault and decodes it, making it accessible for local use.

- `vaultify init`: Initializes the Vaultify project, setting up configuration files and authentication with HashiCorp Vault.

- `vaultify configure`: Configures the Vaultify project, allowing customization of settings such as the Vault address, authentication method, and data paths.

- `vaultify status`: Checks the status of the Vaultify project and its connection to HashiCorp Vault.

- `vaultify version`: Displays the current version of the Vaultify CLI.

Each command has its specific use case, enabling you to seamlessly integrate Vaultify into your workflow for secure data management and collaboration. Detailed documentation for each command can be found in their respective sections below.

Let's dive into the details of each command and explore how they can be employed to enhance your interaction with HashiCorp Vault and Terraform statefiles.


---

<details>
<summary><strong>Vaultify Push</strong></summary>

The `vaultify push` command is part of the Vaultify project, which allows you to easily push, pull, and encrypt Terraform statefiles and other sensitive data to and from HashiCorp Vault.

## Usage

To use the `vaultify push` command, follow these steps:

1. First, make sure you have created an encoded Terraform state file using the `vaultify wrap` command. This encoded file will be used for pushing to HashiCorp Vault.

2. Navigate to the directory containing the `.encoded_wrap` file.

3. Run the `vaultify push` command.

## What It Does

The `vaultify push` command performs the following steps:

1. Checks for the existence of the `.encoded_wrap` file in the current directory. If the file is not found, it displays an error message and exits.

2. Reads the contents of the `.encoded_wrap` file, which is assumed to contain the encoded Terraform state.

3. Validates that the content is in a valid base64 format.

4. Sets an environment variable named `TERRAFORM_STATE_BASE64` with the content of the encoded Terraform state.

5. Constructs a URL to HashiCorp Vault using the `VAULT_ADDR` environment variable.

6. Determines the secret path in Vault where the encoded state will be stored. This path is constructed as follows:
   - `dataPath`: A configurable data path (default is "vaultify").
   - `workspaceName`: The current Terraform workspace name (obtained using `terraform workspace show`).
   - `workingDirName`: The name of the current working directory (the folder where the command is executed).
   - Additional "_terraform.tfstate" is appended to `workingDirName` to create the final path.

7. Checks if the secret path already exists in HashiCorp Vault. If not, it creates the path.

8. Pushes the encoded state to Vault under the calculated secret path.

9. Displays the status of the push operation, including the path where the state was stored, the size of the uploaded file, and any additional messages.

</details>

---

<details>
<summary><strong>Vaultify Wrap</strong></summary>

# Vaultify Wrap Command

The `vaultify wrap` command is part of the Vaultify project, designed to simplify the process of encrypting and encoding Terraform statefiles for secure storage in HashiCorp Vault.

## Usage

To use the `vaultify wrap` command, follow these steps:

1. Navigate to the directory containing the Terraform statefile (`*.tfstate`) you want to wrap.

2. Run the `vaultify wrap` command.

3. The wrapped and encoded statefile will be saved in the current directory as `.encoded_wrap`.

## What It Does

The `vaultify wrap` command performs the following logical steps:

1. Searches for Terraform statefiles (`*.tfstate`) in the current directory and selects the first found file.

2. Reads the contents of the selected Terraform statefile.

3. Compresses the statefile using the gzip compression algorithm.

4. Encodes the compressed statefile into base64 format.

5. Sets an environment variable named `TERRAFORM_STATE_BASE64` with the encoded statefile data, making it accessible to other Vaultify commands.

6. Saves the encoded statefile to a file named `.encoded_wrap` in the current directory.

7. Displays a success message indicating that the statefile has been wrapped and saved.

The wrapped statefile in `.encoded_wrap` is now ready for secure storage in HashiCorp Vault using the `vaultify push` command.

## Example

Suppose you have a Terraform statefile named `terraform.tfstate` in your current directory. Running the `vaultify wrap` command will perform the following actions:

1. Compress the `terraform.tfstate` file.

2. Encode the compressed data into base64 format.

3. Save the encoded data to a file named `.encoded_wrap`.

After running the command, you'll have a file named `.encoded_wrap` that contains the encoded and compressed Terraform state, which can be securely pushed into HashiCorp Vault for storage and later retrieval.

## Notes

- The `vaultify wrap` command operates on the first `.tfstate` file found in the current directory. Ensure that you are in the correct directory with the desired statefile before running the command.

- To use the wrapped statefile with other Vaultify commands, make sure to have the `TERRAFORM_STATE_BASE64` environment variable correctly set, as it holds the encoded state data.

- You can customize the behavior of Vaultify and the data path used for storage by configuring settings with the `vaultify configure` command.

For detailed information on the Vaultify project and other available commands, refer to the project's documentation and README.

</details>

---

<details>
<summary><strong>Vaultify Unwrap</strong></summary>

</details>

---

<details>
<summary><strong>Vaultify Init</strong></summary>

</details>

---

<details>
<summary><strong>Vaultify Validate</strong></summary>

</details>

---

<details>
<summary><strong>Vaultify Pull</strong></summary>

</details>

---