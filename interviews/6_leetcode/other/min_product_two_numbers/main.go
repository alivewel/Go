package main

import (
	// "fmt"
	// "math"
	"sort"
)

// 1 Задача
// Дается массив числе, нужно найти минимально возможное произведение 2 чисел

func MinProduct(nums []int) int {
	if len(nums) < 2 {
		panic("Массив должен содержать минимум два элемента")
	}

	sort.Ints(nums)
	n := len(nums)

	// Минимальное произведение — либо два наименьших, либо два наибольших
	product1 := nums[0] * nums[1]
	product2 := nums[n-1] * nums[n-2]
	product3 := nums[0] * nums[n-1]

	// Найдём минимум из трёх
	minProduct := product1
	if product2 < minProduct {
		minProduct = product2
	}
	if product3 < minProduct {
		minProduct = product3
	}
	// fmt.Println(product1, product2, product3)
	return minProduct

	// product1 := nums[0] * nums[1]
	// product3 := nums[0] * nums[n-1]
	// if product1 < product3 {
	// 	return product1
	// }
	// fmt.Println(product1, product3)

	// return product3
}

// func MinProduct(nums []int) int {
// 	sort.Ints(nums)
//     f := nums[0]
//     l := nums[len(nums)-1]

//     if (f > -1 && l > -1) || (f < 0 && l < 0) {
//         return nums[0] * nums[1]
//     }
// 	fmt.Println()
//     return nums[0] * nums[len(nums)-1]
// }

// func MinProduct(nums []int) int {
// 	sort.Ints(nums)

// 	f := nums[0]
// 	l := nums[len(nums)-1]

// 	if f > -1 {
// 		return nums[0] * nums[1]
// 	}

// 	if f < 0 && l < 0 {
// 		return nums[len(nums)-1] * nums[len(nums)-2]
// 	}

// 	return nums[0] * nums[len(nums)-1]
// }


