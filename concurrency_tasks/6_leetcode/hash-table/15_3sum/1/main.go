package main

import (
	"fmt"
	"sort"
)

// решение с помощью двух указателей
func threeSum(nums []int) [][]int {
    // отсортировать массив
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		l, r := i + 1, len(nums)-1
		for l < r {
			if nums[i] + nums[l] + nums[r] > 0 {
				r--
			} else if nums[i] + nums[l] + nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// забыл сдвинуть указатели
				r--
				l++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//      sort: -4, -1, -1, 0, 1, 2
	fmt.Println(threeSum(nums))
}

