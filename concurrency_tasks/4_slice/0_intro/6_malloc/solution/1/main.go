package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(double(arr))
}

func double(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		res = append(res, num*2)
	}
	return res
}
