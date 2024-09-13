package main

import "fmt"

// Функция для замены -1 на вычисленные значения, используя разницу
func fillMissingValues(arr []int, diff int) []int {

	// Проход слева направо
	for i := 0; i < len(arr); i++ {
		if arr[i] == -1 && i > 0 && arr[i-1] != -1 {
			arr[i] = arr[i-1] + diff
		}
	}

	// Проход справа налево
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == -1 && i < len(arr)-1 && arr[i+1] != -1 {
			arr[i] = arr[i+1] - diff
			if arr[i] < 1 {
				return nil
			}
		}
	}

	return arr
}

func main() {
	array := []int{-1, -1, 6, 10, -1}
	diff := 4
	result := fillMissingValues(array, diff)
	if len(result) > 0 {
		fmt.Println(result) // Ожидаемый вывод: [-2, 2, 6, 10, 14]
	}
}
