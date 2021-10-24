package tgr

import (
	"fmt"
	"strings"
)

func Punctone(n []string) string {

	var empty []rune
	sjoined := strings.Join(n, " ")

	fmt.Println("\nThis is sjoined:", sjoined)

	nrune := []rune(sjoined)

	fmt.Println("\n This is nrune", nrune)

	for i := 0; i < len(nrune)-1; i++ {

		if IsPunc(nrune[i]) && nrune[i-1] == 32 {
			nrune[i], nrune[i-1] = nrune[i-1], nrune[i]

		}
		empty = append(empty, nrune[i])
	}
	return string(empty)

}
