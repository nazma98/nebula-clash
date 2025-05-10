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

	n, _ := strconv.Atoi(parts[0])
	t, _ := strconv.Atoi(parts[1])
	a, _ := strconv.Atoi(parts[2])

	if t == a {
		fmt.Println("No")
	} else if t > a {
		if ((n - (t + a)) + a) > t {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	} else {
		if ((n - (t + a)) + t) > a {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}
