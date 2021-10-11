package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFromFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
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
