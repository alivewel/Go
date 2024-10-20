package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indices := make([]int, 2) // индексы
	for index1, i := range nums {
		for index2, j := range nums {
			if index1 == index2 {
				continue
			}
			if i+j == target {
				indices[0] = index2
				indices[1] = index1
			}
		}
	}
	return indices
}

func main() {
	// nums := []int{2, 7, 11, 15}
	nums := []int{3, 3}
	res := twoSum(nums, 6)
	fmt.Println(res)
}
