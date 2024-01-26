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



---

## Install Vaultify

```bash
go install github.com/DFW1N/vaultify@latest
export PATH=$PATH:$(go env GOPATH)/bin
```

# Vaultify - A CLI Tool for Managing Statefiles

Vaultify is a command-line interface (CLI) tool for managing statefiles, particularly for Terraform projects. It provides functionality to interact with HashiCorp Vault, including secret encryption and pushing/pulling state from a remote Vault server.

## Commands

Vaultify supports the following commands:

| Command       | Description                                       |
|---------------|---------------------------------------------------|
| `init`        | Initialize Vaultify in your operating system.     |
| `validate`    | Validate the JSON format of your Terraform statefile. |
| `update`      | Update the Vaultify CLI to the latest version.    |
| `wrap`        | Wrap a secret in base64 encoding.                |
| `unwrap`      | Unwrap a secret from base64 encoding.            |
| `pull`        | Pull state from a remote HashiCorp Vault server.  |
| `push`        | Push state to a remote HashiCorp Vault server.    |
| `-h, --help`        | Display the help display for Vaultify.    |
| `-v, --version`        | Display Vaultify version installed on the operating system.    |

## Usage

To use Vaultify, simply execute one of the commands listed above. For example, to initialize Vaultify, you can run:

```bash
vaultify <command>
```

---

## Development

### Initialize Go Module

```bash
go mod init vaultify
```

---

## Contributing

This section covers how to contrinute to this project.

---

## Author

Vaultify is maintained by `Sacha Roussakis-Notter`.

## License

This project is licensed under the `GNU General Public License, Version 3 (GPL-3.0)` - see the [LICENSE](LICENSE) file for details.