package cmd

import "fmt"

// Help prints the help information for the program
func Help() {
	fmt.Println("Vaultify - A CLI tool for managing statefiles.")
	fmt.Println("\nCommands:")
	fmt.Println("  -v, -version   Show the Vaultify version")
	fmt.Println("  -h, -help      Show this help message")
	// Add more help entries for other commands here
}
