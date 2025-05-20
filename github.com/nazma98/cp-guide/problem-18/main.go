package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	a, _ := strconv.Atoi(parts[0])

	for a != 1 {
		if a%2 == 0 {
			a /= 2
		} else {
			a *= 3
			a += 1
		}
		fmt.Print(a, " ")
	}
	fmt.Println()
}
