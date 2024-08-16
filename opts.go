package main

import (
	"fmt"
	"math/rand"
	"os"
)

func split(inp string) []string {
	var res []string
	var t string
	for _, c := range inp {
		if c == ' ' {
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

func isValid(entry string) bool {
	if len(entry) < 13 || len(entry) > 16 {
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

		entries = split(readline())
	} else {
		entries = os.Args[2:]
	}
	for _, entry := range entries {
		if isValid(entry) {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			os.Exit(1)
		}
	}
}

func isTemplate(entry string) bool {
	if len(entry) < 13 || len(entry) > 16 {
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
			fmt.Println(generated[rand.Int()%len(generated)])
		}
	} else { // template
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

func information() {
}

func issue() {
}
