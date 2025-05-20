package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	for b != 0 {
		rem := a % b
		a = b
		b = rem
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		nums := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])

		res := gcd(a, b)
		fmt.Println(res)
	}
}
