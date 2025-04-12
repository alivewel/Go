package main

// first argument to append must be a slice; have allComb (variable of type *[][]int)
// забыл звездочки добавить
func bruteForce(idx int, nums []int, curComb []int, allComb *[][]int) {
    if idx == len(nums) {
        // здесь нужно копию curComb добавлять в allComb
        copySubset := make([]int, len(curComb))
        copy(copySubset, curComb)
        *allComb = append(*allComb, copySubset) // res = append(res, curComb)
        return
    }
    bruteForce(idx+1, nums, curComb, allComb)
    curComb = append(curComb, nums[idx])
    bruteForce(idx+1, nums, curComb, allComb)
    // забыл обрезать длину curComb после вставки
    curComb = curComb[:len(curComb)-1]
}

func generateSubsets(nums []int) [][]int {
    res := make([][]int, 0)
    bruteForce(0, nums, []int{}, &res)
    return res
}
