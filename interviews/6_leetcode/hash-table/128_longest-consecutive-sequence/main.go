func longestConsecutive(nums []int) int {
    setNums := make(map[int]bool, 0)
    for _, num := range nums {
        setNums[num] = true
    }
    maxLength := 0
    for num := range setNums {
        if !setNums[num-1] { // !setNums[num]
            currNum := num
            currLength := 0
            for setNums[currNum] {
                currNum++
                currLength++
            }
            if maxLength < currLength {
                maxLength = currLength 
            }
        }
    }
    return maxLength
}
