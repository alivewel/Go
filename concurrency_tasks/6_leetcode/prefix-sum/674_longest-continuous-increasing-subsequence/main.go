package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	maxLen := 1
	curLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curLen++
		} else {
			if maxLen < curLen {
				maxLen = curLen
			}
			curLen = 1
		}
	}
	if maxLen < curLen {
		maxLen = curLen
	}
	return maxLen
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums))
}
