package main

import "fmt"

func main() {
	var arr []int
	fmt.Println(arr == nil) // true

	arr = []int{}
	fmt.Println(arr == nil) // false
}
