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

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil || n < 1 {
		return
	}

	candies := (n * (n + 1)) / 2
	fmt.Println(candies)
}
