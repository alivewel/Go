package main

import "fmt"

// Rotate Array
func rotate(nums []int, k int) {
	if k < 0 {
		k = len(nums) + k
	}
	length := len(nums)
	k = k % length
	fmt.Println(k)
	newArray := make([]int, length)
	for i := range nums {
		newIndex := (i + k) % length
		fmt.Println(i, newIndex)
		newArray[newIndex] = nums[i]
	}
	for i, j := range newArray {
		nums[i] = j
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	// nums := []int{1, 2, 3, 4, 4}
	// nums := []int{1, 2, 3, 4, 5}
	// nums := []int{7,1,5,3,6,4}
	// nums := []int{7,6,4,3,1}
	// nums := []int{}
	k := -6
	rotate(nums, k)
	fmt.Println(nums)
}
