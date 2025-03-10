package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
    l, r := 0, len(nums) - 1
    minDiff := abs(nums[r] - nums[l])
    for l < r {
        curDiff := abs(nums[r] - nums[l])
        if curDiff < minDiff {
            minDiff = curDiff
        }
        if nums[r] > nums[l] {
            r--
        } else {
            l++
        }
    }
    return minDiff
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

// некорректное решение

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(minimumDifference(nums, 2))
}
