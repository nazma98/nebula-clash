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

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])

		var rem, num int
		num = 0

		slice := make([]int, 0)

		for a != 0 {
			rem = a % b
			a = a / b
			slice = append(slice, rem)
		}

		for i := len(slice) - 1; i >= 0; i-- {
			num = num*10 + slice[i]
		}
		fmt.Println(num)

	}
}
