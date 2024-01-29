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

# Contributing to Vaultify

Thank you for considering contributing to Vaultify! Here are some guidelines to help you get started with development and the release process.

## Development

### Local Testing without Rebuilding

To locally test Vaultify without the need to rebuild it each time, follow these steps:

1. Make sure you are in the working directory where `main.go` exists.

2. Export the necessary environment variable to include Go binaries in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
go run main.go <command>
```

This allows you to test Vaultify commands quickly during development without the need for rebuilding the entire application.

## Pull Requests

If you're an external developer raising a pull request to contribute to Vaultify, please follow these guidelines:

  - Fork the Vaultify repository to your own GitHub account.

  - Clone your forked repository to your local machine.

  - Create a new branch for your changes:

    ```bash
    git checkout -b feature/my-new-feature
    ```

- Make your changes, commit them, and push the branch to your forked repository.

- Open a pull request from your branch to the main Vaultify repository.

- Provide a clear and descriptive title for your pull request, along with details of the changes made.

- Ensure your code adheres to the project's coding standards.

- Respond to any feedback or comments on your pull request promptly.

By following these guidelines, you'll help streamline the process of reviewing and merging your contributions into Vaultify.

**Happy contributing!**

---

## Releasing a New Version (Authorised Individuals)

When you're ready to release a new version of Vaultify, follow these steps:

`Create new tag`
To create a new tag, use the following command, replacing v1.0.9 with the desired version number:

```bash
git tag v1.0.9
```


`Push new tag`
Push the newly created tag to the remote repository using:

```bash
git push origin tag v1.0.9
```

The release workflow will automatically trigger when the tag is pushed.