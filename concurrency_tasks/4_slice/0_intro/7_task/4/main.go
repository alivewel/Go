package main

import "fmt"

// Что если мы хотим изменить внутренний массив и не хотим менять внешний?

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
