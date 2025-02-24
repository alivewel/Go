package main

import "fmt"

// Способы создания массива

func main() {
	var arr1 [4]int
	arr1 = [4]int{1, 2, 3, 4}

	arr2 := [4]int{5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
