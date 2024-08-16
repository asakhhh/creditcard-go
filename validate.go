package main

import (
	"fmt"
	"os"
)

func validate() {
	var entries []string
	if os.Args[2] == "--stdin" {
		if len(os.Args) > 3 {
			fmt.Printf(color("Warning", purple) + ": the args after --stdin are ignored.\n")
			fmt.Printf("Use './creditcard validate --help' for more information.\n====================\n")
		}

		entries = split(readline(), ' ')
	} else {
		entries = os.Args[2:]
	}
	for _, entry := range entries {
		entry = removeSpaces(entry)
		if isValid(entry) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)
		}
	}
}
