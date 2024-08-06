package main

import "fmt"

// Создание слайс через make, с заранее указанной длиной и емкостью:

func main() {
	arr := make([]int, 5, 10)
	fmt.Println(len(arr), cap(arr))
}
