package main

import (
	"fmt"
	"sort"
)

// Сначала нужно отсортировать массив по первой точке
// нам нужно 2 указателя slow fast
// slow стоит на текущем указателе
// fast двигается по следующим и пытается найти пересечение,
// как только мы находим элемент, который не входит в данное пересечение
// увеличиваем counter
// ставим slow указатель на него и смотрим следующие интервалы

func countShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    }) // sort.Slice
    slow, fast := 0, 0
    counter := 0
    for slow < len(points) {
        for fast + 1 < len(points) && max(points[slow][0], points[fast+1][0]) <= min(points[slow][1], points[fast+1][1]) {
            fast++
        }
        counter++
        slow = fast + 1
        fast = fast + 1
    }
    return counter
}

func main() {
	points := [][]int{{1,5},{8,12},{0,3},{6,8},{7,8}}// Вывод: 3
	fmt.Println(countShots(points)) 
}
