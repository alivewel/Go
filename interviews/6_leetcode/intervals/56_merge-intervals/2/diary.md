1) Не сразу догадался, как отсортировать двумерный массив. Для этого нужно использовать sort.Slice

```
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0] // Сортируем по первому элементу (началу отрезка)
})
```

2) Условие сравнения <= должно быть, а не <

Ошибка из-за невнимательности
Имеется:
if max(lastElemRes[0], interval[0]) < min(lastElemRes[1], interval[1]) 
Должно быть:
if max(lastElemRes[0], interval[0]) <= min(lastElemRes[1], interval[1])
