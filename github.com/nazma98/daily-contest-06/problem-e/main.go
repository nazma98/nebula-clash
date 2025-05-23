package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read number of test cases
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		// Read how many numbers to take
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		// Read the line of numbers
		scanner.Scan()
		fields := strings.Fields(scanner.Text())

		// Take first n numbers only
		numbers := make([]int, 3)
		for j := 0; j < 3; j++ {
			numbers[j], _ = strconv.Atoi(fields[j])
		}

		if numbers[n-1] == 0 {
			fmt.Println("Yes")
		} else {
			if numbers[n-1] == 2 && numbers[1] == 1 && numbers[0] == 0 {
				fmt.Println("Yes")
			} else if numbers[n-1] == 1 && numbers[0] == 3 && numbers[3] == 0 {
				fmt.Println("Yes")
			}
		}
		fmt.Println("No")
	}
}
