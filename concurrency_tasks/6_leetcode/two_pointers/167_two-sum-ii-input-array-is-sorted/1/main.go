package main

import "fmt"

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l != r {
		if nums[l]+nums[r] > target {
			r--
		} else if nums[l]+nums[r] < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return nil
}

func main() {
	// nums := []int{2, 7, 11, 15}
	// nums := []int{3, 3}
	nums := []int{-2, 1, 6, 9, 12, 25, 101}
	res := twoSum(nums, 18)
	fmt.Println(res)
}
