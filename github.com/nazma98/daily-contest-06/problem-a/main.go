package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	line, _ := in.ReadString('\n')

	line = strings.TrimSpace(line)

	var slice []int

	for _, ch := range line {
		slice = append(slice, int(ch))
	}

	sort.Ints(slice)

	var diff int

	for i := 1; i < len(slice); i++ {
		if slice[i-1] != slice[i] {
			diff++
		}
	}

	if diff == 1 {

		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
	out.Flush()
}
