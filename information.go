package main

import (
	"fmt"
	"os"
)

func information() {
	var brandsFile, issuersFile string
	brandsIndex, issuersIndex := -1, -1
	stdinIndex := -1
	for ind, arg := range os.Args[2:] {
		if len(arg) >= 8 && arg[:8] == "--brands" {
			if len(arg) >= 10 && arg[8] == '=' && brandsIndex == -1 {
				brandsIndex = ind
				brandsFile = arg[9:]
			} else {
				fmt.Printf(color("Error", red) + ": brands file is specified incorrectly.\n")
				fmt.Printf("Use './creditcard information --help' for more information.\n")
				os.Exit(1)
			}
		} else if len(arg) >= 9 && arg[:9] == "--issuers" {
			if len(arg) >= 11 && arg[9] == '=' && issuersIndex == -1 {
				issuersIndex = ind
				issuersFile = arg[10:]
			} else {
				fmt.Printf(color("Error", red) + ": issuers file is specified incorrectly.\n")
				fmt.Printf("Use './creditcard information --help' for more information.\n")
				os.Exit(1)
			}
		} else if arg == "--stdin" {
			stdinIndex = ind
		}
	}
	if brandsIndex == -1 {
		fmt.Printf(color("Error", red) + ": brands file is not specified.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	} else if issuersIndex == -1 {
		fmt.Printf(color("Error", red) + ": issuers file is not specified.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	} else if brandsIndex >= 3+min(stdinIndex, 0) || issuersIndex >= 3+min(stdinIndex, 0) {
		fmt.Printf(color("Error", red) + ": flags should come before card numbers.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	} else if stdinIndex != -1 && stdinIndex >= 3 {
		fmt.Printf(color("Error", red) + ": flags should come before card numbers.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	}

	brands, err1 := os.ReadFile(brandsFile)
	issuers, err2 := os.ReadFile(issuersFile)
	if err1 != nil {
		fmt.Printf(color("Error", red) + ": \"" + brandsFile + "\" file does not exist.\n")
	}
	if err2 != nil {
		fmt.Printf(color("Error", red) + ": \"" + issuersFile + "\" file does not exist.\n")
	}
	if err1 != nil || err2 != nil {
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	}

	var entries []string
	if stdinIndex == -1 {
		entries = os.Args[4:]
	} else {
		if len(os.Args) > 5 {
			fmt.Printf(color("Warning", purple) + ": the args after flags are ignored.\n")
			fmt.Printf("Use './creditcard information --help' for more information.\n====================\n")
		}
		entries = split(readline(), ' ')
		fmt.Printf("====================\n")
	}

	if len(entries) == 0 {
		fmt.Printf(color("Error", red) + ": no card numbers provided.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	}

	for _, entry := range entries {
		entry = removeSpaces(entry)
		fmt.Println(entry)
		fmt.Printf("Correct: ")
		if !isValid(entry) || !correctLength(entry) {
			fmt.Printf("no\nCard Brand: -\nCard Issuer: -\n\n")
		} else {
			fmt.Printf("yes\nCard Brand: ")
			found := false
			for _, brand := range split(string(brands), '\n') {
				prefix := split(brand, ':')[1]
				if len(entry) >= len(prefix) && entry[:len(prefix)] == prefix {
					found = true
					fmt.Printf(split(brand, ':')[0] + "\nCard Issuer: ")
					break
				}
			}
			if !found {
				fmt.Printf("-\nCard Issuer: -\n\n")
			} else {
				found = false
				for _, issuer := range split(string(issuers), '\n') {
					prefix := split(issuer, ':')[1]
					if len(entry) >= len(prefix) && entry[:len(prefix)] == prefix {
						found = true
						fmt.Printf(split(issuer, ':')[0] + "\n\n")
						break
					}
				}
				if !found {
					fmt.Printf("-\n")
				}
			}
		}
	}
}
