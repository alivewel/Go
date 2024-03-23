package main

import "fmt"

// Best Time to Buy and Sell Stock II
func maxProfit(prices []int) int {
	profit := 0
	for i := range prices {
		if i == 0 {
			continue
		}
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	// prices := []int{1, 2, 3, 4, 4}
	// prices := []int{1,2,3,4,5}
	// prices := []int{7,1,5,3,6,4}
	prices := []int{7,6,4,3,1}
	// prices := []int{}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
