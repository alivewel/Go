package main

import (
	"fmt"
)

func maxArea(height []int) int {
    l, r := 0, len(height) - 1
	maxArea := 0
	for l < r { // l <= r с таким условием тоже работает
		minHeight := height[l]
		if height[l] > height[r] {
			minHeight = height[r]
		}
		curArea := (r - l) * minHeight
		if curArea > maxArea {
			maxArea = curArea
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
