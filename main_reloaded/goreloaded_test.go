package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// This test file tests the go-reloaded project against the test cases on audit page
func TestGoReloaded(t *testing.T) {
	inputFile, outputFile := "test.txt", "output.txt"

	/*	Set the args that will be seen by the go-reloaded file to the file name, input
		file and output file */
	os.Args = []string{"cmd", inputFile, outputFile}

	/*	Each element of testCases contains a pair of strings, the first string of the
		pair is what is to be written to the input file, the second is what should be
		the contents of the output file */
	testCases := [][]string{
		{"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure"},
		{"Don't be sad ,because sad backwards is das . And das not good",
			"Don't be sad, because sad backwards is das. And das not good"},
		{"harold wilson (cap, 2) : ' I’m a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I’m an optimist, but an optimist who carries a raincoat.'"},
	}

	for _, testCase := range testCases {
		// Write the string to be processed to the input file
		if err := ioutil.WriteFile(inputFile, []byte(testCase[0]), os.ModePerm); err != nil {
			panic(err)
		}

		// CHANGE THIS FUNCTION TO YOUR FUNCTION NAME IF DIFFERENT
		GoReloaded()

		/*	Read the contents of the output file, checking if it is equal to the
			expected output */
		if result, err := ioutil.ReadFile(outputFile); err != nil {
			panic(err)
		} else if string(result) != testCase[1] {
			t.Errorf("\nTest fails when given the test case:\n\t\"%s\","+
				"\n%s should contain:\n\t\"%s\",\n%s contains:\n\t\"%s\"\n\n",
				testCase[0], outputFile, testCase[1], outputFile, string(result))
		}
	}
}
