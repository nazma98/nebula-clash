package main

import "fmt"

func main() {
	slice := []int{-8, -2, 5, -1, 4, -6, 2}

	maxSum := -9999999
	currentSum := 0

	for _, val := range slice {
		currentSum += val
		if maxSum < currentSum {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}

	fmt.Println(maxSum)

	fmt.Println(slice)
}
