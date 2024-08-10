package main

import "fmt"

// Что выведет программа?

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("before", arr)
	arr = handle(arr)
	fmt.Println("after", arr)
}

func handle(arr []int) []int {
	arr = append(arr, 5)
	fmt.Println("append", arr)
	return arr
}
