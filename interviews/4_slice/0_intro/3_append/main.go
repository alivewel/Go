package main

import "fmt"

func main() {
	var arr [8]int
	arr = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice1 := arr[4:7]
	slice1 = append(slice1, 12) // исходный слайс тоже изменится, arr[8] перезапишется
	fmt.Println(slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println(arr)
}
