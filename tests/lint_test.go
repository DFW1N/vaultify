// ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														  #
// # Project: Vaultify																	  #
// # Description: Easily push, pull and encrypt tofu and terraform statefiles from Vault. #
// ########################################################################################

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func ensureGolangciLintInstalled(t *testing.T) {
	// Install golangci-lint if not already installed
	installCmd := exec.Command("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2")
	if _, err := installCmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to install golangci-lint: %v", err)
	}

	// Determine GOPATH
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(os.Getenv("HOME"), "go")
	}

	// Set PATH to include GOPATH/bin
	os.Setenv("PATH", os.Getenv("PATH")+string(os.PathListSeparator)+filepath.Join(gopath, "bin"))
}

func TestLinting(t *testing.T) {
	ensureGolangciLintInstalled(t)

	// Run golangci-lint
	cmd := exec.Command("golangci-lint", "run")
	cmd.Dir = "../"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Linting failed: %v\n%s", err, output)
	}
}
