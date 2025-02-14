package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for i, num := range nums {
		if val, found := sumMap[target-num]; found {
			return []int{i, val}
		}
		sumMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
