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

---

# Introduction

This document has been created to showcase, automation solutions that can be used with vaultify.

---

# Azure DevOps

This section shows how you would pass the required environment variables if you want to use vaultify within your self-hosted agents or Microsoft cloud agents.

| Command     | Templated Task                                      | Usage                                         |
|-------------|------------------------------------------------|-----------------------------------------------|
| `command1`  | [View Template](#install-vaultify)                    | [View Usage](#usage1)                         |
| `command2`  | [View Template](#template2)                    | [View Usage](#usage2)                         |
| `command3`  | [View Template](#template3)                    | [View Usage](#usage3)                         |


<details>
<summary>Click to view pipeline templated examples</summary>

<br>

## Install Vaultify

`install_vaultify.yml`
```yml
parameters:
  defaultWorkingDirectory: ''

steps:

####################
# Install Vaultify #
####################

- task: Bash@3
  displayName: "Install > Vaultify"
  continueOnError: false
  enabled: true
  inputs:
    targetType: 'inline'
    workingDirectory: $(System.DefaultWorkingDirectory)
    script: |
        if ! [ -x "$(command -v vaultify)" ]; 
        then
            echo "⚠️ Installing Vaultify..."
            curl --proto '=https' --tlsv1.2 -fsSL https://raw.githubusercontent.com/DFW1N/vaultify/main/scripts/install-vaultify.sh | sudo bash
            vaultify --version
          else
            echo "✅ Vaultify is already installed."
            vaultify --version
          fi
```



---

<br>

`vaultify_retrieve.yml`

```yml
parameters:
  defaultWorkingDirectory: ''
  vaultToken: ''
  vaultServerUrl: ''
  subscriptionId: ''
  clientId: ''
  clientSecretId: ''
  tenantId: ''

steps:

####################
# Install Vaultify #
####################

- task: Bash@3
  displayName: "Vaultify > Retrieve"
  continueOnError: false
  enabled: true
  inputs:
    targetType: 'inline'
    workingDirectory: ${{ parameters.defaultWorkingDirectory }}
    script: |
        vaultify retrieve
  env:
    VAULT_ADDR:          ${{ parameters.vaultServerUrl }}
    VAULT_TOKEN:         ${{ parameters.vaultToken }}
    ARM_SUBSCRIPTION_ID: ${{ parameters.subscriptionId }}
    ARM_CLIENT_ID:       ${{ parameters.clientId }}
    ARM_CLIENT_SECRET:   ${{ parameters.clientSecretId }}
    ARM_TENANT_ID:       ${{ parameters.tenantId }}
```

</details>


## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
