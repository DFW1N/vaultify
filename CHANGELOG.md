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


## [v1.0.15]([diff][1.0.14]) - TBD

### Added

- Fixed a bug with install-vault when i introduced, checking that `.vaultify/settings.json` file exists before presenting the status of your connected storage provider, solution.
- Fixed some minor interface errors and endenting.
- Included into install-vault to run `vaultify init`, first before trying to execute `vaultify status`.
- Added TODO tags, inside my code so its easier to remind me what to complete, for the next release.
- Added a confirmation prompt, when running `delete-vault`, command to just to make sure the user is confident they want to delete the container and its contents. 
- Added support to include the `-y` input, to `delete-vault` command so it can be automated in pipelines or any devops solution.
- Renamed directory `sample/` to `samples/`, and added a `terraform/` directory to provide example code to push states easy for testing and validating the CLI.
- Added support, to export your new localhost vault values to your host environment from the docker instance.

---

## [v1.0.14]([diff][1.0.13]) - 03/02/2024

### Added

- Added, more complex settings struct, to support a more complex json object structure.
- Added, support for `vaultify status` to check your secret storage settings.
- Added, a new feature to push your state file to a S3 amazon bucket or azure storage account instead of only having vault.
- Added so when you run `vaultify init` it will build this new data structure.
- Added a new common go function, `CheckAzureEnvVars()` this will check you have set the required environment variables.
- Added new functions, `checkAzureStorageAccountExists()`, `AuthenticateWithAzureAD()` in `common.go` (All azure functions use rest api to do checks)
- Updated, `configure.go` and `init.go`, `status.go`.
- Updated `push.go` to use a case switch statement deppending on your `config.Settings.DefaultSecretStorage` setting.
- Added new functions in `common.go`, `CreateContainer(accountName, key string)`, `generateSignature(accountName, accountKey, method, contentLength, contentType, date, blobType, containerName, blobName string)`, `listStorageAccountKeys()` and `uploadBlobWithAccessKey(accountName, key, encodedStateFilePath string)`.
- The push, command will now identify if you are pushing to vault or azure storage account and will publish either a vault secret or a azure storage blob following same naming convention model.
- Added case switch statement, to the pull command so now it works pulling, azure storage account container blobs, it uses your terraform workspace and base working directory to determine what blob to download from storage account.
- Added function `pullBlobFromAzureStorage(accountName, key string)` and updated, `generateSignature(accountName, accountKey, method, contentLength, contentType, date, blobType, containerName, blobName string)` to work if PUT or GET queries.

---

## [v1.0.13]([diff][1.0.12]) - 2024/02/01

### Added

- Fixed Linting failed on delete-vault command
- Updated the help command
- Added `.goreleaser.yml` file.
- Added my `go-tests` to the release workflow, and added a `docker` push to my release workflow also.

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