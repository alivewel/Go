package main

import "sort"

type Element struct {
	Idx int
	Val int
} 

// nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]

func smallestRange(nums [][]int) []int {
	elements := []Element{}
	for i, curNums := range nums {
		for _, num := range curNums {
			elements = append(elements, Element{Idx: i, Val: n})
		}
	}

	sort.Slices(elements, func(i, j int) bool {
		return elements[i].Val < elements[j].Val
	})

	left := 0
	count := make(map[int]int, 0)
	res := make([]int, 0)
	bestRange := 0
	for right := range elements {
		rightElem := elements[right]
		count[rightElem.Idx]++
		for len(nums) == len(count) {
			leftElem := elements[left]
			curRange := rightElem.Val - leftElem.Val
			if bestRange > curRange {
				bestRange = curRange
				res = res[:0]
				res = append(res, leftElem.Val, rightElem.Val)
			}
			count[leftElem.Idx]--
			if count[leftElem.Idx] == 0 {
				delete(count, leftElem.Idx)
			}
			left++
		}
	}
    return res
}
