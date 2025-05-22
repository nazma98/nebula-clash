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
	out := bufio.NewWriter(os.Stdout)

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])

	sum := a + b
	sub := a - b
	product := a * b

	if sum > sub {
		if sum > product {
			fmt.Println(sum)
		} else {
			fmt.Println(product)
		}
	} else {
		if sub > product {
			fmt.Println(sub)
		} else {
			fmt.Println(product)
		}
	}
	out.Flush()
}
