package main

import "fmt"

// Что выведет программа?

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("before", arr)
	handle(arr[:1])
	fmt.Println("before", arr)
}

func handle(arr []int) {
	arr = append(arr, 5)
	fmt.Println("append", arr)
}
