package main

import "fmt"

// Что выведет программа?

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("before", arr)
	handle(arr[:1])
	fmt.Println("after", arr)
}

func handle(arr []int) {
	arr = append(arr, 5)
	fmt.Println("append", arr)
}

// Вывод программы:
// before [1 2 3 4]
// append [1 5]
// before [1 5 3 4]
