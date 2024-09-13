package main

import "fmt"

// Функция для замены -1 на предыдущий элемент +1
func replaceNegativeOnes(arr []int) []int {
	if arr[0] == -1 {
		arr[0] = 1
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] == -1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr
}

func main() {
	array := []int{-1, -1, 6, 10, -1}
	result := replaceNegativeOnes(array)
	fmt.Println(result) // Ожидаемый вывод: [1, 2, 6, 10, 11]
}
