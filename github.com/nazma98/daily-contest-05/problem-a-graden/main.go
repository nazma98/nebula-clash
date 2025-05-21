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
	area := (a - 1) * (b - 1)
	fmt.Println(area)
}
