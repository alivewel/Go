1) panic: runtime error: index out of range [6] with length 6

Нужно было поднять проверку if nums[l] != 0 выше, до того как мы l-указатель инкрементировали.

``` go
func longestStockGrowth(nums []int) int {
	l, r := 0, 0
	result := 0
	countZeros := 0
	for r < len(nums) {
		for r+1 < len(nums) && nums[r] == nums[r+1] || countZeros < 1 { 
			if nums[r] == 0 {
				countZeros++
			}
			r++
		}
		curCountZero := r - l + 1
		result = max(result, curCountZero)
		l = r + 1
		r = r + 1
		if nums[l] != 0 {
			countZeros = 0
		}
	}
    return result 
}
```

2) Забыл в цикле условие `nums[r] == nums[r+1] || countZeros < 1` в скобки взять.
Было:
for r+1 < len(nums) && nums[r] == nums[r+1] || countZeros < 1 
Нужно было:
for r+1 < len(nums) && (nums[r] == nums[r+1] || countZeros < 1) 

3) Много тестов зафейлилось. Мое решение неверное.
Не смог сам додуматься, полез смотреть решение.  
Не мог понять зачем изначально инициализировать zeroIdx := -1.

Проверка if len(stock) > 0 && stock[0] == 0:
Нужна для корректной обработки случая, когда первый элемент массива равен 0.

Инициализация zeroIdx := -1:
Нужна для обработки случаев, когда в окне нет нулей, чтобы избежать ошибок при попытке использовать индекс последнего нуля.