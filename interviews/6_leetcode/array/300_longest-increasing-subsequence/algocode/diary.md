1) 
``` go
func findLength(nums []int) int {
	maxLen, curLen := 0, 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] < nums[i+1] {
			curLen++
		} else {
			// curLen = 0
			maxLen = max(maxLen, curLen) // max(maxLen, curLen)
			curLen = 0
		}
	}
    return maxLen
}
```

2) 
``` go
func findLength(nums []int) int {
	maxLen, curLen := 1, 1
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] < nums[i+1] {
			curLen++
		} else {
			// curLen = 0
			 // max(maxLen, curLen)
			curLen = 1
		}
        maxLen = max(maxLen, curLen)
	}
    return maxLen
}
```