package main

import "fmt"

func moveZeroes(nums []int) {
    for i := range nums {
		if nums[i] == 0 {
			fast := i
			for fast < len(nums) { 
				if nums[fast] != 0 {
					nums[i] = nums[fast]
					nums[fast] = 0
					break
				}
				fast++
			}
		}
	}

}

func main() {
	nums := []int{0, 2, 0, 0, 4, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}