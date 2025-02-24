package main

import (
	"fmt"
	"sort"
)

// решение с помощью двух указателей
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		l, r := i + 1, len(nums)-1
		if i > 0 && nums[i] == nums[i-1] {
            continue
        }
		for l < r {
			if nums[i] + nums[l] + nums[r] > 0 {
				r--
			} else if nums[i] + nums[l] + nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[r] == nums[r+1] {
                    r--
                }
				for l < r && nums[l] == nums[l-1] {
                    l++
                }
			}
		}
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

