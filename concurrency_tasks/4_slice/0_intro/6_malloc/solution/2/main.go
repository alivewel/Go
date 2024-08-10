package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(double(arr))
}

func double(nums []int) []int {
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = num * 2
	}
	return res
}
