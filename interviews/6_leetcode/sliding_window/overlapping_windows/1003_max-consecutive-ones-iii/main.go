package main

import "fmt"

func longestOnes(nums []int, k int) int {
    l, r := 0, -1
    maxCount := 0
    countZero := 0
    for l < len(nums) {
        for r + 1 < len(nums) && (countZero < k || nums[r+1] == 1) {
            if nums[r+1] == 0 {
                countZero++
            }
            r++
        }
		windowSize := r - l + 1
        maxCount = max(windowSize, maxCount)
		if nums[l] == 0 {
            countZero--
        }
        l++
    }
    return maxCount
}

func main() {
	nums := []int{1,1,1,0,0,0,1,1,1,1,0}
	k := 2
	fmt.Println(longestOnes(nums, k)) // Expected 6
}
