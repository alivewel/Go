package main

import "fmt"

func isSymmetric(points [][]int) bool {
    minX, maxX := 0, 0
    setPoints := make(map[[2]int]struct{})
    for i, point := range points {
        if i == 0 {
            minX, maxX = point[0], point[0]
        }
        if minX > point[0] {
            minX = point[0]
        }
        if maxX < point[0] {
            maxX = point[0]
        }
        setPoints[[2]int{point[0], point[1]}] = struct{}{}
    }
    for _, point := range points {
                   //3 ==  1 + 4 - 2
                   //4 ==  1 + 4 - 1
        simPoint := minX + maxX - point[0] // minX + maxX - x
        if _, found := setPoints[[2]int{simPoint, point[1]}]; !found {
            return false
        }
    }
    return true
}
//   x y
// [[1,2],[3,1],[4,2],[2,1],[2,1]]

// изначально прочитал readme, чтобы вспомнить алгоритм решения
// не смог вспомнить формулу для нахождения симметричной точки

func main() {
	points := [][]int{{1,2},{3,1},{4,2},{2,1},{2,1}}
	fmt.Println(isSymmetric(points))
}
