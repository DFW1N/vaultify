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

This document has been created to showcase, how you can intergrate `Vaultify`, into your pipelines for automation.

---

<br>

# Azure DevOps

This section covers how you create tasks, to automate pipelines with `Vaultify`, it has been broken down into two sections, one being the templated task, what it is actually doing using `bash` and the second part actually using the templated task and passing the required `parameters` to it.

<br>

`vaultify_retrieve.yml`
### Part 1, templated task.

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

#####################
# Vaultify Retrieve #
#####################

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

> `Note:` Save this templated task as a file `vaultify_retrieve.yml`, somewhere in your repository then point the template to the path of this file.

<br>

---

### Part 2, template usage.

```yml
#####################
# Vaultify Retrieve #
#####################

- template: <relative-directory>/vaultify_retrieve.yml
    parameters:
    defaultWorkingDirectory: ${{ parameters.defaultWorkingDirectory }}
    vaultToken: ${{ parameters.vaultToken }}
    vaultServerUrl: ${{ parameters.vaultServerUrl }}
    clientSecretId: ${{ parameters.clientSecretId }}
    subscriptionId: ${{ parameters.subscriptionId }}
    tenantId: ${{ parameters.tenantId }}
    clientId: ${{ parameters.clientId }}
```

> `Note:` Please make sure, you are passing these parameters, from Azure keyvault or your secret storage manager to pass it securely.

---

<br>

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
