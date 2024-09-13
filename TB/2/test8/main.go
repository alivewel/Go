package main

import "fmt"

// Функция для проверки, возрастают ли числа в массиве хотя бы на 1
func isIncreasingByAtLeastOne(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	array := []int{1, 2, 3, 4, 5}
	result := isIncreasingByAtLeastOne(array)
	fmt.Println(result) // Ожидаемый вывод: true

	array2 := []int{1, 2, 2, 4, 5}
	result2 := isIncreasingByAtLeastOne(array2)
	fmt.Println(result2) // Ожидаемый вывод: false
}
