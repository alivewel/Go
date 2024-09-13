package main

import "fmt"

// Функция сравнения массивов
func compareArrays(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != -1 && arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func main() {
	array1 := []int{1, -1, 6, 10, 15}
	array2 := []int{1, 3, 6, 10, 15}

	result := compareArrays(array1, array2)
	fmt.Println("Результат сравнения:", result)
}
