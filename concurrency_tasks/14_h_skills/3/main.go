package main

import (
	"fmt"
)

// что выведет программа?

func main() {
	slice2 :=[]int{1, 2, 3, 4, 5}
	fmt.Println(len(slice2), cap(slice2))

	slice3 := slice2[1:]
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice4 := slice2[:3]
	fmt.Println(slice4, len(slice4), cap(slice4))

	slice4 = append(slice4, 1, 2)
	fmt.Println(slice4)
	fmt.Println(slice2)
}
