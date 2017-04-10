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
		fmt.Println("Usage: > go run TidyNumbers.go <input.file>")
		return
	}

	// capture input file name
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return
	}
	defer file.Close()

	// ints := []int{222, 8, 123, 555, 224488, 342}
	// for _, v := range ints {
	// 	if clean(v) {
	// 		fmt.Printf("%d clean\n", v)
	// 	} else {
	// 		fmt.Printf("%d not clean\n", v)
	// 	}
	// }
	// return

	// setup scanner to read input file
	scanner := bufio.NewScanner(file)

	k := 0 // line numbers
	testcasecount := 0
	prevlastClean := 1
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

		// read testcase - number Tatiana counts up to from 1 (including limits)
		countLimit, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
