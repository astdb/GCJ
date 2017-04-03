
// https://code.google.com/codejam/contest/6254486/dashboard

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	// ensure input filename entered on command line
	if len(os.Args) <= 1 {
		fmt.Println("Usage: > go run CountingSheep.go <input.file>")
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
				fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
				return
			}
		} else {
			if k > testcasecount {
				break
			}
		}
		
		// read next line
		n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
			return
		}

		i := 1
		for {
			
			i++
		}

		k++
	}
}

// takes a positive integer and returns an int slice containing all unique decimal digits of it 
func getDigits(n int) []int {
	if n < 0 {
		return nil
	}

	digits := []int{}
	rem := n%10
	digits = append(digits, rem)

	for rem > 0 {
		rem = rem%2

		in := false
		for _, v := range digits {
			if v == rem {
				in = true
				break
			}
		}

		if !in {
			digits = append(digits, rem)
		}
	}
}
