package main

import (
	"fmt"
)

// что выведет программа и почему?
func main() {
	a1 := make([]int, 0, 2)
	a1 = append(a1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(cap(a1), len(a1))
	a2 := append(a1, 6)
	a3 := append(a1, 7)

	fmt.Println(a1, a2, a3)
	fmt.Println(cap(a1))
}
