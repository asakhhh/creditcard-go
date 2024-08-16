package main

import (
	"fmt"
	"math/rand"
	"os"
)

func issue() {
	if len(os.Args) < 6 {
		fmt.Printf(color("Error", red) + ": not enough arguments.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	} else if len(os.Args) > 6 {
		fmt.Printf(color("Error", red) + ": too much arguments.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	}

	var brandsFile, issuersFile, brandName, issuerName string
	for _, arg := range os.Args[2:] {
		if len(arg) >= 11 && arg[:10] == "--issuers=" {
			issuersFile = arg[10:]
		} else if len(arg) >= 10 && arg[:9] == "--brands=" {
			brandsFile = arg[9:]
		} else if len(arg) >= 9 && arg[:8] == "--brand=" {
			brandName = arg[8:]
		} else if len(arg) >= 10 && arg[:9] == "--issuer=" {
			issuerName = arg[9:]
		}
	}

	if len(brandsFile) == 0 {
		fmt.Printf(color("Error", red) + ": brands file is not specified.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	} else if len(issuersFile) == 0 {
		fmt.Printf(color("Error", red) + ": issuers file is not specified.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	} else if len(brandName) == 0 {
		fmt.Printf(color("Error", red) + ": brand is not specified.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	} else if len(issuerName) == 0 {
		fmt.Printf(color("Error", red) + ": issuer is not specified.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
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
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	}

	var brandPrefix, issuerPrefix string
	for _, brand := range split(string(brands), '\n') {
		if brandName == split(brand, ':')[0] {
			brandPrefix = split(brand, ':')[1]
			break
		}
	}
	if len(brandPrefix) == 0 {
		fmt.Printf(color("Error", red) + ": brand not found in the file.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	}

	for _, issuer := range split(string(issuers), '\n') {
		if issuerName == split(issuer, ':')[0] {
			issuerPrefix = split(issuer, ':')[1]
			break
		}
	}
	if len(issuerPrefix) == 0 {
		fmt.Printf(color("Error", red) + ": issuer not found in the file.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	}

	if issuerPrefix[:len(brandPrefix)] != brandPrefix {
		fmt.Printf(color("Error", red) + ": brand and issuer prefixes do not match - impossible to issue a card.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	}

	length := 16
	if brandPrefix == "34" || brandPrefix == "37" {
		length--
	}

	for len(issuerPrefix) < length-1 {
		issuerPrefix += string(rune('0' + rand.Int()%10))
	}
	last := (10 - checkSum(issuerPrefix+"0")%10) % 10
	issuerPrefix += string(rune('0' + last))

	fmt.Println(issuerPrefix)
}
