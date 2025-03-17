package main

import (
	"fmt"
)

// задача решается через 2 бинарных поиска с разными функциями
// в первой функции условие x <= target
// во второй функции условие x < target

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)
	m := 0
	// берем последний левый указатель
    for r - l > 1 {
		m = (r + l) / 2
		if nums[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	l2, r2 := -1, len(nums) - 1
	m2 := 0
	// берем первый правый указатель
    for r2 - l2 > 1 {
		m2 = (r2 + l2) / 2
		if nums[m2] < target {
			l2 = m2
		} else {
			r2 = m2
		}
	}
	fmt.Println(r2, l)
	if l < len(nums) && nums[l] == target && r2 < len(nums) && nums[r2] == target {
	// if l < len(nums) && nums[l] == target && l2 < len(nums) && nums[l2] == target {
        return []int{r2, l}
    }
	return []int{-1, -1}
	//                 |
	// []int{1,2,2,2,2,2,5,5,8,19}
	// l, r = 0, 10; m = 5
	//           |
	// []int{1,2,2,2,2,2,5,5,8,19}
	// l, r = 5, 10;  m = 7
	// l, r = 5, 7;  m = 6
}

func main() {
	arr := []int{1,2,2,2,2,2,5,5,8,19}
	target := 2
	// arr := []int{1}
	// target := 1
	index := searchRange(arr, target)
	fmt.Println("index:", index)
}
