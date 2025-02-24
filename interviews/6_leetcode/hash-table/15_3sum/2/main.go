package main

import (
	"fmt"
	"sort"
)

// решение с помощью хэш-таблицы
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		
		
	}
	return res
}

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	//      sort: -4, -1, -1, 0, 1, 2
	nums := []int{-2,0,0,2,2}
	//      sort: -2,0,0,2,2
	fmt.Println(threeSum(nums))
}

