package main

import (
	"fmt"
	"os"
)

func helpValidate() {
	fmt.Printf(color("validate", blue) + " - option for checking the validity of one or more credit cards. Evaluation is based on the Luhn algorithm.\n" + color("Note", cyan) + ": Card numbers should be entered in one line separated by a space.\n")
	fmt.Println(toBold("usages:"))
	fmt.Printf("./creditcard validate [number]...\t- prints OK or INCORRECT for entered card number(s).\n")
	fmt.Printf("./creditcard validate --stdin\t\t- card numbers are read from standard input.\n")
	fmt.Printf("./creditcard validate --help\t\t- prints help message for this option.\n")
}

func helpGenerate() {
	fmt.Printf(color("generate", blue) + " - option for generating valid credit card numbers based on a given template. Prints all combinations by substituting digits for all asterisks.\n" + color("Note", cyan) + ": Template must have at most 4 asterisks, all at the end.\n")
	fmt.Println(toBold("usages:"))
	fmt.Printf("./creditcard generate <template>\t - outputs all possible combinations that satisfy the given template.\n")
	fmt.Printf("./creditcard generate --pick <template>\t - randomly chooses one number that satisfies the given template.\n")
	fmt.Printf("./creditcard generate --help\t\t - prints help message for this option.\n")
}

func helpInformation() {
	fmt.Printf(color("information", blue) + " - option for getting info about the card number's validity, brand and issuer.\n" + color("Note", cyan) + ": Both mandatory flags should come before the card numbers.\n")
	fmt.Println(toBold("usages:"))
	fmt.Printf("./creditcard information [flags] [numbers]...\n")
	fmt.Printf("./creditcard information [flags] --stdin\n")
	fmt.Printf("./creditcard information --help\n\n")
	fmt.Println(toBold("Flags:"))
	fmt.Printf(toBold("--brands=FILE") + "\t- MANDATORY flag that provides the file containing the list of brands.\n")
	fmt.Printf(toBold("--issuers=FILE") + "\t- MANDATORY flag that provides the file containing the list of issuers.\n")
	fmt.Printf(toBold("--stdin") + "\t\t- card numbers are read from standard input.\n")
	fmt.Printf(toBold("--help") + "\t\t- prints help message for this option.\n")
}

func helpIssue() {
	fmt.Printf(color("issue", blue) + " - option for generating a random card number for a given brand and issuer.\n" + color("Note", cyan) + ": Exactly four mandatory flags should be specified without any other args.\n")
	fmt.Println(toBold("usages:"))
	fmt.Printf("./creditcard issue [flags]\n")
	fmt.Printf("./creditcard issue --help\n\n")
	fmt.Println(toBold("Flags:"))
	fmt.Printf(toBold("--brands=FILE") + "\t- MANDATORY flag that provides the file containing the list of brands.\n")
	fmt.Printf(toBold("--issuers=FILE") + "\t- MANDATORY flag that provides the file containing the list of issuers.\n")
	fmt.Printf(toBold("--brand=") + "\t- MANDATORY flag that specifies the brand of the card.\n")
	fmt.Printf(toBold("--issuer=") + "\t- MANDATORY flag that specifies the issuer of the card.\n")
	fmt.Printf(toBold("--help") + "\t\t- prints help message for this option.\n")
}

func printHelp(opt string) {
	if opt == "general" {
		fmt.Printf(color(toBold("creditcard"), cyan) + " is a tool for handling credit card number(s).\n")
		fmt.Println(toBold("usages:"))
		fmt.Printf("./creditcard option [flags] [args]\n")
		fmt.Printf("./creditcard --help\t\t- outputs help message for the program.\n")
		fmt.Printf("./creditcard option --help\t- outputs help message for the option.\n\n")
		fmt.Printf(toBold("Options:\n"))
		helpValidate()
		fmt.Printf("\n\n")
		helpGenerate()
		fmt.Printf("\n\n")
		helpInformation()
		fmt.Printf("\n\n")
		helpIssue()
		fmt.Println()
	} else if opt == "validate" {
		helpValidate()
	} else if opt == "generate" {
		helpGenerate()
	} else if opt == "information" {
		helpInformation()
	} else if opt == "issue" {
		helpIssue()
	} else {
		fmt.Printf(color("Error", red) + ": incorrect option chosen - " + color(os.Args[1], red) + "\n")
		fmt.Printf("Use './creditcard --help' for more information.\n")
		os.Exit(1)
	}
}
