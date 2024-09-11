package main

import (
	"fmt"
)

// Функция для разложения числа на простые множители
func primeFactorsUnique(n int) int {
	var factors []int

	// Проверяем делимость на 2
	if n%2 == 0 {
		factors = append(factors, 2)
		for n%2 == 0 {
			n /= 2
		}
	}

	// Проверяем делимость на нечетные числа
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			factors = append(factors, i)
			for n%i == 0 {
				n /= i
			}
		}
	}

	// Если n все еще больше 2, то оно само является простым числом
	if n > 2 {
		factors = append(factors, n)
	}

	// Перемножаем простые множители друг на друга и получаем количество делителей числа
	result := 1
	for _, value := range factors {
		result *= value
	}
	return result

	// return factors
}

func main() {
	number := 720
	fmt.Printf("Простые множители числа %d: %v\n", number, primeFactorsUnique(number))
}
