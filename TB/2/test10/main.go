package main

import "fmt"

// Функция для нахождения двух подряд идущих чисел, которые не равны -1
func findConsecutiveNonNegativeOnes(arr []int) ([]int, int) {
	result := make([]int, len(arr))
	diff := 0
	for i := range result {
		result[i] = -1
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != -1 && arr[i+1] != -1 {
			result[i] = arr[i]
			result[i+1] = arr[i+1]
			if arr[i+1] > arr[i] {
				diff = arr[i+1] - arr[i]
			}
			break
		}
	}

	return result, diff
}

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
	sequence := []int{-1, -1, 7, 9, 11}
	replacingArray, diff := findConsecutiveNonNegativeOnes(sequence)

	// fmt.Println(replacingArray, diff) // Ожидаемый вывод: {-1, -1, 6, 10, -1}

	fillingMissingArray := fillMissingValues(replacingArray, diff)
	if len(fillingMissingArray) > 0 && compareArrays(sequence, fillingMissingArray) {
		fmt.Println(fillingMissingArray) // Ожидаемый вывод: [-2, 2, 6, 10, 14]
	}

}
