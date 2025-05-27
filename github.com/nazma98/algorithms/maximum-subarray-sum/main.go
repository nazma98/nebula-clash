package main

import "fmt"

func main() {
	slice := []int{-8, -2, 5, -1, 4, -6, 2}

	maxSum := -9999999
	currentSum := 0

	startIndex := 0
	endIndex := 0
	tempStart := 0

	for i, val := range slice {
		currentSum += val

		if maxSum < currentSum {
			maxSum = currentSum
			startIndex = tempStart
			endIndex = i
		}

		if currentSum < 0 {
			currentSum = 0
			tempStart = i + 1
		}
	}

	fmt.Println(maxSum)
	fmt.Println(slice[startIndex : endIndex+1])
}
