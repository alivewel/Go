package main

import "fmt"

// создание слайса на основе массива

func main() {
	var arr [9]int
	arr = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := arr[4:7]
	fmt.Println(slice1, "len:", len(slice1), "cap:", cap(slice1))
}

