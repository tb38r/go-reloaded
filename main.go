package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("test.txt") //os.Open opens file in read only mode
	if err != nil {
		log.Fatal("Failed to open")
	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.

	Scanner := bufio.NewScanner(file) //creating scanner, scans input from chosen source (file, as declared in lines 11/12)
	// The bufio.ScanLines is used as an input to the method bufio.Scanner.Split()

	Scanner.Split(bufio.ScanLines) // and then the scanning forwards to each new line using the bufio.Scanner.Scan() method

	var text []string

	for Scanner.Scan() {
		text = append(text, Scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	t2 := strings.Join(text, " ")
	t3 := strings.Split(t2, " ")

	fmt.Printf("%v, %T", t3, t3)

}
