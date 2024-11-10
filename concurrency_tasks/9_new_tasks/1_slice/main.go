package main

import "fmt"

// Что выведет программа?

func main() {
	a := []int{1, 2, 3}

	b := a[:2]
	b = append(b, 4)

	fmt.Println(a)
	fmt.Println(b)
}
