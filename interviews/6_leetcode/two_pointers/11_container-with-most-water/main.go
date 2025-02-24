package main

import (
	"fmt"
)

func maxArea(height []int) int {
    l, r := 0, len(height) - 1
	maxArea := 0
	for l < r { // l <= r с таким условием тоже работает
		curArea := (r - l) * min(height[l], height[r])
		maxArea = max(curArea, maxArea)
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

// оказывается можно не реализовывать функции max и min, они сами подхватываются из стандартной библиотеки
