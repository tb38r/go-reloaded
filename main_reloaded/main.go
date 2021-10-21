package main

import (
	"fmt"
	"os"
	"strings"
	"tgr"
)

func main() {
	arguments := os.Args

	file_input := arguments[1]
	// := arguments[2]

	file := tgr.ReadFile(file_input)

	filesplit := strings.Split(file, " ") //returns slice of strings
	fmt.Println(filesplit)
	fmt.Println((tgr.HexaConvert(filesplit)))

}
