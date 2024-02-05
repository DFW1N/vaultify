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

# Azure DevOps

This section covers how you create tasks, to automate pipelines with `Vaultify`, it has been broken down into two sections, one being the templated task, what it is actually doing using `bash` and the second part actually using the templated task and passing the required `parameters` to it.

<br>

`install_vaultify.yml`
### Part 1, templated task.

```yml
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

### Part 2, template usage.

```yml
####################
# Install Vaultify #
####################

- template: <relative-path>/vaultify_install.yml
```

---

<br>

## Author

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |
