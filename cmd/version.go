package cmd

import "fmt"

const version = "1.0.0" // Update the version as needed

// Version prints the version of the program
func Version() {
	fmt.Printf("Vaultify Version: %s\n", version)
}
