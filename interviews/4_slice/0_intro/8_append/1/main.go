package main

import "fmt"

// Что выведет программа?

func main() {
	list := make([]int, 4, 4)

	list2 := append(list, 1)

	list[0] = 5
	list2[0] = 9

	fmt.Println(list)
	fmt.Println(list2)
}
