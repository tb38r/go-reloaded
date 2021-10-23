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
	file := tgr.ReadFile(file_input)
	filesplit := strings.Split(file, " ")  //returns slice of strings
	filehexa := tgr.HexaConvert(filesplit) //basic (cap/bin/hex...) fixed
	filehexanumb := tgr.Hexnumbered(filehexa)
	_ = filehexanumb // to delete

	fmt.Println("\n", filehexanumb)

}
