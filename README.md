<!-- // ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														                            #
// # Project: Vaultify																	                                  #
// # Description: Easily push, pull and encrypt tofu and terraform statefiles from Vault. #
// ######################################################################################## -->

<div align="center">
    <img src="img/vaultify-logo.png" alt="Vaultify Logo" style="width: 30%;"/>
</div>


```bash
██╗   ██╗ █████╗ ██╗   ██╗██╗  ████████╗██╗███████╗██╗   ██╗
██║   ██║██╔══██╗██║   ██║██║  ╚══██╔══╝██║██╔════╝╚██╗ ██╔╝
██║   ██║███████║██║   ██║██║     ██║   ██║█████╗   ╚████╔╝ 
╚██╗ ██╔╝██╔══██║██║   ██║██║     ██║   ██║██╔══╝    ╚██╔╝  
 ╚████╔╝ ██║  ██║╚██████╔╝███████╗██║   ██║██║        ██║   
  ╚═══╝  ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝   ╚═╝╚═╝        ╚═╝   
                                                            
```

# Introduction 🙋

Welcome to Vaultify, the versatile CLI tool designed to simplify the management of your statefiles while ensuring their security. Vaultify empowers DevOps teams and infrastructure engineers to seamlessly encrypt, store, and retrieve statefiles in HashiCorp Vault. Whether you're automating CI/CD pipelines or collaborating with a team, Vaultify provides the tools you need to protect sensitive data and streamline your workflows.

To learn more you can visit the [Vaultify](https://vaultify.buungroup.com) website.

> NOTE: Documentation on this repository does not get updated please refer to [Vaultify](https://vaultify.buungroup.com) for the latest documentation.

In summary, `Vaultify` empowers you to optimize your Terraform state management, reducing costs, enhancing security, and simplifying automation, all while eliminating limitations imposed by traditional approaches. Reduce the complexity by running `Vaultify` as your `Hashicorp Vault` state manager.

### Get Started

Please refer to [Vaultify](https://vaultify.buungroup.com) for documentation.

Before using `Vaultify`, make sure your system meets the following requirements:

1. **Linux Operating System**: `Vaultify` is currently supported only on Linux-based operating systems.

2. **Dependencies**:
   - **curl**: `Vaultify` requires the `curl` command-line tool to interact with HashiCorp Vault. You can install it using your system's package manager.
   - **gzip**: The `gzip` utility is used for compressing state files. Ensure it is installed on your system.
   - **jq**: The `jq` utility is used for querying json data. Ensure it is installed on your system.

Alternatively you can run the `make` command, to install all the requirements.

Before you can run this `make` must be installed this can be installed with the following command:

```bash
sudo apt-get install make -y
```

Command Usage:
```bash
make
```

3. **Terraform or Opentofu**: `Vaultify` expects either Terraform or Opentofu to be installed on your system. These are used for managing infrastructure and Terraform state files. Install one of these tools based on your needs.

Please ensure that you have these requirements fulfilled on your system before using `Vaultify`.

---

# Install Vaultify 🔨

Run the following commands for installation of Vaultify.
Binary
```bash
sudo apt-get install jq -y
wget -qO- https://raw.githubusercontent.com/DFW1N/vaultify/main/public.key | gpg --import && gpg --verify vaultify.sig vaultify
latestVersion=$(curl -s "https://api.github.com/repos/DFW1N/vaultify/releases/latest" | jq -r '.tag_name'); wget -q "https://github.com/DFW1N/vaultify/releases/download/$latestVersion/vaultify" && chmod +x vaultify
sudo mv vaultify /usr/local/bin/
```

> NOTE: If you still cant run `vaultify` commands try refreshing your terminal.

`bin/bash`
```bash
source $HOME/.bashrc
```

`zsh`
```bash
source $HOME/.zshrc
```

---

# Documentation 📖

You can find more information on `VAULTIFY CLI` documentation at [Vaultify](https://vaultify.buungroup.com/#/CLI) file for details, this section covers the logic of each command and what exactly it is doing.

---

## Vaultify - A CLI Tool for Managing Statefiles 🔐

Vaultify is a command-line interface (CLI) tool for managing statefiles, particularly for Terraform projects. It provides functionality to interact with HashiCorp Vault, including secret encryption and pushing/pulling state from a remote Vault server.

## Commands

Vaultify supports the following commands:

| Command       | Description                                       |
|---------------|---------------------------------------------------|
| `init`        | Initialize Vaultify in your operating system.     |
| `validate`    | Validate the JSON format of your Terraform statefile. |
| `compare`     | Compare the Terraform statefile from local to whats in your vault. |
| `delete`      | Delete your remote terraform statefile in Hashicorp Vault. |
| `path`        | Display your statefile Hashicorp Vault secret path. |
| `update`      | Update the Vaultify CLI to the latest version.    |
| `wrap`        | Wrap a secret in base64 encoding.                |
| `unwrap`      | Unwrap a secret from base64 encoding.            |
| `pull`        | Pull state from a remote HashiCorp Vault server.  |
| `push`        | Push state to a remote HashiCorp Vault server.    |
| `retrieve`    | Combines pull and unwrap together to speed up state existence.    |
| `configure`        | Configures the Vaultify project, allowing customization of settings such as the Vault address, authentication method, and data paths    |
| `status`      | Checks if Vaultify is still authenticated to Hashicorp Vault.    |
| `-h, --help`        | Display the help display for Vaultify.    |
| `-v, --version`        | Display Vaultify version installed on the operating system.    |

## Usage ⭐

To use Vaultify, simply execute one of the commands listed above. For example, to initialize Vaultify, you can run:

```bash
vaultify <command>
```

> NOTE: Vaultify uses your `Terraform Workspace` name and `Working Directory Base` name to dynamically build your vault secret path to publish the encoded state file to.

### Example Usage

![Vaultify Command Examples](docs/img/example-useage.png)

In this example, I'm simply pulling an existing state file, in my `Hashicorp Vault`, unwrapping it then ideally you would run your `terraform` commands here update the statefile then run `vaultify wrap` then `vaultify push` to renew your state with the changes in your `Hashicorp Vault`.

---
---

## Bugs/Errors 💣

Please keep in mind this product is brand new and has not been tested actively and throughly, if you encounter errors or issues please raise a PR or raise a issue refer to [ISSUE TEMPLATE](.github/ISSUE_TEMPLATE/vaultify_issue_report.md) for more information on submission, thank you.

---

# Required Environment Variables

| Description                         | Environment Variable | Required |
|-------------------------------------|----------------------|----------|
| Vault token for authentication     | `VAULT_TOKEN`        | Yes      |
| Vault server address for connection | `VAULT_ADDR`         | Yes  

In summary, the Vault token and Vault address are fundamental authentication parameters that Vaultify relies on to securely interact with HashiCorp Vault. Ensuring that these parameters are correctly configured is essential for Vaultify to perform its actions effectively while maintaining the security of your secrets and statefiles.

---

## Contributing ⌛

This section covers how to contrinute to this project  see the [CONTRIBUTING](CONTRIBUTING.md) file for details.

---

## Author 🔥

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |

---

## License 📃

This project is licensed under the `GNU General Public License, Version 3 (GPL-3.0)` - see the [LICENSE](LICENSE) file for details.