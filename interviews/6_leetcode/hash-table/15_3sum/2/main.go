package main

import (
	"fmt"
	"sort"
)

// решение с помощью хэш-таблицы
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		mapSum := make(map[int]struct{}) 
		for j := i+1; j < len(nums); j++ {
			complement := target - nums[j] // complement - дополнение
			if _, found := mapSum[complement]; found {
				res = append(res, []int{nums[i], complement, nums[j]})
				for j < len(nums) - 1 && nums[j] == nums[j+1] {
					j++	
				}
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

