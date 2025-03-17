1) Забыл вернуть результат при нахождении target.

// Неправильный ответ
// k = 9; nums = [-1,0,3,5,9,12];
// Ожидаемый результат: 4

```
func search(nums []int, target int) int {
    l, r := 0, len(nums)
    for r - l > 1 {
        middle := (r + l) / 2
        if nums[middle] <= target { // nums[middle] < target
            r = middle
        } else {
            l = middle
        }
    }
    return -1
}
```

2) Решение ломается при кейсе arr := []int{5} и target := 5. Мы просто не зайдем в цикл, r = 1, l = 0. Условие r - l > 1 не выполняется. В этом случае мы возвращаем -1.
Пришлось добавить проверку после цикла - if nums[l] == target.

```
func search(nums []int, target int) int {
    l, r := 0, len(nums)
    for r - l > 1 {
        middle := (r + l) / 2
        if nums[middle] < target { // nums[middle] < target
            l = middle // + 1
        } else if nums[middle] > target {
            r = middle // - 1
        } else {
			return middle
		}
    }
    if nums[l] == target {
        return l
    }
    return -1
}
```