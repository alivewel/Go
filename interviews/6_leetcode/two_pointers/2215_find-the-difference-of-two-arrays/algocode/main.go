package main

import (
	"fmt"
)

func findDifference(nums1 []int, nums2 []int) []int {
	p1, p2 := 0, 0
	result := make([]int, 0)
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			result = append(result, nums1[p1])
			p1++
		} else if nums1[p1] > nums2[p2] {
			result = append(result, nums2[p2])
			p2++
		} else {
			p1++
			p2++
		}
	}
	if p1 < len(nums1) {
		for p1 < len(nums1) {
			result = append(result, nums1[p1])
			p1++
		}
	}
	if p2 < len(nums2) {
		for p2 < len(nums2) {
			result = append(result, nums2[p2])
			p2++
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 5, 7, 9}
	nums2 := []int{2, 3, 5, 6, 7, 8}
	fmt.Println(findDifference(nums1, nums2))
}