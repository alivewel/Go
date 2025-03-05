package main

import "fmt"

func containsDuplicate(nums []int) bool {
    numsMap := make(map[int]struct{}, len(nums))
    for _, num := range nums {
        if _, found := numsMap[num]; found {
            return true
        }
        numsMap[num] = struct{}{}
    }
    return false
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(containsDuplicate(nums))
}
