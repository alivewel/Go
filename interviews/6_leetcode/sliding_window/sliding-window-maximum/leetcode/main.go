package main

import (
	"fmt"
)

//         |    |
// nums = [1,3,-1,-3,5,3,6,7]
// queue = [3 -1] - удалить 1

//             |   |
// nums = [1 1 3 2 1 1]
// queue = [3 2 1] - удалить все что меньше 3
// удалили две 1 добавили 3
// peek = 3
// if peek == nums[l]
//   de queue

// делаем операцию Enqueue до тех пор пока вытащенный элемент из очереди меньше текущего  
// в очереди хранить индекс


// queue = [3
//              |   | 
// [1,3,-1,-3,5,3,6,7]

func maxSlidingWindow(nums []int, k int) []int {
    queue := make([]int, 0)
    res := make([]int, 0, len(nums) - k) // len(nums) - k + 1 // res := make([]int, len(nums) - k + 1)
    l, r := 0, -1
    for l + k - 1 < len(nums) {
        for r + 1 < len(nums) && r - l + 1 < k {
            for len(queue) > 0 && nums[r+1] > queue[0] {
                fmt.Println(nums[r+1])
                queue = queue[1:]
            }
            queue = append(queue, nums[r+1])
            r++
        }
        fmt.Println(l, r, queue)
        // самый первый элемент в очереди должен быть максимум окна
        // вытаскиваем из очереди элемент
        if len(queue) > 0 {
            res = append(res, queue[0])
            fmt.Println(queue[0])
        }
        if len(queue) > 0 && queue[0] == nums[l] {
            queue = queue[1:] // вытаскиваем элемент из очереди
        }
        l++
    }
    return res
}

func main() {
	// nums := []int{1,3,-1,-3,5,3,6,7}
    // nums := []int{1,3,-1,-3,5,3,6,7} // Expected [3,3,5,5,6,7]
    nums := []int{1,3,1,2,0,5}
    k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
