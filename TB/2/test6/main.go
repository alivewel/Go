package main

import "fmt"

// Функция для нахождения двух подряд идущих чисел, которые не равны -1
func findConsecutiveNonNegativeOnes(arr []int) []int {
	result := make([]int, len(arr))
	for i := range result {
		result[i] = -1
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != -1 && arr[i+1] != -1 {
			result[i] = arr[i]
			result[i+1] = arr[i+1]
			break
		}
	}

	return result
}

func main() {
	array := []int{1, -1, 6, 10, 15}
	result := findConsecutiveNonNegativeOnes(array)
	fmt.Println(result) // Ожидаемый вывод: {-1, -1, 6, 10, -1}
}
