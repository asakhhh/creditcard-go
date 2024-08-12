package main

import (
	"os"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		printHelp("general")
	} else if !valid(os.Args[1]) || len(os.Args) == 2 || os.Args[2] == "--help" {
		printHelp(os.Args[1])
	} else if os.Args[1] == "validate" { // len >= 3 and valid option and not --help
	} else if os.Args[1] == "generate" {
	} else if os.Args[1] == "information" {
	} else {
	}
}
