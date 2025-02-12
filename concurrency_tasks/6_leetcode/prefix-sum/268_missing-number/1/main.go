package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2 // сумма арифметической прогрессии
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums)) // 8
	nums = []int{3,0,1}
	fmt.Println(missingNumber(nums)) // 2
	nums = []int{0,1}
	fmt.Println(missingNumber(nums)) // 2
}
