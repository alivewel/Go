package main

import (
	"fmt"
	"sort"
)

func countShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// Сортируем интервалы по начальной точке
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	counter := 1
	currentStart, currentEnd := points[0][0], points[0][1]

	for i := 1; i < len(points); i++ {
		interval := points[i]
		// Если интервал пересекается с текущим пересечением
		if interval[0] <= currentEnd {
			currentStart = max(currentStart, interval[0])
			currentEnd = min(currentEnd, interval[1])
		} else {
			// Новый выстрел
			counter++
			currentStart, currentEnd = interval[0], interval[1]
		}
	}

	return counter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	points := [][]int{{3,9}, {7,12}, {3,8}, {6,8}, {9,10}, {2,9}, {0,9}, {3,9}, {0,6}, {2,8}}
	fmt.Println("res:", countShots(points)) // Ожидаемый вывод: 2
}
