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

# Vaultify CLI Documentation


# Overview

Vaultify is a powerful CLI tool developed in Go, designed to enhance productivity and security by encrypting the state-files in `base64` for managing Terraform state files. It streamlines the encryption and storage of state files across multiple platforms, including `HashiCorp Vault`, `Azure Storage Account`, and soon, `AWS S3 buckets`. By automating the encryption and push/pull processes, Vaultify ensures your Terraform state files are securely managed and easily accessible.

> NOTE: You can also refer to vaultify documentation at [Vaultify](https://vaultify.buungroup.com) to learn more.

In summary Vaultify significantly enhances the management of Terraform state files, offering a robust, secure, and cost-effective solution that goes beyond traditional methodologies. By supporting a variety of storage backends including HashiCorp Vault, Azure Storage, and future integrations with AWS S3, Vaultify provides a flexible and scalable approach to state file management. This versatility ensures that regardless of your infrastructure's complexity or the scale of your operations, Vaultify simplifies automation and streamlines workflows across multiple platforms.

<br>

---

<br>

# Why Vaultify?

Vaultify was developed to tackle the significant security and access management challenges observed in large organizations, where the complexity of roles and permissions often led to inadvertent access to sensitive Terraform state files in plaintext. Addressing the need for a secure, efficient, and centralized solution for state file management, Vaultify simplifies encryption processes and integrates seamlessly with major pipeline and DevOps platforms. It supports storage options like HashiCorp Vault and Azure Storage, overcoming Vault's 1MB limitation through gzip compression to securely store larger state files. This makes Vaultify an invaluable tool for organizations aiming to enhance security, streamline workflows, and manage permissions effectively, all while keeping infrastructure deployment and pipeline processes uncomplicated.

<details>
<summary>Click to read why vaultify was created indepth</summary>

Vaultify's creation was inspired by my extensive experience across large-scale organizations, where I encountered significant challenges in managing access rights and permissions for sensitive files. In these environments, the intricate overlay of roles and permissions frequently led to scenarios where individuals could inadvertently access state files in plaintext — a situation that should never occur. Moreover, the size and complexity of these organizations often meant that other teams' errors could introduce security vulnerabilities, sometimes remaining undetected until posing a tangible risk.

The objective behind developing Vaultify was to address these critical issues by providing a Command-Line Interface (CLI) tool specifically designed to simplify the encryption process of state files. This ensures that sensitive information is never left exposed in plaintext, thereby enhancing security without complicating the deployment infrastructure or the continuity of pipeline processes.

Vaultify is designed to seamlessly integrate with virtually any pipeline and DevOps platform provider, offering a versatile solution for secure state file management. The decision to include Vault as a supported storage option was motivated by the desire to centralize role-based access control (RBAC) mechanisms. Utilizing Vault allows for the consolidation of permission management in a single location, leveraging Vault's inherent security features and simplifying the oversight of access rights.

Furthermore, recognizing the potential cost implications of relying exclusively on HashiCorp Vault within enterprise licensing models, Vaultify also extends support to Azure Storage. This inclusion ensures that organizations can maintain the security of their state files, encrypted in base64, without incurring unnecessary expenses. By encrypting all state files, Vaultify significantly diminishes the risks associated with plaintext file storage and reduces potential attack surfaces, providing a comprehensive and secure state management solution.

The initial reluctance to adopt Vault for Terraform state file management was largely due to its 1MB size limit per secret, rendering it unsuitable for larger state files. This restriction posed a significant challenge for using Vault as a unified platform for state management, particularly as the size of state files tends to increase with the complexity of the infrastructure being managed.

Vaultify addresses this challenge head-on by integrating gzip compression into its operational workflow. This step significantly reduces the size of state files before they are encrypted. For example, a state file that originally sizes at 5MB can be compressed to a much smaller size, making it feasible to store within Vault's size limitations. It's important to note that while the example of compressing a 5MB file down to 200KB may be optimistic, the actual compression ratio can vary based on the content of the state files. Generally, gzip compression can result in substantial size reductions, making previously unmanageable files fit comfortably within Vault's constraints after encryption.

By overcoming Vault's size limitation, Vaultify enhances the practicality of storing Terraform state files securely within Vault. This advancement opens up avenues for organizations to centralize their infrastructure management practices, offering a secure, efficient, and consolidated solution for managing sensitive state data. Vaultify's approach not only navigates around the storage size hurdle but also capitalizes on the security and organizational benefits that Vault provides, presenting a compelling case for its adoption in managing Terraform state files across various scales of infrastructure projects.

</details>

<br>

### **Key Features of Vaultify**

| Feature                                      | Vaultify                                                                 | Traditional Methods                                           |
|----------------------------------------------|--------------------------------------------------------------------------|---------------------------------------------------------------|
| **Encryption of State Files**                | Automatically encrypts state files before storage.                       | State files stored in plain text or manually encrypted.       |
| **Decryption for Use**                       | Automatically decrypts on retrieval for immediate use.                   | Manual decryption required if encrypted.                      |
| **State File Management**                    | Centralized management for push, pull, and sync of state files.          | Relies on manual management or Terraform Cloud features.      |
| **Integration with Secret Managers**         | Native integration with HashiCorp Vault, Azure Storage, etc.             | Limited to Terraform Cloud's integration or manual setup.     |
| **Access Control and Permissions**           | Leverages existing secret manager's RBAC for unified access control.     | Managed separately within Terraform Cloud or storage backend. |
| **Version Control and History**              | Integrates with secret managers to utilize their versioning capabilities.| Dependent on backend capabilities or Terraform Cloud.         |
| **Cost Optimization**                        | Potentially reduces costs through efficient storage management.          | Costs can vary based on backend and Terraform Cloud pricing.  |
| **Simplified Workflow**                      | Streamlines the encryption and decryption process with simple commands.  | Often requires additional scripts or manual processes.        |
| **Customizable Configuration**               | Flexible `.vaultify/settings.json` for tailored workflows.               | Configuration limited to Terraform backend syntax.            |
| **DevOps and CI/CD Integration**             | Designed for seamless integration into any pipeline or DevOps platform.  | Requires custom integration or use of Terraform Cloud features.|
| **Scalability**                              | Built to efficiently scale with project complexity and size.             | Scalability can be constrained by backend or platform limits. |
| **Audit and Compliance**                     | Enhances compliance with encrypted storage and detailed audit trails.    | Compliance efforts may require additional tooling or processes.|

### **Current Supported State-file Secret Storage Solutions**

| Secret Storage Manager | Supported by Vaultify |
|------------------------|-----------------------|
| **HashiCorp Vault**        | ✓                     |
| **Azure Storage**          | ✓                     |
| **S3**                     | ✗                     |

<br>

---

<br>

# Installation 🔨

To ensure a smooth experience with Vaultify, we highly recommend reviewing the comprehensive documentation available on the [Vaultify](https://vaultify.buungroup.com) website. The following guidelines will help you prepare your environment for using Vaultify effectively.

  ##### **System Requirements**

  Vaultify has been designed for Linux-based operating systems, underscoring our commitment to security and flexibility. Before proceeding, confirm that your system aligns with this requirement to ensure compatibility.

  ##### **Required Dependencies**

  For Vaultify to function as intended, several dependencies must be present on your system:
    
  - `curl`: Essential for communicating with HashiCorp Vault, curl can be easily installed through your distribution's package manager.
  - `gzip`: Utilized for the compression of state files, `gzip` should be installed to manage file sizes efficiently.
  - `jq`: A powerful tool for querying and manipulating JSON data, `jq` is necessary for processing configuration and state files.

To streamline the installation of these dependencies, you may opt to use the make command, which automates the setup process. Please ensure make is installed on your system before proceeding. If needed, make can be installed using the following command:

```bash
sudo apt-get install make -y
```

> Note: Please install this using your own system distrbution package manager.

After installing make, you can execute the following command to install all required dependencies in one go:

```bash
make
```

After ensuring your system meets the prerequisites and dependencies for `Vaultify`, the next step is to install the tool itself. `Vaultify` simplifies this process with a straightforward installation script that can be executed using either `curl` or `wget`. This script installs `Vaultify` and prepares your system for managing Terraform state files with enhanced security and efficiency.


**Using curl**:
```bash
curl --proto '=https' --tlsv1.2 -fsSL https://raw.githubusercontent.com/DFW1N/vaultify/main/scripts/install-vaultify.sh | sudo bash
```

**Using wget**:
```bash
wget --https-only -O - https://raw.githubusercontent.com/DFW1N/vaultify/main/scripts/install-vaultify.sh | sudo bash
```

#### Optional: Deploying HashiCorp Vault with Docker using Vaultify

<details>
<summary>Click to view the Optional: Deploying HashiCorp Vault with Docker</summary>


### Optional: Deploying HashiCorp Vault with Docker

If your workflow requires deploying HashiCorp Vault, you can use our Docker setup for a quick and easy deployment. This step is optional and only necessary if you need a local instance of HashiCorp Vault for managing your secrets if you want to use the `vault`, secret engine to manage your `terraform.tfstate` files.

#### Prerequisites

- **Docker:** Ensure Docker is installed on your system. The installation script provided supports Ubuntu/Debian and Fedora/Red Hat distributions. If you're using a different OS, please follow the Docker installation instructions for your specific system.

  ##### Installation Steps:

  - Install Docker (if not already installed): Use the following command to install Docker on your system. This script is tailored for Ubuntu/Debian and Fedora/Red Hat distributions.

    ```bash
    curl --proto '=https' --tlsv1.2 -fsSL https://raw.githubusercontent.com/DFW1N/vaultify/main/docker/scripts/install-docker.sh | sudo bash
    ```

    > Feel free to view Buun Group [Docker Hub](https://hub.docker.com/r/buungroup/vault-raft) for version details.

  - Deploy HashiCorp Vault: After setting up Docker, deploy your local HashiCorp Vault instance using Vaultify's installation command. This will set up a Vault server tailored for use with Vaultify.

    ```bash
    vaultify install-vault
    ```

</details>

<br>

---

<br>

# Getting Started with Vaultify

Following the successful installation of Vaultify and ensuring that your system meets all prerequisites, you're now ready to embark on your journey of secure and efficient Terraform state file management. This section will guide you through the initial steps to get up and running with Vaultify.

#### **1. Initial Configuration**

Vaultify operates with a blend of configuration files and environment variables. Begin by initializing Vaultify's configuration:

```bash
vaultify init
```

This command creates a .vaultify/settings.json file in your home directory, which serves as the base configuration for Vaultify. You may need to edit this file to specify your secret storage options and other preferences.

#### 2. **Setting Environment Variables**

Depending on your chosen secret storage provider, set the necessary environment variables in your shell:

**For Azure Storage**:

```bash
export ARM_SUBSCRIPTION_ID="your_subscription_id"
export ARM_CLIENT_ID="your_client_id"
export ARM_CLIENT_SECRET="your_client_secret"
export ARM_TENANT_ID="your_tenant_id"

```

**For Vault:**

```bash
export VAULT_TOKEN="your_vault_token"
export VAULT_ADDR="your_vault_address"

```

#### 3. ⭐ Usage

To use Vaultify, simply execute one of the commands listed in the table below.

```bash
vaultify <command>
```


Ensure these variables are set before proceeding to use Vaultify for managing your state files, to learn more about environment variables please review the section below `[Configuration and Authentication]` for more detail.

<br>

---

<br>

# Configuration and Authentication

Vaultify uses the `.vaultify/settings.json` file and environment variables for configuration and authentication with your secret management provider. Start by running `vaultify init` to generate the initial settings file, which outlines how Vaultify will manage your Terraform state files.

Adjust this file to fit your project's needs and set the necessary environment variables for secure authentication with your chosen secret provider. This setup ensures a secure and efficient management of your infrastructure's state files.

---

#### Required environment variables for supported `secret providers`

<details>
<summary>Click to view required environment variables</summary>

| Secret Provider   | Required Environment Variables                                      |
|-------------------|---------------------------------------------------------------------|
| **Azure Storage**     | `ARM_SUBSCRIPTION_ID`, `ARM_CLIENT_ID`, `ARM_CLIENT_SECRET`, `ARM_TENANT_ID` |
| **Vault**             | `VAULT_TOKEN`, `VAULT_ADDR`
| **AWS S3**             | `NOT SUPPORTED YET`

> Note: These can be easily exported in your operating system using `export VAULT_ADDR=http://localhost:8200`. If you don't have an hashicorp vault already setup.

</details>

---

#### Required table fields for `settings.json` file.

<details>
<summary>Click to view the table fields for settings.json</summary>

| Field                                       | Type          | Default Value         | Accepted Values                                               | Description                                           |
|---------------------------------------------|---------------|-----------------------|---------------------------------------------------------------|-------------------------------------------------------|
| `TerraformWorkspace`                        | `bool`        | `true`                | `true`, `false`                                               | Indicates whether to use Terraform workspace name.         |
| `DefaultEngineName`                         | `string`      | `"kv"`                | Any string, typically secret engines supported by `vault`           | Default engine name for the storage backend.          |
| `DefaultSecretStorage`                      | `string`      | `"vault"`             | `"vault"`, `"azure_storage"`, `"s3"` (future support for AWS S3)    | Default secret storage provider.                      |
| `Azure.StorageAccountName`                  | `string`      | `""` (empty string)   | Any valid Azure Storage Account required if name `DefaultSecretStorage` = `"azure_storage"`                         | Azure storage account name.                           |
| `Azure.StorageAccountResourceGroupName`     | `string`      | `""` (empty string)   | Any valid Azure Resource Group required if name `DefaultSecretStorage` = `"azure_storage"`  name                           | Azure storage account resource group name.            |
| `AWS.S3BucketName`                          | `string`      | `""` (empty string)   | Any valid AWS S3 bucket name                                  | AWS S3 bucket name.                                   |

<br>

This configuration allows, you to have a custom `settings.json` commited and include it into your devops or pipeline deployments. 
<details>
<summary>Click to view the settings.json file format</summary>

```json
{
  "Settings": {
    "TerraformWorkspace": true,
    "DefaultEngineName": "kv",
    "DefaultSecretStorage": "vault",
    "Azure": {
      "StorageAccountName": "",
      "StorageAccountResourceGroupName": ""
    },
    "AWS": {
      "S3BucketName": ""
    }
  }
}
```
</details>

</details>



<br>

---


<br>

# 📖 Documentation

Explore the extensive [Vaultify documentation](https://vaultify.buungroup.com/#/CLI) for a deep dive into the CLI's functionalities, command explanations, and operational logic, ensuring you leverage Vaultify to its fullest potential.

---

## 🔐 Vaultify - Your Multi-Provider Statefile Manager

Vaultify stands out as an advanced command-line interface (CLI) tool engineered for the efficient management of statefiles. Designed with Terraform projects in mind this works perfectly with `Opentofu` also, Vaultify extends its capabilities to accommodate a broad range of secret management providers, including `HashiCorp Vault`, `Azure Storage`, and more. This versatility facilitates secure encryption, as well as the seamless push and pull of state files across diverse environments, streamlining your infrastructure management processes.

## Commands

Vaultify supports the following commands:

| Command                                 | Description                                |
|-----------------------------------------|--------------------------------------------|
| [`vaultify init`](https://vaultify.buungroup.com/#/cli/INIT.md) | Initialize Vaultify in your operating system.                                                                    |
| [`vaultify validate`](https://vaultify.buungroup.com/#/cli/VALIDATE.md) | Vaultify will validate your `terraform.tfstate` file JSON.                                                         |
| [`vaultify compare`](https://vaultify.buungroup.com/#/cli/COMPARE.md) | Vaultify will compare your local `terraform.tfstate` file JSON to your remote Vault `terraform.tfstate` file.  |
| [`vaultify update`](https://vaultify.buungroup.com/#/cli/UPDATE.md) | Update Vaultify to the latest version.                                                                           |
| [`vaultify wrap`](https://vaultify.buungroup.com/#/cli/WRAP.md) | Encrypts and encodes Terraform statefiles for secure storage in HashiCorp Vault.                                |
| [`vaultify unwrap`](https://vaultify.buungroup.com/#/cli/UNWRAP.md) | Decrypts and decodes Terraform statefiles, retrieving them from HashiCorp Vault for use.                       |
| [`vaultify delete`](https://vaultify.buungroup.com/#/cli/DELETE.md) | Delete the HashiCorp secret from Vault.                                                                         |
| [`vaultify path`](https://vaultify.buungroup.com/#/cli/PATH.md) | Display the HashiCorp secret path used to store statefiles.  
| [`vaultify retrieve`](https://vaultify.buungroup.com/#/cli/RETRIEVE.md) | Combines pull and unwrap together to speed up state existence.                                                   |
| [`vaultify permissions`](https://vaultify.buungroup.com/#/cli/PERMISSIONS.md) | It will validate the policies on your token, then attempt to create a test secret on your engine type you have as default it will either suceed or fail.                                                   |
| [`vaultify publish`](https://vaultify.buungroup.com/#/cli/PUBLISH.md) | Combines wrap and push together to speed up pushing your state to Hashicorp Vault.                                                   |
| [`vaultify pull`](https://vaultify.buungroup.com/#/cli/PULL.md) | Pulls encrypted data from HashiCorp Vault and decodes it, making it accessible for local use.                    |
| [`vaultify push`](https://vaultify.buungroup.com/#/cli/PUSH.md) | Pushes encrypted data, such as Terraform statefiles, into HashiCorp Vault, allowing for centralized and secure storage. |
| [`vaultify status`](https://vaultify.buungroup.com/#/cli/STATUS.md) | Checks if Vaultify is still authenticated to HashiCorp Vault.                                                     |
| [`vaultify install-vault`](https://vaultify.buungroup.com/#/cli/INSTALL-VAULT.md) | Vaultify will automatically setup and deploy your developer Hashicorp Vault                                                     |
| [`vaultify delete-vault`](https://vaultify.buungroup.com/#/cli/DELETE-VAULT.md) | Vaultify will automatically delete your developer or test Hashicorp Vault                                                     |
| [`vaultify configure`](https://vaultify.buungroup.com/#/cli/CONFIGURE.md) | Configures the Vaultify project, allowing customization of settings such as the Vault address, authentication method, and data paths. |
| [`vaultify -v, --version`](https://vaultify.buungroup.com/#/cli/VERSION.md) | Show the Vaultify version.                                                                                      |
| [`vaultify -h, --help`](https://vaultify.buungroup.com/#/cli/HELP.md)    | Show this help message.                                                                                         |

<br>

---

<br>

# 💣 Bugs/Errors

Please keep in mind this product is brand new and has not been tested actively and throughly, if you encounter errors or issues please raise a PR or raise a issue refer to [ISSUE TEMPLATE](.github/ISSUE_TEMPLATE/vaultify_issue_report.md) for more information on submission, thank you.

<br>

---

<br>

# ⌛ Contributing

This section covers how to contrinute to this project  see the [CONTRIBUTING](CONTRIBUTING.md) file for details.

<br>

---

<br>

# 🔥 Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |

<br>

---

<br>

# 📃 License

This project is licensed under the `GNU General Public License, Version 3 (GPL-3.0)` - see the [LICENSE](LICENSE) file for details.