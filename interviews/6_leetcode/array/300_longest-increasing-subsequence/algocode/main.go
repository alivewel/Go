package main

import (
	"fmt"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func findLength(nums []int) int {
	maxLen, curLen := 1, 1
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] < nums[i+1] {
			curLen++
		} else {
			curLen = 1
		}
        maxLen = max(maxLen, curLen)
	}
    return maxLen
}

func main() {
	// nums := []int{0,1,2,4,5,7} //
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(findLength(nums))
}
