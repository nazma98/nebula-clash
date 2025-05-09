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
	b, _ := strconv.Atoi(parts[1])

	if a+b < 10 {
		fmt.Println(a + b)
	} else {
		fmt.Println("error")
	}
}
