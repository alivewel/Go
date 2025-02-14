package main

import "fmt"
 
func maxProduct(nums []int) int {
	maxVal1, maxVal2 := 0, 0
	for _, num := range nums {
		if num > maxVal1 {
			if maxVal2 < maxVal1 {
				maxVal2 = maxVal1
			}
			maxVal1 = num
			continue
		}
		if num > maxVal2 {
			maxVal2 = num
		}
	}
	return (maxVal1 - 1) * (maxVal2 - 1)
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(maxProduct(nums))
}
