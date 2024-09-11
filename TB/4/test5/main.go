package main

import (
	"fmt"
	"math"
)

// Функция для проверки, является ли число простым
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Функция для подсчета количества делителей числа
func countDivisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

// Функция для подсчета составных чисел с простым количеством делителей
func countCompositeWithPrimeDivisors(l, r int) int {
	count := 0
	for num := l; num <= r; num++ {
		if num > 1 {
			// Проверяем, является ли число составным
			isComposite := false
			for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
				if num%i == 0 {
					isComposite = true
					break
				}
			}
			if isComposite {
				divisorsCount := countDivisors(num)
				if isPrime(divisorsCount) {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	// l := 10
	// r := 30
	l := 1
	r := 9
	result := countCompositeWithPrimeDivisors(l, r)
	fmt.Printf("Количество составных чисел от %d до %d с простым количеством делителей: %d\n", l, r, result)
}
