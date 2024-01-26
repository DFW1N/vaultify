package main

import (
	"flag"
	"fmt"
	"vaultify/cmd"
)

func main() {
	// Define flags for version and help
	versionFlag := flag.Bool("v", false, "Prints the version of the program")
	versionFlagLong := flag.Bool("version", false, "Prints the version of the program")
	helpFlag := flag.Bool("h", false, "Prints the help information")
	helpFlagLong := flag.Bool("help", false, "Prints the help information")

	flag.Parse()

	switch {
	case *versionFlag || *versionFlagLong:
		cmd.Version()
	case *helpFlag || *helpFlagLong:
		cmd.Help()
	default:
		fmt.Println("Usage: vaultify [command]")
		fmt.Println("Use 'vaultify -h' for help.")
	}
}
