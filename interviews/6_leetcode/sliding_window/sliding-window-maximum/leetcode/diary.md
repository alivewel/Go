```
func maxSlidingWindow(nums []int, k int) []int {
    queue := make([]int, 0)
    res := make([]int, 0, len(nums) - k) // len(nums) - k + 1 // res := make([]int, len(nums) - k + 1)
    l, r := 0, -1
    // for l < len(nums) {
    for l + k - 1 < len(nums) {
        
        // sizeWindow := r - l + 1 // нужно в цикле явно смотреть, иначе она не обновляется
        
        // for r + 1 < len(nums) && r - l + 1 <= k { // r - l + 1 <= k
        for r + 1 < len(nums) && r - l + 1 < k {
            // for len(queue) > 0 && nums[r+1] > queue[0] {
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
        if len(queue) > 0 { // поменял местами 2 if
            res = append(res, queue[0])
            fmt.Println(queue[0])
        }
        if len(queue) > 0 && queue[0] == nums[l] { // поменял местами 2 if
            queue = queue[1:] // вытаскиваем элемент из очереди
        }
        l++ // забыл двигать левый указатель
    }
    return res
}
```
1) Количество итераций должно быть меньше, иначе результат будет меньше.
Имеется:
for l < len(nums) 
Должно быть:
for l + k - 1 < len(nums)

2) Нужно в цикле явно смотреть, иначе она не обновляется
Имеется:
sizeWindow := r - l + 1 
for r + 1 < len(nums) && sizeWindow < k 
Должно быть:
for r + 1 < len(nums) && r - l + 1 < k

3) Условие движения правой границы правого окна некорректное

Имеется:
for r + 1 < len(nums) && r - l + 1 <= k
Должно быть:
for r + 1 < len(nums) && r - l + 1 < k

4) l++ // забыл двигать левый указатель

5) Поменял местами 2 if
if len(queue) > 0 && queue[0] == nums[l]
if len(queue) > 0