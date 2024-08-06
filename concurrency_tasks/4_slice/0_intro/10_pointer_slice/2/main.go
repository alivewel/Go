package main

import "fmt"

// Что выведет программа?

func modifySlicePointer(slice *[]int) {
	*slice = append(*slice, 6, 7, 8)
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	modifySlicePointer(&arr)

	fmt.Println(arr)
}
