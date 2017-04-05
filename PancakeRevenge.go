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
			// fmt.Printf("%d\n", testcasecount)
			k++
			continue
		} else {
			// only need to read and process the specified number of testcases
			if k > testcasecount {
				break
			}
		}

		// read testcase
		pancakestack := strings.TrimSpace(scanner.Text())
		
		flipcount := 0
		var top byte
		
		for i := 0; (i < len(pancakestack)) && mustFlip(pancakestack); i++ {
			// fmt.Printf("-----------------\ni = %d", i)
			// fmt.Printf("\ni = %d\n", i)
			if i == 0 {
				// remember the top cake
				top = pancakestack[i]
				// fmt.Printf("Top cake is <%c>\n", rune(top))

				if len(pancakestack) == 1 {
					// fmt.Printf("Single-length stack and must be flipped - flipping..\n")
					pancakestack = flipStack(pancakestack, 1)
					flipcount++
					// fmt.Printf("Flipcount: %d, stack: %s\n", flipcount, pancakestack)
					break
				}

				continue
			}

			if pancakestack[i] != top {
				// flip from top to i-1, which is the top i pancakes
				// fmt.Printf("Change from top detected at i = %d (stack: %s) - flipping the top %d...\n", i, pancakestack, i)
				pancakestack = flipStack(pancakestack, i)
				// pancakestack_ = []rune(pancakestack)
				flipcount++
				// fmt.Printf("Flipcount: %d, stack: %s\n", flipcount, pancakestack)
				i = -1	// check from the start 
			}

			if i == len(pancakestack)-1 && top == '-' {
				pancakestack = flipStack(pancakestack, len(pancakestack))
				flipcount++
				// fmt.Printf("Flipcount: %d, stack: %s\n", flipcount, pancakestack)
				
			}
		}
		
		fmt.Printf("Case #%d: %d\n", k, flipcount)
		k++
	}
}

// 'x-rays' a stack of pancakes and decides if it needs flipping of not
func mustFlip(stack string) bool {
	for _, pancake := range stack {
		if pancake == '-' {
			// found cake with happy side down - must flip
			return true
		}
	}

	return false
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
