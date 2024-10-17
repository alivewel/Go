package main

import (
	"fmt"
)

// какая длина и емкость будет в первом и втором массиве
func main() {
	a := make([]int32, 0)
	a = append(a, []int32{1,2,3,4,5,6}...)
	a = append(a, 4)

	fmt.Println(len(a), cap(a))

	b := make([]int64, 0)
	b = append(b, []int64{1,2,3,4,5,6}...)
	b = append(b, 4)

	fmt.Println(len(b), cap(b))
}
