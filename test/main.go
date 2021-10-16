package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPunc(category string) bool {
	switch category {
	case
		".",
		",",
		"!",
		":",
		";",
		"?":
		return true
	}
	return false
}

func removebrackets(n string) string {

	var opening int
	var closing int

	for index, char := range n {
		if char == '(' {
			opening = index
		} else if char == ')' {
			closing = index + 1
		}
	}

	toAdd := closing - opening

	var empty string

	for i := 0; i < len(n); i++ {
		if i == opening {
			i += toAdd
		} else {
			empty += string(n[i])
			//empty = append(empty, string(n[i]))
		}
	}
	return empty
}

func reverse(n []string) []string {

	reversed := []string{}

	for _, v := range n {
		reversed = append([]string{v}, reversed...)
		// result = string(v) + result
	}
	return reversed //fmt.Sprint(reserved)
}

func isvowelh(n string) bool {

	var result bool

	for i := 0; i < len(n); i++ {
		if n[0] == 'A' || n[0] == 'E' || n[0] == 'I' || n[0] == 'O' || n[0] == 'U' || n[0] == 'a' || n[0] == 'e' ||
			n[0] == 'i' || n[0] == 'o' || n[0] == 'u' || n[0] == 'h' {
			result = true

		} else {
			result = false
		}
	}

	return result

}

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

	for i := 0; i < len(t3); i++ {
		if t3[i] == "(hex)" && i > 0 {
			t3[i-1] = hexaconvert(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)
			t4 := strings.Join(t3, " ")
			_ = t4 //omit

		} else if t3[i] == "(bin)" && i > 0 {
			t3[i-1] = binaryconvert(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)
			t4 := strings.Join(t3, " ")
			_ = t4 // omit

		} else if t3[i] == "(up)" && i > 0 {
			t3[i-1] = strings.ToUpper(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)
			t4 := strings.Join(t3, " ")
			_ = t4 // omit

		} else if t3[i] == "(low)" && i > 0 {
			t3[i-1] = strings.ToLower(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)
			t4 := strings.Join(t3, " ")
			_ = t4 //omit

		} else if t3[i] == "(cap)" && i > 0 {
			t3[i-1] = strings.Title(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)
			t4 := strings.Join(t3, " ")
			_ = t4 //omit

		} else if (t3[i] == "a") && (isvowelh(t3[i+1]) == true) {
			t3[i] = "an"
			t4 := strings.Join(t3, " ")
			_ = t4 //omit

		} else if (t3[i] == "A") && (isvowelh(t3[i+1]) == true) {
			t3[i] = "An"
			t4 := strings.Join(t3, " ")
			_ = t4 //omit

		}
	}
	reversed := reverse(t3)
	for i := 0; i < len(reversed); i++ {
		if (reversed[i][0] > '0' && reversed[i][0] <= '9') && reversed[i][len([]rune(reversed[i]))-1] == ')' {
			n, _ := strconv.Atoi(reversed[i][:1]) //[:1] reads as string , [0] reads as a solitary byte, not wanted
			if strings.Contains(reversed[i+1], "(up") {
				for j := 1; j <= n; j++ {
					reversed[(i+1)+j] = strings.ToUpper(reversed[(i+1)+j])
				}
				reform := reverse(reversed)
				rejoined := strings.Join(reform, " ")
				final := removebrackets(rejoined)
				fmt.Println(final)

			} else if strings.Contains(reversed[i+1], "(low") {
				for j := 1; j <= n; j++ {
					reversed[(i+1)+j] = strings.ToLower(reversed[(i+1)+j])
				}
				reform := reverse(reversed)
				rejoined := strings.Join(reform, " ")
				final := removebrackets(rejoined)
				fmt.Println(final)
			} else if strings.Contains(reversed[i+1], "(cap") {
				for j := 1; j <= n; j++ {
					reversed[(i+1)+j] = strings.Title(reversed[(i+1)+j])
				}
				reform := reverse(reversed)
				rejoined := strings.Join(reform, " ")
				final := removebrackets(rejoined)
				_ = final //to omit
			}
		}
	}
	//PUNCTUATION #1
	srune := []rune(t2)

	for i := 0; i < len(srune); i++ {
		if isPunc(string(srune[i])) && srune[i-1] == ' ' {
			srune[i], srune[i-1] = srune[i-1], srune[i]
		} else if srune[len(srune)-1] == 39 && srune[len(srune)-2] == ' ' {
			srune[len(srune)-2], srune[len(srune)-1] = srune[len(srune)-1], srune[len(srune)-2]
		} else if srune[i] == 39 && srune[i-1] == ' ' && srune[i+1] == ' ' {
			srune[i], srune[i+1] = srune[i+1], srune[i]
		}

	}
	t2 = string(srune)
	fmt.Printf("%#v\n", t2)
	t2unspaced := strings.ReplaceAll(t2, "  ", " ")
	final := strings.TrimSpace(t2unspaced)
	fmt.Printf("%#v", final)

}
