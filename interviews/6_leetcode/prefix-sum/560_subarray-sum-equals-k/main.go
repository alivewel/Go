package main

import "fmt"

func subarraySum(nums []int, targetSum int) int {
    prefSum := make(map[int]int)
    prefSum[0] = 1
    curSum := 0
    countArr := 0
    for _, num := range nums { // for _, num := nums {
        curSum += num
        if count, found := prefSum[curSum-targetSum]; found { // prefSum[num-targetSum]
           countArr += count 
        }
        prefSum[curSum]++ // prefSum[num]++
    } 
    return countArr
}

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}
