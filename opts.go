package main

import (
	// "flag"
	"fmt"
	"math/rand"
	"os"
)

func split(inp string, sep rune) []string {
	var res []string
	var t string
	for _, c := range inp {
		if c == sep {
			if len(t) > 0 {
				res = append(res, t)
			}
			t = ""
		} else {
			t += string(c)
		}
	}
	if len(t) > 0 {
		res = append(res, t)
	}
	return res
}

func checkSum(entry string) int {
	sm := 0
	p := 2
	for _, c := range entry {
		if c < '0' || c > '9' {
			return -1
		}
		modified := int(c-'0') * p
		if modified >= 10 {
			modified = modified/10 + modified%10
		}
		sm += modified
		p = 3 - p // 1 and 2 factor roggle
	}
	return sm
}

func removeSpaces(entry string) string {
	var res string
	for _, c := range entry {
		if c != ' ' {
			res += string(c)
		}
	}
	return res
}

func isValid(entry string) bool {
	if len(entry) < 13 || len(entry) > 19 {
		return false
	}
	for _, c := range entry {
		if c < '0' || c > '9' {
			return false
		}
	}

	return checkSum(entry)%10 == 0
}

func readline() string {
	var res string
	var c rune
	fmt.Scanf("%c", &c)
	for c != '\n' {
		res += string(c)
		fmt.Scanf("%c", &c)
	}
	return res
}

func validate() {
	var entries []string
	if os.Args[2] == "--stdin" {
		if len(os.Args) > 3 {
			fmt.Printf(colorPurple("Warning") + ": the args after --stdin are ignored.\n")
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

func isTemplate(entry string) bool {
	if len(entry) < 13 || len(entry) > 19 {
		return false
	}
	for _, c := range entry {
		if c != '*' && !(c >= '0' && c <= '9') {
			return false
		}
	}

	asterisks := 0
	for i := len(entry) - 1; i >= 0; i-- {
		if entry[i] == '*' {
			asterisks++
		} else {
			break
		}
	}
	if asterisks > 4 || asterisks == 0 {
		return false
	}

	for i := len(entry) - 1 - asterisks; i >= 0; i-- {
		if entry[i] == '*' {
			return false
		}
	}

	return true
}

func generatedNumbers(template string) []string {
	if template[len(template)-2] != '*' && template[len(template)-1] == '*' {
		last := rune('0' + (10-checkSum(template[:len(template)-1])%10)%10)
		return []string{template[:len(template)-1] + string(last)}
	}
	var res []string
	ind := len(template) - 1
	for ind > 0 && template[ind-1] == '*' {
		ind--
	}

	for _, c := range "0123456789" {
		newTemplate := template[:ind] + string(c) + template[ind+1:]
		res = append(res, generatedNumbers(newTemplate)...)
	}

	return res
}

func generate() {
	if os.Args[2] == "--pick" {
		if len(os.Args) > 3 {
			os.Args[3] = removeSpaces(os.Args[3])
		}
		if len(os.Args) == 3 || !isTemplate(os.Args[3]) {
			fmt.Printf(colorRed("Error") + ": " + [2]string{"no", "invalid"}[min(4, len(os.Args))-3] + " template provided.\n")
			fmt.Printf("Use './creditcard generate --help' for more information.\n")
			os.Exit(1)
		} else {
			if len(os.Args) > 4 {
				fmt.Printf(colorPurple("Warning") + ": args after the template are ignored.\n")
				fmt.Printf("Use './creditcard generate --help' for more information.\n====================\n")
			}
			generated := generatedNumbers(os.Args[3])

			if len(generated) == 0 {
				fmt.Printf(colorRed("Error") + ": couldn't generate card numbers.\n")
				fmt.Printf("Use './creditcard generate --help' for more information.\n")
				os.Exit(1)
			}

			fmt.Println(generated[rand.Int()%len(generated)])
		}
	} else { // template
		os.Args[2] = removeSpaces(os.Args[2])
		if !isTemplate(os.Args[2]) {
			fmt.Printf(colorRed("Error") + ": invalid template provided.\n")
			fmt.Printf("Use './creditcard generate --help' for more information.\n")
			os.Exit(1)
		} else {
			if len(os.Args) > 3 {
				fmt.Printf(colorPurple("Warning") + ": args after the template are ignored.\n")
				fmt.Printf("Use './creditcard generate --help' for more information.\n====================\n")
			}
			generated := generatedNumbers(os.Args[2])
			for _, entry := range generated {
				fmt.Println(entry)
			}
		}
	}
}

func correctLength(entry string) bool {
	if entry[0] == '4' || (entry[0] == '5' && entry[1] >= '1' && entry[1] <= '5') {
		return len(entry) == 13 || len(entry) == 16
	}
	if entry[:2] == "34" || entry[:2] == "37" {
		return len(entry) == 15
	}
	return len(entry) >= 13 && len(entry) <= 19
}

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
				fmt.Printf(colorRed("Error") + ": brands file is specified incorrectly.\n")
				fmt.Printf("Use './creditcard information --help' for more information.\n")
				os.Exit(1)
			}
		} else if len(arg) >= 9 && arg[:9] == "--issuers" {
			if len(arg) >= 11 && arg[9] == '=' && issuersIndex == -1 {
				issuersIndex = ind
				issuersFile = arg[10:]
			} else {
				fmt.Printf(colorRed("Error") + ": issuers file is specified incorrectly.\n")
				fmt.Printf("Use './creditcard information --help' for more information.\n")
				os.Exit(1)
			}
		} else if arg == "--stdin" {
			stdinIndex = ind
		}
	}
	if brandsIndex == -1 {
		fmt.Printf(colorRed("Error") + ": brands file is not specified.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	} else if issuersIndex == -1 {
		fmt.Printf(colorRed("Error") + ": issuers file is not specified.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	} else if brandsIndex >= 3+min(stdinIndex, 0) || issuersIndex >= 3+min(stdinIndex, 0) {
		fmt.Printf(colorRed("Error") + ": flags should come before card numbers.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	} else if stdinIndex != -1 && stdinIndex >= 3 {
		fmt.Printf(colorRed("Error") + ": flags should come before card numbers.\n")
		fmt.Printf("Use './creditcard information --help' for more information.\n")
		os.Exit(1)
	}

	brands, err1 := os.ReadFile(brandsFile)
	issuers, err2 := os.ReadFile(issuersFile)
	if err1 != nil {
		fmt.Printf(colorRed("Error") + ": \"" + brandsFile + "\" file does not exist.\n")
	}
	if err2 != nil {
		fmt.Printf(colorRed("Error") + ": \"" + issuersFile + "\" file does not exist.\n")
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
			fmt.Printf(colorPurple("Warning") + ": the args after flags are ignored.\n")
			fmt.Printf("Use './creditcard information --help' for more information.\n====================\n")
		}
		entries = split(readline(), ' ')
		fmt.Printf("====================\n")
	}

	if len(entries) == 0 {
		fmt.Printf(colorRed("Error") + ": no card numbers provided.\n")
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

func issue() {
	var brandsFile, issuersFile, brandName, issuerName string

	if len(os.Args) < 6 {
		fmt.Printf(colorRed("Error") + ": not enough arguments.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	} else if len(os.Args) > 6 {
		fmt.Printf(colorRed("Error") + ": too much arguments.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	}

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
		fmt.Printf(colorRed("Error") + ": brands file is not specified.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	} else if len(issuersFile) == 0 {
		fmt.Printf(colorRed("Error") + ": issuers file is not specified.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	} else if len(brandName) == 0 {
		fmt.Printf(colorRed("Error") + ": brand is not specified.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	} else if len(issuerName) == 0 {
		fmt.Printf(colorRed("Error") + ": issuer is not specified.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	}

	brands, err1 := os.ReadFile(brandsFile)
	issuers, err2 := os.ReadFile(issuersFile)
	if err1 != nil {
		fmt.Printf(colorRed("Error") + ": \"" + brandsFile + "\" file does not exist.\n")
	}
	if err2 != nil {
		fmt.Printf(colorRed("Error") + ": \"" + issuersFile + "\" file does not exist.\n")
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
		fmt.Printf(colorRed("Error") + ": brand not found in the file.\n")
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
		fmt.Printf(colorRed("Error") + ": issuer not found in the file.\n")
		fmt.Printf("Use './creditcard issue --help' for more information.\n")
		os.Exit(1)
	}

	if issuerPrefix[:len(brandPrefix)] != brandPrefix {
		fmt.Printf(colorRed("Error") + ": brand and issuer prefixes do not match - impossible to issue a card.\n")
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
	last := (10 - checkSum(issuerPrefix)%10) % 10
	issuerPrefix += string(rune('0' + last))

	fmt.Println(issuerPrefix)
}
