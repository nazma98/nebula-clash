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

	x, _ := strconv.Atoi(parts[0])
	t, _ := strconv.Atoi(parts[1])

	if x <= t {
		fmt.Println(0)
	} else {
		fmt.Println(x - t)
	}
}
