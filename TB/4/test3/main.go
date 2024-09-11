package main

import (
	"fmt"
)

// Функция для вычисления простых чисел с помощью решета Эратосфена
func sieveOfEratosthenes(max int) []int {
	if max < 2 {
		return []int{}
	}

	// Создаем слайс для отметки простых чисел
	isPrime := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		isPrime[i] = true
	}

	// Реализация алгоритма решета Эратосфена
	for p := 2; p*p <= max; p++ {
		if isPrime[p] {
			for i := p * p; i <= max; i += p {
				isPrime[i] = false
			}
		}
	}

	// Собираем простые числа в слайс
	var primes []int
	for p := 2; p <= max; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func main() {
	max := 30
	primes := sieveOfEratosthenes(max)
	fmt.Printf("Простые числа до %d: %v\n", max, primes)
}
