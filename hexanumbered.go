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

		}

		if s[i] == "(low," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				empty[i-j] = strings.ToLower(empty[i-j])

			}

		}

		if s[i] == "(cap," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				empty[i-j] = strings.Title(empty[i-j])

			}

		}

		if s[i] != "(up" && s[i] != "(low" && s[i] != "(cap" {
			empty = append(empty, s[i])
		}

	}
	return empty

}
