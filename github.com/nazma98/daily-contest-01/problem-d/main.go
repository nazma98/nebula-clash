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
	var payment int
	payBack := n / 15
	payment = (n * 800) - (payBack * 200)
	fmt.Println(payment)
}
