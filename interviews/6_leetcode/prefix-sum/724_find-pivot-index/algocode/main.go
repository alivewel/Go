package main

import "fmt"

func pivotIndex(nums []int) int {
	allSum := 0
	pxSum := 0
	for i := 0; i < len(nums); i++ {
		allSum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		// pxSum += nums[i] // неправильно
		if allSum - nums[i] - pxSum == pxSum {
			return i
		}
		pxSum += nums[i]
	}

	return -1 // возвращать -1, если ничего не нашли
}

func main() {
	nums := []int{7, 3, 4, 5, 5}
	fmt.Println(pivotIndex(nums))
}
