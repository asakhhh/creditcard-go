package main

import "fmt"

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

func main() {
	s := readline()
	fmt.Println(isValid(s))
}
