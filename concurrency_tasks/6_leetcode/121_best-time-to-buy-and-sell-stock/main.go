package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4} // Пример массива цен
	profit := maxProfit(prices)
	fmt.Println("Максимальная прибыль:", profit) // Выводим результат
}

// https://medium.com/@arijitnath92/leetcode-121-best-time-to-buy-and-sell-stock-edd10e9885de
