package main

import (
	"fmt"
	"sort"
)

// решение с помощью хэш-таблицы
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		// mapSum := make(map[[3]int]struct{}) 
		mapSum := make(map[int]struct{}) 
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			complement := target - nums[j] // complement - дополнение
			if _, found := mapSum[complement]; found {
				// mapSum[[3]int{nums[i], complement, nums[j]}] = struct{}{}
				res = append(res, []int{nums[i], complement, nums[j]})
			}
			mapSum[nums[j]] = struct{}{}
		}	
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//      sort: -4, -1, -1, 0, 1, 2
	// nums := []int{-2,0,0,2,2}
	//      sort: -2,0,0,2,2
	fmt.Println(threeSum(nums))
}

