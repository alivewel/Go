package main

import "fmt"

// Rotate Array
func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	newArray := make([]int, length)
	count := 0
	for i := range nums {
		if i+k < length {
			newArray[i] = nums[i+k]
		} else {
			newArray[i] = nums[count]
			count++
		}
	}
	for i, j := range newArray {
		nums[i] = j
	}
}

func main() {
	// nums := []int{1, 2, 3, 4, 4}
	nums := []int{1, 2, 3, 4, 5}
	// nums := []int{7,1,5,3,6,4}
	// nums := []int{7,6,4,3,1}
	// nums := []int{}
	k := 7
	rotate(nums, k)
	fmt.Println(nums)
}
