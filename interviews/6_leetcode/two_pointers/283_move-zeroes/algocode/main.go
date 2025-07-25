package main

import "fmt"

// moveZeros перемещает нули в конец массива
func moveZeros(nums []int) []int {
    // указывает на какую позицию поставим следующий элемент не равный 0
    p1 := 0
    // указывает на следующий не нулевой элемент
    p2 := 0
    for p2 < len(nums) {
        if nums[p2] != 0 {
            nums[p1], nums[p2] = nums[p2], nums[p1]
            p1++
        }
        p2++
    }
    return nums
}

func main() {
	nums := []int{0, 2, 0, 0, 4, 0}
	moveZeros(nums)
	fmt.Println(nums)
}