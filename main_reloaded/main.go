package main

import (
	"fmt"
	"os"
	"strings"
	"tgr"
)

func main() {
	arguments := os.Args[1:]

	file_input := arguments[0]
	file := tgr.ReadFile(file_input)
	filesplit := strings.Split(file, " ")  //returns slice of strings
	filehexa := tgr.HexaConvert(filesplit) //basic (cap/bin/hex...) fixed
	filehexanumb := tgr.Hexnumbered(filehexa)
	filepunctone := tgr.Punctone(filehexanumb) // to delete
	filepuncttwo := tgr.Puncttwo(filepunctone)
	fileatoan := tgr.Atoan(filepuncttwo)

	man := os.WriteFile(arguments[1], []byte(fileatoan), 0644)
	if man != nil {
		panic(man)
	}
	// fmt.Printf("\nFINAL OUTPUT : %v", filehexa)
	fmt.Printf("\nFINAL OUTPUT : %v", filehexanumb)
	fmt.Printf("\nFINAL OUTPUT : %v", filepunctone)
	fmt.Printf("\nFINAL OUTPUT : %v", filepuncttwo)
	fmt.Printf("\nFINAL OUTPUT : %v", fileatoan)

}
