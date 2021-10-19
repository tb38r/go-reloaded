package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"tgr"
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

	t2 := strings.Join(text, "")
	t3 := strings.Split(t2, " ")

	for i := 0; i < len(t3); i++ {
		if t3[i] == "(hex)" && i > 0 {
			t3[i-1] = tgr.HexaConvert(t3[i-1])
			t3 = append(t3[:i], t3[i+1:]...)
			t4 := strings.Join(t3, " ")
			_ = t4 //omit

		} else if t3[i] == "(bin)" && i > 0 {
			t3[i-1] = tgr.BinaryConvert(t3[i-1])
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
			fmt.Printf("%#v\n", t4)

		} else if (t3[i] == "a") && (tgr.IsVowelh(t3[i+1]) == true) {
			t3[i] = "an"
			t4 := strings.Join(t3, " ")
			_ = t4 //omit

		} else if (t3[i] == "A") && (tgr.IsVowelh(t3[i+1]) == true) {
			t3[i] = "An"
			t4 := strings.Join(t3, " ")
			_ = t4 //omit
			fmt.Printf("%#v", t4)

		}
	}
	// (LOW,UP,CAP [NUMBER])
	reversed := tgr.Reverse(t3)
	for i := 0; i < len(reversed); i++ {
		if (reversed[i][0] > '0' && reversed[i][0] <= '9') && reversed[i][len([]rune(reversed[i]))-1] == ')' {
			n, _ := strconv.Atoi(reversed[i][:1]) //[:1] reads as string , [0] reads as a solitary byte, not wanted
			if strings.Contains(reversed[i+1], "(up") {
				for j := 1; j <= n; j++ {
					reversed[(i+1)+j] = strings.ToUpper(reversed[(i+1)+j])
				}
				reform := tgr.Reverse(reversed)
				rejoined := strings.Join(reform, " ")
				final := tgr.RemoveBrackets(rejoined)
				fmt.Println(final)

			} else if strings.Contains(reversed[i+1], "(low") {
				for j := 1; j <= n; j++ {
					reversed[(i+1)+j] = strings.ToLower(reversed[(i+1)+j])
				}
				reform := tgr.Reverse(reversed)
				rejoined := strings.Join(reform, " ")
				final := tgr.RemoveBrackets(rejoined)
				fmt.Println(final)
			} else if strings.Contains(reversed[i+1], "(cap") {
				for j := 1; j <= n; j++ {
					reversed[(i+1)+j] = strings.Title(reversed[(i+1)+j])
				}
				reform := tgr.Reverse(reversed)
				rejoined := strings.Join(reform, " ")
				final := tgr.RemoveBrackets(rejoined)
				_ = final //to omit
			}
		}
	}
	//PUNCTUATION
	srune := []rune(t2)

	for i := 0; i < len(srune); i++ {
		if tgr.IsPunc(string(srune[i])) && srune[i-1] == ' ' {
			srune[i], srune[i-1] = srune[i-1], srune[i]
		} else if srune[len(srune)-1] == 39 && srune[len(srune)-2] == ' ' {
			srune[len(srune)-2], srune[len(srune)-1] = srune[len(srune)-1], srune[len(srune)-2]
		} else if srune[i] == 39 && srune[i-1] == ' ' && srune[i+1] == ' ' {
			srune[i], srune[i+1] = srune[i+1], srune[i]
		}

	}
	// t2 = string(srune)
	// fmt.Printf("%#v\n", t2)
	// t2unspaced := strings.ReplaceAll(t2, "  ", " ")
	//  //final := strings.TrimSpace(t2unspaced)
	//  fmt.Printf("%#v", t4)

}
