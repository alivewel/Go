package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(double(arr))
}

func double(nums []int) []int {
	res := make([]int, len(nums))
	for i, _ := range nums {
		// fmt.Println(res[i])
		res[i] = num * 2
	}
	// _ = num
	return res
}
