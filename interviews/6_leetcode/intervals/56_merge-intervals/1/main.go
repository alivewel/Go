package main

import (
	"fmt"
	"sort"
)
// Отсортировать массив по первому числу в массиве.
// Проверить на пересечение
// Слить

//  a1 b1 a2 b2
// [[0 6] [0 3] [2 5] [6 8] [9 11]]
// условие пересечения 
// max(a1,a2) <= min(b1,b2)

// сортируем массив используя sort.Slice
// создаем результирующий массива и кладем в него первый интервал
// мы двигаемся по массиву с индекса [1]
// сравниваем последний элемент результирующего массива c текущим интервалом
	
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] // Сортируем по первому элементу (началу отрезка)
	})
	res := make([][]int, 0)
	res = append(res, []int{intervals[0][0], intervals[0][1]})
	for i := range intervals {
		if i == 0 {
			continue
		}
		if max(res[len(res)-1][0], intervals[i][0]) <= min(res[len(res)-1][1], intervals[i][1]) {
			if len(res) == 0 {
				row := make([]int, 2)
				res = append(res, row)
			}
			res[len(res)-1][0] = min(res[len(res)-1][0], intervals[i][0]) 
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])  
		} else {
			res = append(res, intervals[i])
		}
	}
	
	return res
}

func main() {
	// segments := [][]int{{2,5},{0,6},{0,3},{9,11},{6,8}} // Вывод: [][]int{{0,8},{9,11}}
	segments := [][]int{{1,4},{0,2},{3,5}} // Вывод: [0,5]
	fmt.Println(merge(segments)) // [[0 5] [2 5] [6 8]]
}
