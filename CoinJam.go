// https://code.google.com/codejam/contest/dashboard?c=6254486#s=p2

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
		fmt.Println("Usage: > go run CoinJam.go <input.file>")
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

	k := 0             // line numbers
	// testcasecount := 0 // only one testcase for both inputs for coinjam
	for scanner.Scan() {
		if k == 0 {
			// read test case number
			fmt.Printf("Case #: %s\n", strings.TrimSpace(scanner.Text()))
			k++
			continue
		}

		testcase_data := strings.Split(strings.TrimSpace(scanner.Text()), " ")

		if len(testcase_data) == 2 {
			cj_length, _ := strconv.Atoi(strings.TrimSpace(testcase_data[0])) // jamcoin length
			cj_num, _ := strconv.Atoi(strings.TrimSpace(testcase_data[1]))    // number of coins to generate
			fmt.Printf("%d %d\n", cj_length, cj_num)
		}
	}
}

// validateCoin takes a coin and checks its validity - if valid returns null error and slice of proof 
func validateCoin(coin String) (proof []int, error err) {

}
