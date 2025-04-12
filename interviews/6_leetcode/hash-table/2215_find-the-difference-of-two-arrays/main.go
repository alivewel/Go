// nums1 = [1,2,3] - set1
// nums2 = [2,4,6] - set2
// arr1, arr2

// 1 1 1 1 1 1 1
// nums1 = [1,2,3,3] - set1 1, 2, 3
// nums2 = [1,1,2,2] - set2 1, 2

// time O(2N + 2M)
// mem O(N + M)

// 8:20
// 14:27

func findDifference(nums1 []int, nums2 []int) [][]int {
    set1 := make(map[int]struct{})
    set2 := make(map[int]struct{})
    for _, num := range nums1 {
       set1[num] = struct{}{} 
    }
    for _, num := range nums2 {
       set2[num] = struct{}{} 
    }
    res1 := make([]int, 0)
    for num := range set1 {
       if _, found := set2[num]; !found {
            res1 = append(res1, num)
       }
    }
    res2 := make([]int, 0)
    for num := range set2 {
       if _, found := set1[num]; !found {
            res2 = append(res2, num)
       }
    }
    return [][]int{
        res1,
        res2,
    }
}