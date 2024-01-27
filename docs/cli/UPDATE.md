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

# Vaultify CLI Command - Update

## Overview
The `update` command in the Vaultify CLI allows you to update the Vaultify tool to the latest version available on GitHub. It checks for updates by comparing the installed version with the latest release tag from the official GitHub repository.

## Functionality
- **Check Installed Version:**
  The `update` command first checks the version of the currently installed Vaultify binary using the `vaultify --version` command.

- **Check Latest Release:**
  It then queries the GitHub API to fetch the latest release tag from the official Vaultify repository.

- **Update if Needed:**
  If the installed version is older than the latest release, the command will:
  - Download the latest Vaultify binary from the official GitHub release.
  - Make the downloaded binary executable.
  - Move the binary to the `/usr/local/bin` directory (requires sudo privileges).
  - Report the successful update.

## Usage
To use the `update` command in the Vaultify CLI, run the following command:

```bash
vaultify update
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
