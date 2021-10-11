package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFromFile() {
	args := os.Args[1:]
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file) // Can use bufio.Scanner to implement "print each line" function
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	readFromFile()
}
