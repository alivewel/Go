package main

import "fmt"

// [[1,2,3,4,5],
// [6,7,8,9,10],
// [11,12,13,14,15]]

// matrix = [1,3,5,7][10,11,16,20][23,30,34,60]; target = 3;
// Ожидаемый результат true
// Результат исполнения false

func search(matrix [][]int, target int) bool {
	totalSize := len(matrix) * len(matrix[0])
	l, r := 0, totalSize
	middle := 0
	for r - l > 1 {
		middle = (r + l) / 2 // middle := (r + l) / 2
		// r l 12 0 mid 6
		// r l 12 0 mid 6
		i := middle / len(matrix[0])
		j := middle % len(matrix[0])
		if (matrix[i][j] <= target) { // (matrix[i][j] >= target) // (matrix[i][j] > target)
			l = middle // l = middle
		} else {
			r = middle // r = middle
		}
	}
	if matrix[l / len(matrix[0])][l % len(matrix[0])] == target { // matrix[middle / len(matrix[0])][middle % len(matrix[0])] == target 
		return true
	}
	return false
}

func main() {
	matrix := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	target := 3;

	res := search(matrix, target)
	fmt.Println("res:", res)
}
