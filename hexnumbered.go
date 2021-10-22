package tgr

import (
	"fmt"
	"strings"
)

func TrimAtoi(s string) int {
	neg := false       // intialise neg as false
	slice := []rune(s) // slice string into a rune to manipulate it
	trim := 0          // initialise trim as 0
	for i := 0; i < len(slice); i++ {

		if !neg && trim == 0 && slice[i] == '-' {
			neg = true
		}
		if slice[i] >= '0' && slice[i] <= '9' {
			trim *= 10
			trim += int(slice[i] - 48)
		}
	}
	if neg {
		return trim * -1
	}
	return trim
}

func HexNumbered(s []string) []string {

	pd := 0
	var empty []string

	for i := 0; i < len(s); i++ {

		if s[i] == "(up," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				empty[i-j] = strings.ToUpper(empty[i-j])
				fmt.Println("\n", s[i-j])

			}

		}

		if s[i] == "(low," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				empty[i-j] = strings.ToLower(empty[i-j])
				fmt.Println("\n", s[i-j])

			}

		}

		if s[i] == "(cap," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				empty[i-j] = strings.Title(empty[i-j])
				fmt.Println("\n", s[i-j])

			}

		}

		if s[i] != "(up" && s[i] != "(low" && s[i] != "(cap" {
			empty = append(empty, s[i])
		}

	}
	return empty

}
