package main

import (
	"fmt"
)

func search(nums []int, target int) int {
    offset := 0
	if !(nums[0] < nums[len(nums)-1]) { // если есть сдвиг
		l, r := -1, len(nums) - 1 // l, r := 0, len(nums)
		for r - l > 1 {
			m := (r + l) / 2
			if nums[m] > nums[len(nums)-1] { // if nums[m] < nums[len(nums)-1] { // <= // <=
				l = m // l = m
			} else {
				r = m // r = m
			}
		}
		offset = r
	}
	l, r := offset, len(nums) + offset
	for r - l > 1 {
		m := (r + l) / 2
		index := m % len(nums)
		if nums[index] <= target {
			l = m // r = m
		} else {
			r = m // l = m 
		}
	}
	if nums[l % len(nums)] == target {
		return l % len(nums) // return l
	}
	return -1
}

func main() {
	arr := []int{4,5,6,7,0,1,2}
	target := 0 // Output 4
	index := search(arr, target)
	fmt.Println("index:", index)
}
