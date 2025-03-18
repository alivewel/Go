package main

import (
	"fmt"
)

func search(nums []int, target int) int {
    offset := 0
	if !(nums[0] < nums[len(nums)-1]) { // если есть сдвиг
		l, r := 0, len(nums)
		for r - l > 1 {
			m := (r + l) / 2
			if nums[m] <= nums[len(nums)-1] {
				r = m // уточнить
			} else {
				l = m // уточнить
			}
		}
		offset = l
	}

	l, r := offset, len(nums) + offset
	for r - l > 1 {
		m := (r + l) / 2
		index := m % len(nums)
		if nums[index] <= target {
			r = m // уточнить
		} else {
			l = m // уточнить
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func main() {
	arr := []int{4,5,6,7,0,1,2}
	target := 0 // Output 4
	index := search(arr, target)
	fmt.Println("index:", index)
}
