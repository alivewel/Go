package main

import (
	"fmt"
	"math"
	"sort"
)

type Element struct {
	Idx int
	Val int
} 

// nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
// elements = [{Idx: 1, Val: 0}, {Idx: 0, Val: 4}, {Idx: 2, Val: 5},
//             {Idx: 1, Val: 9}, {Idx: 0, Val: 10}, {Idx: 1, Val: 12},
//             {Idx: 0, Val: 15}, {Idx: 2, Val: 18}, {Idx: 1, Val: 20}]

func smallestRange(nums [][]int) []int {
	elements := []Element{}
	// перезаписываем в один общий список
	for i, curNums := range nums {
		for _, num := range curNums { // for i, num := range curNums {
			elements = append(elements, Element{Idx: i, Val: num})
		}
	}
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Val < elements[j].Val
	})

	count := make(map[int]int, 0) // key - idx, value - count
	// res := []int{math.MinInt, math.MaxInt}
	res := []int{math.MinInt32, math.MaxInt32}
	l := 0
	// сначала двигаем правую границу плавающего окна
	for r := range elements {
		rightElem := elements[r]
		count[rightElem.Idx]++
		for len(count) == len(nums) { // && l < len(elements) 
			leftElem := elements[l]
			bestRange := res[1] - res[0]
			curRange := rightElem.Val - leftElem.Val
			fmt.Println(bestRange, curRange)
			if curRange < bestRange || (curRange == bestRange && leftElem.Val < res[0]) { // && 
				// fmt.Println("!")
				res = res[:0] // clean array
				// res = []int{leftElem.Val, rightElem.Val}
				res = append(res, leftElem.Val, rightElem.Val)

			}
			count[leftElem.Idx]--
			if count[leftElem.Idx] == 0 {
				delete(count, leftElem.Idx)
			}
			l++
		}
	}

    return res
}

func main() {
	nums := [][]int{
		{4, 10, 15, 24, 26},
		{0, 9, 12, 20},
		{5, 18, 22, 30},
	}

	result := smallestRange(nums)
	fmt.Println(result) // Ожидается: [20 24]
}
