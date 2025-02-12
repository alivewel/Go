package main

import "fmt"

// Rotate Array
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	rotateSubArray(nums, 0, n-1) // подумать
	rotateSubArray(nums, 0, k-1)
	rotateSubArray(nums, k, n-1)
}

func rotateSubArray(nums []int, l, r int) {
	for l < r { // подумать
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
