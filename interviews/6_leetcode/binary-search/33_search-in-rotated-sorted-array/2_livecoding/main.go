package main

import (
	"fmt"
)

func search(nums []int, target int) int {
    offset := 0 
    if !(nums[0] < nums[len(nums)-1]) {
        l, r := -1, len(nums) - 1 // l, r := -1, len(nums)
        for r - l > 1 {
            m := (r + l) / 2
            if nums[len(nums) - 1] < nums[m] {
                l = m
            } else {
                r = m
            }
        }
        offset = r // offset = l
    }
    // l, r := offset, len(nums) + offset + 1
    l, r := offset, len(nums) + offset
    // l = 3 r = 9 m = 4 index = 4
    // l = 4 r = 9 m = 6 index = 1
    // l = 6 r = 9 m = 7 index = 2
    //      |       |
    // [1,2,3,4,5]
    // l = 0 r = 6 m = 2
    // l = 2 r = 6 m = 4
    for r - l > 1 {
        m := (r + l) / 2
        index := m % len(nums)
        if nums[index] <= target { // nums[index] < target
            l = m
        } else {
            r = m
        }
    }
    index := l % len(nums) // 2
    if nums[index] == target {
        return index
    }
    return -1
}


func main() {
	arr := []int{4,5,6,7,0,1,2}
	target := 0 // Output 4
	index := search(arr, target)
	fmt.Println("index:", index)
}
