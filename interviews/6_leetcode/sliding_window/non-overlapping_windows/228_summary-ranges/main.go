package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	l, r := 0, 0
    for r < len(nums) {
		for r < len(nums) - 1 && nums[r] == nums[r+1] - 1 {
			r++
		}
		if r > l {
			res = append(res, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r]))
		} else {
			res = append(res, strconv.Itoa(nums[l]))
		}
		l = r + 1
		r = r + 1
	}
	return res
}

func main() {
	nums := []int{0,1,2,4,5,7} // ["0->2","4->5","7"]
	fmt.Println(summaryRanges(nums))
}
