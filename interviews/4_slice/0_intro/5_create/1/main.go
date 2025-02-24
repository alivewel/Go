package main

import "fmt"

// Создавать слайс через var, а потом использовать append можно:

func main() {
	var arr []int
	fmt.Println(append(arr, 2))
}
