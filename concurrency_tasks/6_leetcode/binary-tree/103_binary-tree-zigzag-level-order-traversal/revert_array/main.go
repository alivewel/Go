package main

import "fmt"

// перевернуть массив

func main() {
	nodes := []int{1, 2, 3, 4, 5}
	// nodes := []int{1, 2, 3, 4}
	// nodes := []int{1}
	// nodes := []int{}

	length := len(nodes)
	if length > 0 {
		for i := 0; i < length/2; i++ {
			temp := nodes[i]
			nodes[i] = nodes[length-i-1]
			nodes[length-i-1] = temp
		}
	}
	fmt.Println(nodes)
}
