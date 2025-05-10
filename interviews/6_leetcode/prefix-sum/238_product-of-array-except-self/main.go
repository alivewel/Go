// [1,2,3,4]

// [1,2,6,24]
// [24,24,12,4]

// [24,12,8,6]

func productExceptSelf(nums []int) []int {
    preffSum := make([]int, len(nums))
    suffSum := make([]int, len(nums))
    for i, num := range nums {
        if i == 0 {
            preffSum[0] = nums[0]
        } else {
            preffSum[i] = num * preffSum[i-1]
        }
    }
    // for i, j := len(nums)-1, 0; i >= 0; i, j = i - 1, j + 1 {
    for i := len(nums)-1; i >= 0; i-- {
        if i == len(nums)-1 {
            suffSum[len(nums)-1] = nums[len(nums)-1]
        } else {
            suffSum[i] = nums[i] * suffSum[i+1]
        }
    }
    res := make([]int, len(nums))
    // [1,2,6,24]
    // [24,24,12,4]

    // [24,12,8,6]
    for i := range nums {
        if i == 0 {
            res[i] = suffSum[i+1]
        } else if i == len(nums)-1 {
            res[i] = preffSum[i-1]
        } else {
            res[i] = preffSum[i-1] * suffSum[i+1]
        }
    }
    return res
}
