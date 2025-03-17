1) Забыл добавить проверку на сравнение элементов с полученными указателями и target.

Входные данные k = 6; nums = [5,7,7,8,8,10];
Ожидаемый результат [-1,-1]
Результат исполнения [1,0]

```
func search(nums []int, target int) []int {
	l, r := 0, len(nums)
	m := 0
    for r - l > 1 {
		m = (r + l) / 2
		if nums[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	l2, r2 := 0, len(nums)
	m2 := 0
    for r2 - l2 > 1 {
		m2 = (r2 + l2) / 2
		if nums[m2] < target {
			l2 = m2
		} else {
			r2 = m2
		}
	}
	return []int{r2, l}
}
```

2) Словил панику.

Входные данные k = 1; nums = [1];



```
func search(nums []int, target int) []int {
	l, r := 0, len(nums)
	m := 0
    for r - l > 1 {
		m = (r + l) / 2
		if nums[m] <= target {
			l = m
		} else {
			r = m
		}
	}
    
	l2, r2 := 0, len(nums)
	m2 := 0
    for r2 - l2 > 1 {
		m2 = (r2 + l2) / 2
		if nums[m2] < target {
			l2 = m2
		} else {
			r2 = m2
		}
	}
    if nums[l] == target && nums[r2] == target {
        return []int{r2, l}
    }
	return []int{-1, -1}
}
```
