package main

import (
	"fmt"
	"math"
)

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

// Функция для получения n-го треугольного числа
func triangularNumber(n int) int {
	return n * (n + 1) / 2
}

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
// и еще проверяем нет ли нулей в массиве
func isIncreasingByAtLeastOne(arr []int) bool {
	if arr[0] == 0 {
		return false
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] || arr[i] == 0 {
			return false
		}
	}
	return true
}

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

func main() {
	// Чтение количества записей
	var n int
	fmt.Scan(&n)

	sequence := make([]int, 0, n)
	// Чтение записей
	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		sequence = append(sequence, value)
	}

	// восстанавливаем пропуски и подставляем элементы треугольной последовательности
	filledSequence := fillTriangularSequence(sequence)

	// проверяем является ли восстановленная последовательность треугольной
	if isTriangularSequence(filledSequence) && countNonNegativeOnes(sequence) > 2 {
		fmt.Println("YES")
		fmt.Println(calculateDifferences(filledSequence))
		return
	}

	// арифметическая последовательность
	replacingArray, diff := findConsecutiveNonNegativeOnes(sequence)
	fillingMissingArray := fillMissingValues(replacingArray, diff)
	if len(fillingMissingArray) > 0 && compareArrays(sequence, fillingMissingArray) && fillingMissingArray[0] != 0 {
		fmt.Println("YES")
		fmt.Println(calculateDifferences(fillingMissingArray))
		return
	}

	replacingSequence := replaceNegativeOnes(sequence)

	if isIncreasingByAtLeastOne(replacingSequence) {
		fmt.Println("YES")
		fmt.Println(calculateDifferences(replacingSequence))
		return
	}

	fmt.Println("NO")
}
