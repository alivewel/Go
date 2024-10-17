package main

import "fmt"

func main() {
	var arr []int
	fmt.Println(arr == nil) // true

	arr = []int{}
	fmt.Println(arr == nil) // false

	if len(arr) == 0 {
		fmt.Println("arr is empty")
	}
}
