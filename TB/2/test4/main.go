package main

import "fmt"

// Функция для вычисления разниц между элементами массива
func calculateDifferences(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	differences := make([]int, len(arr))
	differences[0] = arr[0] // Первый элемент остается тем же

	for i := 1; i < len(arr); i++ {
		differences[i] = arr[i] - arr[i-1]
	}

	return differences
}

func main() {
	sequence := []int{1, 3, 6, 10, 15}
	differences := calculateDifferences(sequence)
	fmt.Println(differences) // Вывод: {1, 2, 3, 4, 5}
}
