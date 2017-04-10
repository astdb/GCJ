// Google CodeJam 2017 - Qualification Round - Oersized Pancake Flipper 
// https://code.google.com/codejam/contest/3264486/dashboard#s=p0

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ensure input filename entered on command line
	if len(os.Args) <= 1 {
		fmt.Println("Usage: > go run OversizedFlipper.go <input.file>")
		return
	}

	// capture input file name
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return
	}
	defer file.Close()

	// setup scanner to read input file
	scanner := bufio.NewScanner(file)

	k := 0 // line numbers
	testcasecount := 0
	for scanner.Scan() {
		if k == 0 {
			// read number of testcases
			testcasecount, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading testcase count: %v\n", err)
				return
			}
			fmt.Printf("%d\n", testcasecount)
			k++
			continue
		} else {
			// only need to read and process the specified number of testcases
			if k > testcasecount {
				break
			}
		}

		// read testcase - number Tatiana counts up to from 1 (including limits)
		testcase := strings.TrimSpace(scanner.Text())
		testcase_data := strings.Split(testcase, " ")

		if len(testcase_data) == 2 {
			pancakeLine := testcase_data[0]
			flipperLength, err := strconv.Atoi(testcase_data[1])

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error acquiring testcase %d flipperlength - TC: %s\n", k, testcase)
				return
			}

			fmt.Printf("%s %d\n", pancakeLine, flipperLength)

		} else {
			fmt.Fprintf(os.Stderr, "Error acquiring testcase %d data -  TC: %s\n", k, testcase)
			// return
			continue
		}
	}
}
