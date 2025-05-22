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
	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)

	// Read number of test cases
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		// Read length of the slice
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		// Read the slice
		scanner.Scan()
		numStr := strings.Fields(scanner.Text())

		slice := make([]int, n)
		for j := 0; j < n; j++ {
			slice[j], _ = strconv.Atoi(numStr[j])
		}

		sort.Ints(slice)

		slice[0] += 1
		product := 1

		for _, val := range slice {
			product *= val
		}

		fmt.Println(product)
	}
}
