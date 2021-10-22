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
	filesplit := strings.Split(file, " ")    //returns slice of strings
	filecapped := tgr.HexaConvert(filesplit) //basic (cap/bin/hex...) fixed


	

}
