package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1
	for l <= r {
		if numbers[l] + numbers[r] == target {
			return []int{l+1, r+1} // не забывать увеличить индексы на 1
		} else if numbers[l] + numbers[r] > target {
			r--
		} else {
			l++
		}
	}
    return []int{-1, -1} // не забывать вернуть [-1, -1] в случае неуспеха
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}