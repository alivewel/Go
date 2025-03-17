package main

import (
	"fmt"
)

// бинарный поиск
// [4,5,6,7,0,1,2]

// len(nums)/2
// index = 7 / 2 = 3
// if nums[index] > target {
//    index = index + index / 2 + 1
//                  3 + 1 + 1
// (левый + правый индекс) / 2
// r - l + 1
// }
// [0,1,2,4,5,6,7]
// 7 + 0 / 2 = 3
// 3 + 7 / 2 = 5

// time O(log n)
// mem O(1)

// [-1,0,3,5,9,12] target = 2

func search(nums []int, target int) int {
    l, r := 0, len(nums) -1 // len(nums) // 0
    // middle := 0
    // for l + 1 < r {
    for l <= r {
        middle := l + (r-l)/2 
        if nums[middle] < target {
            l = middle + 1
        } else if nums[middle] > target {
            r = middle - 1
        } else {
            return middle
        }
    }
    return -1
}

// func search(nums []int, target int) int {
//     l, r := 0, len(nums)-1 // Set r to len(nums) - 1

//     for l <= r {
//         middle := l + (r-l)/2 // Correctly calculate the middle index
//         if nums[middle] < target {
//             l = middle + 1 // Move left boundary up
//         } else if nums[middle] > target {
//             r = middle - 1 // Move right boundary down
//         } else {
//             return middle // Target found
//         }
//     }
//     return -1 // Target not found
// }

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	
	target := 9
	index := search(arr, target)
	fmt.Println("index:", index)
}
