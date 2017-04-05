// https://code.google.com/codejam/contest/6254486/dashboard#s=p1

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
		fmt.Println("Usage: > go run PancakeRevenge.go <input.file>")
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
			if k > testcasecount {
				break
			}
		}

		pancakestack := strings.TrimSpace(scanner.Text())

		fmt.Printf("%s\n", pancakestack)

		for i := 1; i <= len(pancakestack); i++ {
			fmt.Printf("\tFlipping first %d pancake(s): %s\n", i, flipStack(pancakestack, i))
		}

		k++
	}
}

// flip the first n pancakes (positions 0-(n-1), top position being 0) of a given pancake stack
func flipStack(stack string, n int) string {
	// strings are immutable - turn it into a rune slice for flipping
	stack_runes := []rune(stack)

	if n > len(stack_runes) {
		return ""
	}

	if n == 1 {
		stack_runes[0] = flipCake(stack_runes[0])
		return string(stack_runes)
	}

	for i := 0; i <= (n - 1); i++ {
		if i == (n-1) {
			stack_runes[i] = flipCake(stack_runes[i])
		} else {
			temp := stack_runes[i]
			stack_runes[i] = flipCake(stack_runes[n-1])
			stack_runes[n-1] = flipCake(temp)
		}
		
		n--
	}

	return string(stack_runes)
}

func flipCake(pancake rune) rune {
	if pancake == '+' {
		return '-'
	} else {
		return '+'
	}
}
