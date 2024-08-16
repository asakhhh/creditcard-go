package main

import (
	"fmt"
)

var (
	red    = "\u001b[31m"
	blue   = "\u001b[34m"
	cyan   = "\u001b[36m"
	purple = "\x1b[35m"
	reset  = "\u001b[0m"
)

func toBold(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func color(s, col string) string {
	return col + s + reset
}

func valid(s string) bool {
	return s == "validate" || s == "generate" || s == "information" || s == "issue"
}

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
	p := 1
	for i := len(entry) - 1; i >= 0; i-- {
		if entry[i] < '0' || entry[i] > '9' {
			return -1
		}
		modified := int(entry[i]-'0') * p
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

func correctLength(entry string) bool {
	if entry[0] == '4' || (entry[0] == '5' && entry[1] >= '1' && entry[1] <= '5') {
		return len(entry) == 13 || len(entry) == 16
	}
	if entry[:2] == "34" || entry[:2] == "37" {
		return len(entry) == 15
	}
	return len(entry) >= 13 && len(entry) <= 19
}
