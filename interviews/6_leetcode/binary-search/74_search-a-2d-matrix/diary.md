1) 

matrix = [1,3,5,7][10,11,16,20][23,30,34,60]; target = 3;
Ожидаемый результат true
Результат исполнения false

```
func search(matrix [][]int, target int) bool {
	totalSize := len(matrix) * len(matrix[0])
	l, r := 0, totalSize
	middle := 0
	for r - l > 1 {
		middle := (r + l) / 2
		i := middle / len(matrix[0])
		j := middle % len(matrix[0])
		if (matrix[i][j] >= target) {
			l = middle
		} else {
			r = middle
		}
	}
	if matrix[middle / len(matrix[0])][middle % len(matrix[0])] == target {
		return true
	}
	return false
}
```

2) 	
Имеется:
if matrix[middle / len(matrix[0])][middle % len(matrix[0])] == target 
Должно быть:
if matrix[l / len(matrix[0])][l % len(matrix[0])] == target
