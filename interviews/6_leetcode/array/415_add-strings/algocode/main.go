package main
// [9, 9, 9, 9]
//       [1, 5]

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func calculate(nums1 []int, nums2 []int) []int {
    carry := 0
    p1, p2 := len(nums1) - 1, len(nums2) - 1
    // indexRes := max(len(nums1), len(nums2)) + 1
    // res := make([]int, indexRes)
    res := make([]int, 0)
    for carry != 0 || p1 >= 0 || p2 >= 0 {
        num1, num2 := 0, 0
        if p1 >= 0 {
            num1 = nums1[p1]
            p1--
        }
        if p2 >= 0 {
            num2 = nums2[p2]
            p2--
        }
        sum := num1 + num2 + carry
        res = append(res, sum % 10)
        carry = sum / 10
    }
    // reverse res
    for i, j := 0, len(res) - 1; i < j; i, j = i+1, j-1  { // i++, j--
        res[i], res[j] = res[j], res[i]
    }
    return res
}
