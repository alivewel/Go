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
func fillTriangularSequence(seq []int) []int {
	filledSeq := make([]int, len(seq))
	copy(filledSeq, seq)

	for i, num := range filledSeq {
		if num == -1 {
			// Найти n, для которого T_n = num
			for n := 1; ; n++ {
				if isTriangularNumber(filledSeq[i-1]) && triangularNumber(n) > filledSeq[i-1] {
					filledSeq[i] = triangularNumber(n)
					break
				}
			}
		}
	}

	return filledSeq
}

func main() {
	// sequence := []int{1, 3, 6, 10, 15}
	// sequence := []int{1, 3, 6}
	// sequence := []int{1, -1, 6, 10, 15}
	sequence := []int{1, 3, -1, -1, -1}
	// sequence := []int{2, 5, 8, 11, 14}
	// sequence := []int{2, 5, 8}
	// sequence := []int{3, 9, 27, 81}
	// sequence := []int{3, 9, 27}

	filledSequence := fillTriangularSequence(sequence)
	fmt.Println(filledSequence)

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
