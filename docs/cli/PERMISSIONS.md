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

# Vaultify CLI Command - Permissions

## Overview
The `permissions` command in the Vaultify CLI has been created to validate if you have the relevant, permissions on your selected secret engine, to create and publish secrets.

## Functionality
- **Executes lookup on token:**
  It goes up to vault and looks up the token policies at: `"/v1/auth/token/lookup-self"`
- **Create Test Token**
  It will then validate by attempting to create a secret on your selected engine Type, for example `kv` if it fails it will output if it lacks permissions or output suceed and will automatically delete and clean up the test secret which gets published to the directory `vaultAddr + "/v1/" + engineName + "/data/`.


To use the `permissions` command in the Vaultify CLI, run the following command:

```bash
vaultify permissions
```

---

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
