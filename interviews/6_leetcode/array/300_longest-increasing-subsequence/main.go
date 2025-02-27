package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	maxLen, curLen := 0, 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] < nums[i+1] {
			curLen++
		} else {
			// curLen = 0
			maxLen = max(maxLen, curLen) // max(maxLen, curLen)
			curLen = 0
		}
	}
    return maxLen
}

func main() {
	// nums := []int{0,1,2,4,5,7} //
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(nums))
}
