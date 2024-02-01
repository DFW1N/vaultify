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

# Vaultify CLI Command - Install-Vault

## Overview
The `install-vault` command in the Vaultify CLI provides a easy way to deploy a developer or test `Hashicorp Vault` locally, so you can play or test with `Vaultify` before releasing it into your production pipelines or production environments.

## Functionality
- **Display Command List:**
  The `install-vault` command utilizes the operating system running `Vaultify`, CLI it checks it the operating system, has the `Docker` engine installed and takes advantage of it by deploying vault, on docker. It will automatically install `Docker` if it does not detect it. It will automatically set the required `Environment variables` inside the container which comes with `Vaultify` pre-installed. It `VAULT_TOKEN` and `VAULT_ADDR`. Based off the outputs from the command.

## Limitations

You will have to manually update `Vaultify` versions, inside the container if you do not want to destroy and pull from new docker tag.

## Finding your Token:

To find your `Hashicorp Vault Token` simply run the following commands:

```bash
docker exec -it vault-raft-backend /bin/bash
vaultToken=$(grep VAULT_ADDR /root/.bashrc | cut -d'=' -f2 | tr -d "'")
echo $vaultToken
```

## Usage
To use the `install-vault` command in the Vaultify CLI, run the following command:
This install a developer `Hashcorp Vault`, using `RAFT` storage backend.

```bash
vaultify install-vault
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
