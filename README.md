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

Welcome to Vaultify, the versatile CLI tool designed to simplify the management of your statefiles while ensuring their security. Vaultify empowers DevOps teams and infrastructure engineers to seamlessly encrypt, store, and retrieve statefiles in HashiCorp Vault. Whether you're automating CI/CD pipelines or collaborating with a team, Vaultify provides the tools you need to protect sensitive data and streamline your workflows.

### Get Started

To begin using Vaultify, simply follow the installation instructions in the documentation. Once installed, you can explore its powerful features and incorporate it into your DevOps toolchain.

Experience the convenience of managing statefiles securely with Vaultify. Let's simplify DevOps together!

---

## Supported Operating Systems and Requirements

Vaultify is currently supported on the following operating system:

| Operating System | Supported Version |
| ---------------- | ----------------- |
| **Linux**            | Any distribution  |

### Requirements

Before using `Vaultify`, make sure your system meets the following requirements:

1. **Linux Operating System**: `Vaultify` is currently supported only on Linux-based operating systems.

2. **Dependencies**:
   - **curl**: `Vaultify` requires the `curl` command-line tool to interact with HashiCorp Vault. You can install it using your system's package manager.
   - **gzip**: The `gzip` utility is used for compressing state files. Ensure it is installed on your system.
   
3. **Terraform or Opentofu**: `Vaultify` expects either Terraform or Opentofu to be installed on your system. These are used for managing infrastructure and Terraform state files. Install one of these tools based on your needs.

Please ensure that you have these requirements fulfilled on your system before using `Vaultify`.

---

# Install Vaultify

Run the following commands for installation of Vaultify.
Binary
```bash
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

# Documentation

You can find more information on `VAULTIFY CLI` documentation at [DOCS](docs/CLI.md) file for details, this section covers the logic of each command and what exactly it is doing.

---

# Vaultify - A CLI Tool for Managing Statefiles

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
| `configure`        | Configures the Vaultify project, allowing customization of settings such as the Vault address, authentication method, and data paths    |
| `status`      | Checks if Vaultify is still authenticated to Hashicorp Vault.    |
| `-h, --help`        | Display the help display for Vaultify.    |
| `-v, --version`        | Display Vaultify version installed on the operating system.    |

## Usage

To use Vaultify, simply execute one of the commands listed above. For example, to initialize Vaultify, you can run:

```bash
vaultify <command>
```

---

## Bugs/Errors

Please keep in mind this product is brand new and has not been tested actively and throughly, if you encounter errors or issues please raise a PR or raise a issue, thank you.

---

## Unlocking the Power of Vaultify for DevOps Pipelines and Automation

Vaultify, a versatile CLI tool designed for managing statefiles securely, offers a wealth of capabilities that can greatly enhance DevOps pipelines and automation workflows. By seamlessly integrating Vaultify into your development and deployment processes, you can harness its power to store statefiles encrypted in HashiCorp Vault, thereby bolstering security, efficiency, and reliability across your entire software development lifecycle.

<details>
  <summary>1. Enhanced Security</summary>

Vaultify empowers DevOps teams to enhance the security of their statefiles by encrypting and storing them in HashiCorp Vault. This ensures that sensitive information and infrastructure configurations remain protected at rest. With Vaultify, secrets and statefiles are shielded from unauthorized access, reducing the risk of data breaches.
</details>

<details>
  <summary>2. Automated Workflows</summary>

Integrate Vaultify into your CI/CD pipelines to automate the encryption and storage of statefiles. By seamlessly incorporating Vaultify commands into your scripts, you can achieve consistency and reliability in managing statefiles across different environments. Automation ensures that every statefile is encrypted and stored securely without manual intervention.
</details>

<details>
  <summary>3. Version Control</summary>

Vaultify allows you to version control your statefiles efficiently. By wrapping and unwrapping secrets within the statefiles using base64 encoding, you can maintain a historical record of changes. This ensures that you can easily track, compare, and revert to previous versions of statefiles when necessary.
</details>

<details>
  <summary>4. Continuous Integration and Continuous Deployment (CI/CD)</summary>

Leverage Vaultify's capabilities within your CI/CD processes. Ensure that statefiles are encrypted before deployment and securely pushed to HashiCorp Vault. This guarantees that only authorized systems and personnel can access and retrieve statefiles, reducing the risk of unauthorized modifications or data exposure.
</details>

<details>
  <summary>5. Flexible Configuration</summary>

Vaultify's support for environment variables, such as VAULT_TOKEN and VAULT_ADDR, allows you to customize and adapt its behavior to various deployment scenarios. Whether you're working in a multi-environment setup or dealing with different Vault instances, Vaultify's flexibility accommodates your specific requirements.
</details>

<details>
  <summary>6. Error Detection and Validation</summary>

Utilize Vaultify's validate command to automatically check the JSON format of your Terraform statefiles. This built-in validation ensures that your statefiles are well-structured and error-free before deployment, reducing the risk of issues in production environments.
</details>

<details>
  <summary>7. Scaling and Collaboration</summary>

As your DevOps infrastructure scales and teams collaborate on projects, Vaultify remains a robust tool for managing secrets and statefiles securely. Each team member can easily use Vaultify to access and update statefiles in a standardized and secure manner, promoting efficient collaboration.
</details>

---

## Authentication with HashiCorp Vault

Vaultify requires two essential pieces of information for authentication with HashiCorp Vault: the Vault token (`VAULT_TOKEN`) and the Vault address (`VAULT_ADDR`). These authentication parameters are crucial for Vaultify to complete its actions securely and effectively. Below, we explain why each of these parameters is necessary:

### Vault Token (`VAULT_TOKEN`)

The Vault token is an authentication credential that provides access to HashiCorp Vault's resources and secrets. Vaultify uses this token to prove its identity to Vault and gain access to the necessary secrets and statefiles. Here's why the Vault token is required:

- **Authorization**: The Vault token acts as a key that grants permission to perform actions within Vault. Without a valid token, Vaultify cannot authenticate itself to Vault, and therefore, it won't be authorized to perform any operations.

- **Access Control**: HashiCorp Vault relies on access control policies associated with tokens. The token you provide to Vaultify must have the necessary permissions to perform the actions you intend to execute. Vaultify assumes the privileges of the token provided.

### Vault Address (`VAULT_ADDR`)

The Vault address specifies the location and endpoint of the HashiCorp Vault server that Vaultify should communicate with. It defines where Vaultify should send its requests and retrieve the required secrets and statefiles. Here's why the Vault address is necessary:

- **Endpoint Resolution**: HashiCorp Vault may be running on different servers or environments, each with its own Vault instance. The Vault address allows Vaultify to locate the correct Vault instance to connect to.

- **Network Communication**: Vaultify needs to establish a network connection to the specified Vault server. The Vault address ensures that Vaultify communicates with the right server, whether it's hosted locally or in a remote environment.

### Environment Variables

Vaultify relies on the following environment variables for configuration. You can set these variables to customize Vaultify's behavior:

| Description                         | Environment Variable | Required |
|-------------------------------------|----------------------|----------|
| Vault token for authentication     | `VAULT_TOKEN`        | Yes      |
| Vault server address for connection | `VAULT_ADDR`         | Yes  

In summary, the Vault token and Vault address are fundamental authentication parameters that Vaultify relies on to securely interact with HashiCorp Vault. Ensuring that these parameters are correctly configured is essential for Vaultify to perform its actions effectively while maintaining the security of your secrets and statefiles.

---

## Contributing

This section covers how to contrinute to this project  see the [CONTRIBUTING](CONTRIBUTING.md) file for details.

---

## Author

Vaultify is maintained by `Sacha Roussakis-Notter`.

## License

This project is licensed under the `GNU General Public License, Version 3 (GPL-3.0)` - see the [LICENSE](LICENSE) file for details.