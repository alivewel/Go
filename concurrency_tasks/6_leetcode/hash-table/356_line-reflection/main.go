package main

import "fmt"

func isReflected(points [][]int) bool {
	minX, maxX := points[0][0], points[0][0]
	pointsMap := make(map[[2]int]struct{}, len(points))
	for _, point := range points {
		if point[0] < minX {
			minX = point[0]
		}
		if point[0] > maxX {
			maxX = point[0]
		}
		pointsMap[[2]int{point[0], point[1]}] = struct{}{}
	}
	avgLine := maxX - minX
	for _, point := range points {
		pointX := 0
		if point[0] > avgLine {
			pointX = avgLine - (point[0] - avgLine)
		} else {
			pointX = avgLine + (avgLine - point[0])
		}
		if _, found := pointsMap[[2]int{pointX, point[1]}]; !found {
			return false
		}
	}
	return true
}

func main() {
	points := [][]int{
		{1, 2},
		{3, 1},
		{4, 2},
		{2, 1},
		{2, 1},
	}
	fmt.Println(isReflected(points))
}
