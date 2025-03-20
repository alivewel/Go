1) Слишком рано выхожу из цикла

Неправильный ответ
Входные данные intervals1 = [8,15]; intervals2 = [2,6][8,10][12,20];
Ожидаемый результат [8,10][12,15]
Результат исполнения [8,10]

```
func intersect(firstList [][]int, secondList [][]int) [][]int {
	p1, p2 := 0, 0
	res := make([][]int, 0)
	for p1 < len(firstList) && p2 < len(secondList) {
		if max(firstList[p1][0], secondList[p2][0]) <= min(firstList[p1][1], secondList[p2][1]) {
			intersect := []int{max(firstList[p1][0], secondList[p2][0]), min(firstList[p1][1], secondList[p2][1])}
			res = append(res, intersect)
		}
		if firstList[p1][0] > secondList[p2][0] {
			p2++
		} else {
			p1++
		}
	} 
	return res 
}
```

Важно сравнивать именно концы интервалов, а не начало
Имеется:
if firstList[p1][0] > secondList[p2][0]
Должно быть:
if firstList[p1][1] > secondList[p2][1]
