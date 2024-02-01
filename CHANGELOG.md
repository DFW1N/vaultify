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

```bash
██╗   ██╗ █████╗ ██╗   ██╗██╗  ████████╗██╗███████╗██╗   ██╗
██║   ██║██╔══██╗██║   ██║██║  ╚══██╔══╝██║██╔════╝╚██╗ ██╔╝
██║   ██║███████║██║   ██║██║     ██║   ██║█████╗   ╚████╔╝ 
╚██╗ ██╔╝██╔══██║██║   ██║██║     ██║   ██║██╔══╝    ╚██╔╝  
 ╚████╔╝ ██║  ██║╚██████╔╝███████╗██║   ██║██║        ██║   
  ╚═══╝  ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝   ╚═╝╚═╝        ╚═╝   
                                                            
```

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).


## [v1.0.14]([diff][1.0.12]) - TBA

### Added


---

## [v1.0.13]([diff][1.0.12]) - HOTFIX [2024/02/01]

### Added

- Fixed Linting failed on delete-vault command

---

## [v1.0.12]([diff][1.0.11]) - 2024/02/01

### Added

- Added a `sample/`directory
- Added fake dummy data into `sample/` directory so a valid `.json` file can be used for testing purposes to publish and pull from `Hashicorp Vault`.
- Added a `README.md` file to `sample/` to explain this is fake data and does not represent anything.
- Updated example useage screenshot
- Updated all documentation to support the new command added `permissions`.
- Added `vaultify permissions` to validate if you have relevant roles or policies on your token to create secrets on your selected `engineType` in `Hashicorp Vault`
- Added a `policies/` directory which has a `README.md` and a `kv_secrets_policy.hcl` file as examples of how to publish policies for granting permissions to your selected `engineType`.
- Added `docker/` directory with a dockerFile and `docker-compose.yml` file with a `README.md` file.
- Added `scripts/` to the `docker/` directory that contains multiple scripts.
- Deleted so the environment variables dont need to be set to run `-h` or `-v` commands for `vaultify`
- Updated `README.md` file for docker support if you want to deploy `Hashicorp Vault` locally.
- Updated `docker-compose.yml` and `initialize-vault.sh`.
- Added `2` new `Vaultify` commands `install-vault` and `delete-vault`.
- Added a new .go file to have `common` functions that is used across the whole `CLI`
- Created a Docker Hub and published the newest `vaultify-vault` [containers](https://hub.docker.com/r/buungroup/vault-raft) 
- Updated the `install.docker.sh` script, to support multiple OS types.
- Updated all the `README.md`, including `CLI.md` and created new CLI docs for the commands.

---

## [v1.0.11]([diff][1.0.10]) - 2024-01-31

### Added

- Added a new command called `vaultify publish` it comes wrap and push to vault.
- Added `Go!` tests
- Added `CODEOWNERS` file
- Added pull request template file in `/workflows`
- Added `ISSUE_TEMPLATE` directory under` /workflows` & issue report template
- Changed the `README.md` to contain less information and move everything to the vaultify website at `https://vaultify.buungroup.com`
- Added a debian folder to future release apt-get package as part of the release workflow.
- Updated command list in` README.md` and also added links to documentation.
- Added `color` to alot of the outputs to make it to visually read the outputs and commands.
- Updated `configure` command
- Added so `init` command creates a `.vaultify` directory with a `settings.json` file this ensures, that your engine name can be changed and added additional functionality for future releases if you want to use terraform workspace name as part of the secret being published.
- Moved functions that where shared across multiple commands to a single `.go` file this way its easier to update and manage those functions without looking within commands.
- Improved the `delete` command to actually go to the latest version and validate that the json content data has been deleted, it was previously failing because the secret path was getting checked for existence but the secret data was deleted if multiple versions, this has been improved so it will go to latest version and try find the actual data of the secret and if it doesn't find any it outputs the latest secret has already been deleted.
- Made other modules check if `vaulift init` has been executed to ensure it creates the `settings.json` file created, only commands that need the data in it has been adjusted.

---

## [v1.0.10]([diff][1.0.9]) - 2024-01-29

### Added

- Added CHANGELOG.md file for future logging of changelogs.

---