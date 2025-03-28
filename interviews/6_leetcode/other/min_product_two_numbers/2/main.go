package main

import (
	// "fmt"
	"math"
	// "sort"
	// "errors"
)

// 1 Задача
// Дается массив числе, нужно найти минимально возможное произведение 2 чисел

func MinProduct(nums []int) (int) {
	min1, min2 := math.MaxInt, math.MaxInt
	max1, max2 := math.MinInt, math.MinInt

	for _, num := range nums {
		// Обновляем два наименьших
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}

		// Обновляем два наибольших
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}

	// Вычисляем три возможных произведения
	prod1 := min1 * min2
	prod2 := max1 * max2
	prod3 := min1 * max1

	// Находим минимум
	result := prod1
	if prod2 < result {
		result = prod2
	}
	if prod3 < result {
		result = prod3
	}

	return result
}
