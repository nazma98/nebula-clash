package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	if t == 1 {
		fmt.Println("Hello World")
	} else {
		sum := 0
		for i := 0; i < t; i++ {
			scanner.Scan()
			n, _ := strconv.Atoi(scanner.Text())
			sum += n
		}
		fmt.Println(sum)
	}

}
