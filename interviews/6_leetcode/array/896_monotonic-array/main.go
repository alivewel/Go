package main

import (
	"fmt"
	"strconv"
)
func isMonotonic(nums []int) bool {
	isInc, isDec := true, true
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			isInc = false
		}
		if nums[i] < nums[i+1] {
			isDec = false
		}
	}
    return isInc || isDec
}

func main() {
	nums := []int{0,1,2,4,5,7} //
	fmt.Println(isMonotonic(nums))
}
