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
The `inject` command in the Vaultify CLI is designed to securely store secrets in HashiCorp Vault. It allows users to specify a custom path, key, and value for the secret, or use default settings based on the current workspace. The command now supports reading secret values from files and specifying metadata.

## Functionality
- **Parse Command-Line Flags:**
  The function parses flags for the Vault path, secret key, metadata, and secret file.

- **Check Vaultify Setup:**
  Ensures that Vaultify is properly set up and Vault is initialized.

- **Determine Secret Path:**
  Uses the provided path or generates a default path based on the current workspace and directory.

- **Read Secret Value:**
  Reads the secret value from a file if specified, or uses the provided inline value.

- **Parse Metadata:**
  Parses metadata from a JSON string or file if provided.

- **Ensure KV Path Exists:**
  Checks if the specified path exists in Vault and creates it if necessary.

- **Write Secret to Vault:**
  Writes the provided secret to Vault at the specified path and key, along with any metadata.

- **Encrypt Secret:** 
  Encrypts the secret using the VAULTIFY_PASSPHRASE.

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

### With Metadata (inline JSON)
```bash
vaultify inject -path "foo/bar/config" -metadata '{"owner":"john","project":"vaultify"}' "my_secret_token"
```

### With Metadata (from JSON file)
```bash
vaultify inject -path "foo/bar/config" -metadata metadata.json "my_secret_token"
```

### Secret Value from File
```bash
vaultify inject -path "foo/bar/config" -secret-file /path/to/secret/file
```

### Secret Value from File with Metadata File
```bash
vaultify inject -secret-file example.json -path "foo/bar/config" -metadata metadata.json
```

- `-path`: Optional. Specifies the Vault path to store the secret. If not provided, a default path based on the current workspace is used.
- `-key`: Optional. Specifies the key for the secret. Defaults to "default" if not provided.
- `-metadata`: Optional. Specifies metadata for the secret. Can be an inline JSON string or a path to a JSON file.
- `-secret-file`: Optional. Specifies a file path to read the secret value from.
- `<secret_value>`: Required if `-secret-file` is not used. The secret value to be stored in Vault.

---

## Metadata

Metadata can be provided either as an inline JSON string or as a path to a JSON file. It allows you to attach additional information to your secrets, such as owner, project, or any other relevant data.

### Example Metadata File

**metadata.json**
```json
{
  "owner":"sacha",
  "project":"vaultify"
}
```

## Hashicorp Vault

This will publish metadata to your secret that looks like this in VAULT UI.

**Custom metadata**

| Key   | Value |
|-------------------|-------------|
| **owner**     | `sacha` |
| **project**   | `vaultify` |

## Azure Storage Account

This will publish metadata to your secret that looks like this in your Azure Storage Account Blob.

> NOT YET SUPPORTED UNDER DEVELOPMENT

---

### Azure Storage Account Blobs

The `inject` command acts abit differently, for Azure Storage Account Blobs, the command 

```bash
vaultify inject -path "foo/bar/tfc_token" "my_secret_token"
```

This splits the -path into the following logic:

`foo/` becomes the name of a new container thats going to be created. While `bar` becomes a directory inside, the `foo` container and creates a new secret inside the `bar/` directory with a blob file name of `tfc_token`. Refer to the image to understand it better.

![alt text](/docs/img/azure-storage-account-blob-injection-logic-example-1.png)

This approach allows for a dynamic approach to controlling and setting fine grained RBAC access on Azure blobs that can adapt to any enterprise work environment with use of directory structures to split secrets or terraform statefiles. Keep in mind even though the secret is stored in blobs, they are encrypted and even if downloaded and accessed they cannot decrypt the value without knowing the `VAULTIFY_PASSPHRASE` used to inject the secret.

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |