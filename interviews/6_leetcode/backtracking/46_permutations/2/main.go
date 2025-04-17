package main

func bruteForce(nums []int, curComb []int, allComb *[][]int) {
    if len(nums) == 0 {
        copyComb := make([]int, 0, len(curComb))
        // copy(curComb, copyComb)
        copy(copyComb, curComb) // забыл про это
        *allComb = append(*allComb, copyComb)
        return
    }
    for i, num := range nums {
        curComb = append(curComb, num)
        // меняем местами текущий и последний элемент
        nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i] 
        // удаляем элемент последний
        nums = nums[:len(nums)-1]
        // вызываем бектрекинг без текущего элемента
        bruteForce(nums, curComb, allComb)
        // возвращаем на место текущий элемент
        nums = append(nums, num)
        // меняем местами текущий и последний элемент
        nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
        curComb = curComb[:len(curComb)-1] // забыл про это
    }
}

func permutations(nums []int) [][]int {
    if len(nums) == 1 {
        return [][]int{nums}
    }
    res := make([][]int, 0)
    bruteForce(nums, []int{}, &res)
    return res
}
