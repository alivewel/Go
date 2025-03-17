package main

import (
	"fmt"
)

func search(nums []int, target int) int {
    l, r := 0, len(nums)
    for r - l > 1 {
        middle := (r + l) / 2
        // [-100,1,3,12,35,134]
        // target = 35
        // middle := (r + l) / 2 = 6 + 0 / 2 = 3; 12 <= 35
        // middle := (r + l) / 2 = 6 + 3 / 2 = 4; 35 <= 35

		// []int{-1,0,3,5,9,12} target = 9
		// middle := (r + l) / 2 = 6 + 0 / 2 = 3; 5 < 9
		// middle := (r + l) / 2 = 3 + 0 / 2 = 1; 0 < 9
        if nums[middle] < target { // nums[middle] < target
            l = middle // + 1
        } else if nums[middle] > target {
            r = middle // - 1
        } else {
			return middle
		}
    }
    if nums[l] == target {
        return l
    }
    return -1
}

func main() {
	// arr := []int{-1, 0, 3, 5, 9, 12}
	// arr := []int{-1,0,3,5,9,12} // Ожидаемый результат: 4
	// target := 9

    arr := []int{5}
    target := 5

	index := search(arr, target)
	fmt.Println("index:", index)
}
