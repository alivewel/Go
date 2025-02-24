package main

import (
	"fmt"
)

// какая длина и емкость будет в первом и втором массиве
func main() {
	a := make([]int32, 0) // len == 0, cap == 0
	a = append(a, []int32{1,2,3}...) // len == 3, cap == 3
	a = append(a, 4)  // len == 4, cap == 4

	fmt.Println(len(a), cap(a))

	b := make([]int64, 0) // len == 0, cap == 0
	b = append(b, []int64{1,2,3}...) // len == 3, cap == 3
	b = append(b, 4) // len == 4, cap == 6

	fmt.Println(len(b), cap(b))
}
