package main

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	l, r, index := 0, len(nums) - 1, len(nums) - 1
	for l <= r {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			result[index] = nums[l] * nums[l]
			l++
		} else {
			result[index] = nums[r] * nums[r]
			r--
		}	
		index--
	} 
	
	return result
}

func main() {
	nums := []int{-3, -2, 0, 1, 3, 5}
	fmt.Println(sortedSquares(nums))
}
