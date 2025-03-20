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

func merge(intervals [][]int) [][]int {
	// sort.Ints(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] // Сортируем по первому элементу (началу отрезка)
	})
	for i := range intervals {
		if i == len(intervals) - 1 {
			continue
		}
		if max ()
	}
	
	return intervals
}

func main() {
	segments := [][]int{{2,5},{0,6},{0,3},{9,11},{6,8}} // Вывод: [][]int{{0,8},{9,11}}
	fmt.Println(merge(segments))
}
