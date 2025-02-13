package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	result := math.MaxInt
	for i := 0; i <= len(nums)-k; i++ { // нужно вовремя выйти из цикла, чтобы не выйти за границы массива
		curResult := nums[i+k-1] - nums[i]
		if curResult < result {
			result = curResult
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(minimumDifference(nums, 2))
}
