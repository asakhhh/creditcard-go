package main

import (
	"fmt"
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

func isValid(entry string) bool {
	if len(entry) < 13 || len(entry) > 16 {
		return false
	}

	sm := 0
	p := 2
	for _, c := range entry {
		if c < '0' || c > '9' {
			return false
		}
		modified := int(c-'0') * p
		if modified >= 10 {
			modified = modified/10 + modified%10
		}
		sm += modified
		p = 3 - p // 1 and 2 factor roggle
	}
	return sm%10 == 0
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
			fmt.Printf(colorRed("Warning") + ": the args after --stdin are ignored.\n")
			fmt.Printf("Use './creditcard validate --help' for more information.\n\n")
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

func generate() {
}

func information() {
}

func issue() {
}
