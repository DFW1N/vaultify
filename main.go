package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version = "1.0.0" // You can update the version as needed
)

func main() {
	versionFlag := flag.Bool("v", false, "Prints the version of the program")
	helpFlag := flag.Bool("h", false, "Prints the help information")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("Vaultify Version: %s\n", version)
		os.Exit(0)
	}

	if *helpFlag {
		fmt.Println("Vaultify - A tool to push OpenTofu or Terraform statefiles to HashiCorp Vault and encrypt them with Base64.")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Add your main program logic here
}
