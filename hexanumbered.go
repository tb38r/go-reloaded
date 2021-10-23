package tgr

import (
	"strings"
)

func Hexnumbered(s []string) []string {

	pd := 0
	var empty []string

	for i := 0; i < len(s); i++ {

		if s[i] == "(up," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				empty[i-j] = strings.ToUpper(empty[i-j])
			}
			//empty = append(empty, s[i+2:]...)
			//fmt.Printf("\nEMPTY AFTER UP: %v\n", empty)

		}

		if s[i] == "(low," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				empty[i-j] = strings.ToLower(empty[i-j])
			}
			//empty = append(empty, s[i+2:]...)
			//fmt.Printf("\nEMPTY AFTER LOW: %v\n", empty)
		}

		if s[i] == "(cap," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				empty[i-j] = strings.Title(empty[i-j])
			}
			//empty = append(empty, s[i+2:]...)
			//fmt.Printf("\nEMPTY AFTER CAP: %v\n", empty)

		}

		if s[i] != "(up" && s[i] != "(low" && s[i] != "(cap" {
			empty = append(empty, s[i])
			//fmt.Printf("\nEMPTY AFTER NONE: %v\n", empty)

		}

	}
	return empty

}
