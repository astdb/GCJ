// https://code.google.com/codejam/contest/6254486/dashboard

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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
				fmt.Fprintf(os.Stderr, "Error reading testcase count: %v\n", err)
				return
			}
			k++
			continue
		} else {
			if k > testcasecount {
				break
			}
		}

		// read next testcase
		n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading testcase: %v\n", err)
			return
		}

		// process case: start counting
		i := 1
		// n_orig := n
		digits := []int{}
		for !complete(digits) {
			val := n * i // OVERFLOW??
			val_digits := getDigits(val)

			if (i > 1) && (n == val) {
				fmt.Printf("Case #%d: INSOMNIA\n", k)
				break
			}

			for _, v := range val_digits {
				if !contains(digits, v) {
					digits = append(digits, v)

					if complete(digits) {
						fmt.Printf("Case #%d: %d\n", k, val)
					}
				}
			}

			i++
		}

		// increment testcase count
		k++
	}
}

// check if an int slice contains a given number
func contains(digits []int, n int) bool {
	for _, v := range digits {
		if v == n {
			return true
		}
	}

	return false
}

// check if a given int slice has all individual decimal digits
func complete(digits []int) bool {
	if len(digits) != 10 {
		return false
	}

	sort.Ints(digits)
	j := 0
	for _, v := range digits {
		if v != j {
			return false
		}

		j++
	}

	return true
}

// takes a positive integer and returns an int slice containing all unique decimal digits of it e.g. 54682 -> []int{5,4,6,8,2}
func getDigits(n int) []int {
	if n < 0 {
		return nil
	}

	rem := n % 10
	n = n / 10
	digits := []int{}
	digits = append(digits, rem)

	for n > 0 {
		rem = n % 10
		n = n / 10

		if contains(digits, rem) {
			continue
		} else {
			digits = append(digits, rem)
		}
	}

	return digits
}
