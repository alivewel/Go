1) 
1 ошибка: В цикле движения указателя было: nums[r] == nums[r+1] + 1
2 ошибка: Для перевода в строку из числа использовал string(). Правильно использовать strconv.Itoa().
3 ошибка: Вместо nums[l] пытался вставить l в результирующий массив.

```
func summaryRanges(nums []int) []string {
	res := []string{}
	l, r := 0, 0
    for r < len(nums) {
		for r < len(nums) - 1 && nums[r] == nums[r+1] - 1 { // 
			r++
		}
		if r > l {
			// res = append(res, string(l)+"->"+string(r))
			res = append(res, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r]))
		} else {
			// res = append(res, string(l))
			res = append(res, strconv.Itoa(nums[l]))
		}
		l = r + 1
		r = r + 1
	}
	return res
}
```