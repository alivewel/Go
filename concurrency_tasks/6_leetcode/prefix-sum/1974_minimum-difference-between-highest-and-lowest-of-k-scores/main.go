package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	result := math.MaxInt
	for i := range nums {
		l := i
		r := i + k - 1
		curResult := nums[r] - nums[l]
		if curResult < result {
			result = curResult
		}
		if i == len(nums) - k {
			break
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(minimumDifference(nums, 2))
}
