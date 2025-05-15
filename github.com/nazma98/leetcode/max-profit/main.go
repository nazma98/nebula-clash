package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {

			if prices[i] < prices[j] && profit < (prices[j]-prices[i]) {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
