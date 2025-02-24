1) r++ в цикле почему-то на шаг вперед делает в конце. 
С помощью дебага определил, что цикл доходит до индеков r = 2 r = 6. А должно быть r = 1 r = 2 r = 5. Один из диапозонов просто пропускается.

Решил проблему добавлением r-- после цикла. Это помогло, решение прошло тесты, но это некрасиво. 

``` go
func findMaxConsecutiveOnes(nums []int) int {
	l, r := 0, 0
	countOnes := 0
	for r < len(nums) {
		for r < len(nums) && nums[l] == nums[r] {
			r++
		}
		if nums[l] == 1 {
			curCount := r - l + 1
			if curCount > countOnes {
				countOnes = curCount
			}
		}
		l = r + 1
		r = l
	}
    return countOnes 
}
```

2) Решение рабочее, но некрасивое.

``` go
func findMaxConsecutiveOnes(nums []int) int {
	l, r := 0, 0
	countOnes := 0
	for r < len(nums) {
		for r < len(nums) && nums[l] == nums[r] {
			r++
		}
		r--
		if nums[l] == 1 {
			curCount := r - l + 1
			if curCount > countOnes {
				countOnes = curCount
			}
		}
		l = r + 1
		r = l
	}
    return countOnes 
}
```

3) Забавно, что решение уходит в бесконечный цикл.

``` go
func findMaxConsecutiveOnes(nums []int) int {
	l, r := 0, 0
	countOnes := 0

	for r < len(nums) {
		// Ищем начало последовательности
		if nums[l] == 1 {
			for r < len(nums) && nums[l] == nums[r] {
				r++
			}
			// Проверяем длину найденной последовательности
			curCount := r - l
			if curCount > countOnes {
				countOnes = curCount
			}
		}
		// Переход к следующему сегменту
		l = r
	}

	return countOnes
}
```

4) Решение от Макса:

``` go
func findMaxConsecutiveOnes(nums []int) int {
	l := 0
	r := 0
	result := 0
	for l < len(nums) {
		// before
		// l
		// 1 1 1 1 0 0 1 1
		// r

		// бежим правым указателем пока в интервале [l, r]
		// находятся все последовательные числа
		for r+1 < len(nums) && nums[r] == nums[r+1] {
			r += 1
		}
		// after
		// l
		// 1 1 1 1 0 0 1 1
		//       r

		// обновляем ответ только если в палавающем окне были 1-цы
		// нас не просят искать нули, поэтому окна с нулями игнорируем
		if nums[r] == 1 {
			result = max(result, r-l+1)
		}

		// интервалы не пересекаются, поэтому сдвигаем
		// на r + 1 - именно отсюда будет начинаться
		// следующий интервал
		l = r + 1
		r = r + 1
	}
	return result
}
```
``` go
for r+1 < len(nums) && nums[r] == nums[r+1] { 
    // Первое условие (r+1 < len(nums)) гарантирует, что r+1 не выйдет за границы массива, 
    // предотвращая обращение к несуществующему индексу nums[r+1].
    
    // Второе условие (nums[r] == nums[r+1]) проверяет, что текущий элемент равен следующему.
    // Это позволяет увеличивать r, пока идет последовательность одинаковых чисел.
    
    // Когда nums[r] != nums[r+1], цикл завершается, и r останется на последнем элементе группы.
    r++
}
```