package main

import "fmt"

// Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	prevElem, len := 0, 0
	for i, j := range nums {
		if i == 0 {
			prevElem = j
			len++
			continue
		}
		if j != prevElem {
			nums[len] = j
			len++
			prevElem = j
		}
	}
	return len
}

func main() {
	nums := []int{1, 2, 3, 4, 4}
	// nums := []int{}
	k := removeDuplicates(nums)
	fmt.Println(nums, k)
}
