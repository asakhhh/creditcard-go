package main

import (
	"fmt"
	"math/rand"
	"os"
)

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
			fmt.Printf(color("Error", red) + ": " + [2]string{"no", "invalid"}[min(4, len(os.Args))-3] + " template provided.\n")
			fmt.Printf("Use './creditcard generate --help' for more information.\n")
			os.Exit(1)
		} else {
			if len(os.Args) > 4 {
				fmt.Printf(color("Warning", purple) + ": args after the template are ignored.\n")
				fmt.Printf("Use './creditcard generate --help' for more information.\n====================\n")
			}
			generated := generatedNumbers(os.Args[3])

			if len(generated) == 0 {
				fmt.Printf(color("Error", red) + ": couldn't generate card numbers.\n")
				fmt.Printf("Use './creditcard generate --help' for more information.\n")
				os.Exit(1)
			}

			fmt.Println(generated[rand.Int()%len(generated)])
		}
	} else { // template
		os.Args[2] = removeSpaces(os.Args[2])
		if !isTemplate(os.Args[2]) {
			fmt.Printf(color("Error", red) + ": invalid template provided.\n")
			fmt.Printf("Use './creditcard generate --help' for more information.\n")
			os.Exit(1)
		} else {
			if len(os.Args) > 3 {
				fmt.Printf(color("Warning", purple) + ": args after the template are ignored.\n")
				fmt.Printf("Use './creditcard generate --help' for more information.\n====================\n")
			}
			generated := generatedNumbers(os.Args[2])
			for _, entry := range generated {
				fmt.Println(entry)
			}
		}
	}
}
