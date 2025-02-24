package main

import "fmt"

// Что выведет программа?

func modifySlice(slice []int) {
	for i := range slice {
		slice[i] *= 2
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	modifySlice(arr)

	fmt.Println(arr)
}
