package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(color("Error", red) + ": no option chosen.\n")
		fmt.Printf("Use './creditcard --help' for more information.\n")
		os.Exit(1)
	} else if os.Args[1] == "--help" {
		printHelp("general")
	} else if !valid(os.Args[1]) || (len(os.Args) >= 3 && os.Args[2] == "--help") {
		printHelp(os.Args[1])
	} else if len(os.Args) == 2 {
		fmt.Printf(color("Error", red) + ": no args provided.\n")
		fmt.Printf("Use './creditcard " + os.Args[1] + " --help' for more information.\n")
		os.Exit(1)
	} else if os.Args[1] == "validate" { // len >= 3 and valid option and not --help
		validate()
	} else if os.Args[1] == "generate" {
		generate()
	} else if os.Args[1] == "information" {
		information()
	} else {
		issue()
	}
}
