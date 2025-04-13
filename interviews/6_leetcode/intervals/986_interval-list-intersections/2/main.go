package main

import (
	"fmt"
)

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    p1, p2 := 0, 0
    res := make([][]int, 0)
    for p1 < len(firstList) && p2 < len(secondList) {
        if max(firstList[p1][0], secondList[p2][0]) <= min(firstList[p1][1], secondList[p2][1]) {
            intersection := []int{max(firstList[p1][0], secondList[p2][0]), min(firstList[p1][1], secondList[p2][1])}
            res  = append(res, intersection)
        }
        if firstList[p1][1] > secondList[p2][1] {
            p2++
        } else {
            p1++
        }
    }
    return res
}

func main() {
	firstList := [][]int{{0,2},{5,10},{13,23},{24,25}}
	secondList := [][]int{{1,5},{8,12},{15,24},{25,26}}
	fmt.Println(intervalIntersection(firstList, secondList)) // [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
}