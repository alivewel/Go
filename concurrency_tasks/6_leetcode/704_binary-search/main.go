package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	mid, low, high := 0, 0, len(nums)-1
	for low <= high {
		mid = (low + high) / 2
		fmt.Println(mid, low, high)
		if target == nums[mid] {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12} // Пример массива цен
	target := 9
	index := search(arr, target)
	fmt.Println("index:", index) // Выводим результат
}

// https://medium.com/@arijitnath92/leetcode-121-best-time-to-buy-and-sell-stock-edd10e9885de
