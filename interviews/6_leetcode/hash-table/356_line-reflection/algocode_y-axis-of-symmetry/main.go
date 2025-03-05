package main

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
}
//   x y
// [[1,2],[3,1],[4,2],[2,1],[2,1]]

func main() {
	points := [][]int{{1,2},{3,1},{4,2},{2,1},{2,1}}
	fmt.Println(isSymmetric(points))
}
