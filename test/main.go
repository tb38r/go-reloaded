package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func hexaconvert(n string) string {

	num, err := strconv.ParseInt(n, 16, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(num)

}
func binaryconvert(n string) string {

	num, err := strconv.ParseInt(n, 2, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(num)

}

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

	t2 := strings.Join(text, "")
	t3 := strings.Split(t2, " ")

	// fmt.Printf("%#v %T\n", t3, t3)

	for i := 0; i < len(t3); i++ {
		if t3[i] == "(hex)" && i > 0 {
			t3[i-1] = hexaconvert(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)

		} else if t3[i] == "(bin)" && i > 0 {
			t3[i-1] = binaryconvert(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)

		} else if t3[i] == "(up)" && i > 0 {
			t3[i-1] = strings.ToUpper(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)

		} else if t3[i] == "(low)" && i > 0 {
			t3[i-1] = strings.ToLower(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)

		} else if t3[i] == "(cap)" && i > 0 {
			t3[i-1] = strings.Title(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)
		}

	}
	fmt.Println(t3)
}
