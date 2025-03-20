``` go
func merge(intervals [][]int) [][]int {
	// sort.Ints(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] // Сортируем по первому элементу (началу отрезка)
	})
	res := make([][]int, 0)
	res = append(res, []int{intervals[0][0], intervals[0][1]})
	for i := range intervals {
		if i == 0 {
			continue
		}
		// if max(intervals[i][0], intervals[i+1][0]) <= min(intervals[i][1], intervals[i+1][1]) {
		if max(res[len(res)-1][0], intervals[i][0]) <= min(intervals[len(res)-1][1], intervals[i][1]) {
			if len(res) == 0 {
				row := make([]int, 2)
				res = append(res, row)
			}
			// res = append(res, []int{min(intervals[i][0], intervals[i+1][0]),  max(intervals[i][1], intervals[i+1][1])})
			// res[len(res)-1][0] = min(intervals[i][0], intervals[i+1][0]) 
			// res[len(res)-1][1] = max(intervals[i][1], intervals[i+1][1]) 
			// res[len(res)-1][0] = min(res[len(res)-1][0], intervals[i+1][0]) 
			// res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i+1][1])
			res[len(res)-1][0] = min(res[len(res)-1][0], intervals[i][0]) 
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])  
		} else {
			res = append(res, intervals[i])
		}
	}
	
	return res
}
```

1) Не сразу догадался, как отсортировать двумерный массив. Для этого нужно использовать sort.Slice

```
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0] // Сортируем по первому элементу (началу отрезка)
})
```

2) Не сразу догадался, что нужно создать результирующий массив и положить в него первый интервал, а потом сранивать результирующий массив с текущим интервалом и при слиянии тоже брать минимумы и максимумы от результирующего массива и текущего интервала.

3) Неправильный ответ
Входные данные: intervals = [1,4][0,2][3,5];
Ожидаемый результат: [0,5]
Результат исполнения: [0,4][3,5]