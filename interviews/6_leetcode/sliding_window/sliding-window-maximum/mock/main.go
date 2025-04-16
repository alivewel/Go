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


func maxSlidingWindow(nums []int, k int) []int {
    queue := make([]int, 0)
    res := make([]int, len(nums) - k + 1)
    l, r := 0, -1
    for l < len(nums) {
        
    }
}
