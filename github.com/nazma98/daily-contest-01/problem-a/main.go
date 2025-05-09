package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			fmt.Println("vowel")
		} else {
			fmt.Println("consonant")
		}
		break
	}
}
