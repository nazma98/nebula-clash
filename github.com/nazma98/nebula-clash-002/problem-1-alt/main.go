package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInts() []int {
	line := readLine()
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, field := range fields {
		nums[i], _ = strconv.Atoi(field)
	}
	return nums
}

func main() {
	defer writer.Flush()

	t, _ := strconv.Atoi(readLine())

	for i := 0; i < t; i++ {
		nums := readInts()

		cnt := 0
		for _, val := range nums {
			if val == 1 {
				cnt++
			}
			if val > 1 {
				cnt = 0
				break
			}
		}

		if cnt >= 5 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
