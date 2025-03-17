package main

import (
	"fmt"
)

// задача решается через 2 бинарных поиска с разными функциями
// в первой функции условие x <= target
// во второй функции условие x <= target

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)
	m := 0
	// берем левый указатель
    for r - l > 1 {
		m = (r + l) / 2
		// if nums[m] <= target {
		if nums[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(l, r)
	return nil
	//                 |
	// []int{1,2,2,2,2,2,5,5,8,19}
	// l, r = 0, 10; m = 5
	//           |
	// []int{1,2,2,2,2,2,5,5,8,19}
	// l, r = 5, 10;  m = 7
	// l, r = 5, 7;  m = 6
}

func main() {
	// arr := []int{1,2,2,2,2,2,5,5,8,19}
	
	target := 2
	index := searchRange(arr, target)
	fmt.Println("index:", index)
}
