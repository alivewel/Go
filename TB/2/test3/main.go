package main

import (
	"fmt"
	"math"
)

// Функция для проверки арифметической прогрессии
func isArithmetic(seq []int) bool {
	if len(seq) < 2 {
		return false
	}
	diff := seq[1] - seq[0]
	for i := 2; i < len(seq); i++ {
		if seq[i]-seq[i-1] != diff {
			return false
		}
	}
	return true
}

// Функция для проверки геометрической прогрессии
func isGeometric(seq []int) bool {
	if len(seq) < 2 || seq[0] == 0 {
		return false
	}
	ratio := float64(seq[1]) / float64(seq[0])
	for i := 2; i < len(seq); i++ {
		if float64(seq[i])/float64(seq[i-1]) != ratio {
			return false
		}
	}
	return true
}

// Функция для проверки, является ли число треугольным
func isTriangularNumber(x int) bool {
	// Решаем уравнение n(n + 1)/2 = x для n
	n := (math.Sqrt(8*float64(x)+1) - 1) / 2
	return n == float64(int(n))
}

// Функция для проверки, является ли последовательность треугольной
func isTriangularSequence(seq []int) bool {
	for _, num := range seq {
		if !isTriangularNumber(num) {
			return false
		}
	}
	return true
}

// Функция для проверки, является ли число треугольным
// func isTriangularNumber(x int) bool {
// 	n := (math.Sqrt(8*float64(x)+1) - 1) / 2
// 	return n == float64(int(n))
// }

// Функция для получения n-го треугольного числа
func triangularNumber(n int) int {
	return n * (n + 1) / 2
}

// Функция для заполнения недостающих чисел в треугольной последовательности
// func fillTriangularSequence(seq []int) []int {
// 	filledSeq := make([]int, len(seq))
// 	copy(filledSeq, seq)

// 	for i, num := range filledSeq {
// 		if num == -1 {
// 			// Найти n, для которого T_n = num
// 			for n := 1; ; n++ {
// 				if isTriangularNumber(filledSeq[i-1]) && triangularNumber(n) > filledSeq[i-1] {
// 					filledSeq[i] = triangularNumber(n)
// 					break
// 				}
// 			}
// 		}
// 	}

// 	return filledSeq
// }

// Функция для заполнения недостающих чисел в треугольной последовательности
func fillTriangularSequence(seq []int) []int {
	filledSeq := make([]int, len(seq))
	copy(filledSeq, seq)

	for i, num := range filledSeq {
		if num == -1 {
			// Используем индекс i для вычисления треугольного числа
			filledSeq[i] = triangularNumber(i + 1)
		}
	}

	return filledSeq
}

// Функция для подсчета количества элементов, не равных -1
func countNonNegativeOnes(arr []int) int {
	count := 0
	for _, value := range arr {
		if value != -1 {
			count++
		}
	}
	return count
}

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
	// sequence := []int{1, 3, 6, 10, 15}
	// sequence := []int{1, 3, 6}
	// sequence := []int{1, -1, 6, 10, 15}
	// sequence := []int{-1, -1, 6, 10, 15}
	// sequence := []int{1, 3, -1, -1, -1}
	// sequence := []int{1, 3, -1, -1, 15}
	// sequence := []int{2, 5, 8, 11, 14}
	// sequence := []int{2, 5, 8}
	// sequence := []int{3, 9, 27, 81}
	// sequence := []int{3, 9, 27}
	// sequence := []int{1, 3, -1}

	// sequence := []int{-1, -1, 6, 11, -1}
	sequence := []int{10, -1, 4}

	// восстанавливаем пропуски и подставляем элементы треугольной последовательности
	filledSequence := fillTriangularSequence(sequence)
	// fmt.Println(sequence)
	// fmt.Println(filledSequence)
	// compareArrays(sequence, filledSequence) похоже бесполезная
	// проверяем является ли восстановленная последовательность треугольной
	if isTriangularSequence(filledSequence) {
		fmt.Println("YES")
		fmt.Println(calculateDifferences(filledSequence))
		return
	}

	// функция для нахождения двух чисел подряд в массиве, которые не равны -1,
	// вернуть слайс с длиной исходного массива, все остальные заполнить числами -1

	// функция для проверки, возрастают ли числи

	replacingSequence := replaceNegativeOnes(sequence)

	if isIncreasingByAtLeastOne(replacingSequence) {
		fmt.Println("YES")
		fmt.Println(calculateDifferences(replacingSequence))
		return
	}

	fmt.Println("NO")

	if isArithmetic(sequence) {
		fmt.Println("Это арифметическая прогрессия")
	} else if isGeometric(sequence) {
		fmt.Println("Это геометрическая прогрессия")
	} else if isTriangularSequence(sequence) {
		fmt.Println("Последовательность является треугольной")
	} else {
		fmt.Println("Последовательность не является ни арифметической, ни геометрической прогрессией")
	}
}

// []int{1, 3, -1}
// при таком случае тяжело определить тип последовательности
// при таком случае тяжело точно определить следующий элемент
