package main

import (
	"fmt"
)

// Функция для вычисления простых чисел с помощью решета Эратосфена
func sieveOfEratosthenes(n int) map[int]struct{} {
	primes := make(map[int]struct{})
	if n < 2 {
		return primes
	}

	// Создаем слайс для отметки простых чисел
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Реализация алгоритма решета Эратосфена
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	// Записываем простые числа в множество
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			primes[p] = struct{}{}
		}
	}

	return primes
}

func main() {
	n := 30
	primes := sieveOfEratosthenes(n)
	fmt.Printf("Простые числа до %d: %v\n", n, primes)
}
