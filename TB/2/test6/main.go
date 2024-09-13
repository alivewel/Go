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

func main() {
	array := []int{1, -1, 6, 10, 15}
	diff, result := findConsecutiveNonNegativeOnes(array)
	fmt.Println(result, diff) // Ожидаемый вывод: {-1, -1, 6, 10, -1}
}
