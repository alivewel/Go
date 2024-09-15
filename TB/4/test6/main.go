package main

import (
	"fmt"
	"math"
)

// Функция для генерации простых чисел
func generatePrimes(n int) []int64 {
	primes := []int64{}
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}

	for i := 2; i <= n; i++ {
		if sieve[i] {
			primes = append(primes, int64(i))
			for j := i * i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}

	return primes
}

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	counter := 0
	primes := generatePrimes(int(math.Sqrt(float64(r))) + 1)

	for _, primePow := range primes {
		if primePow == 2 {
			continue
		}
		for _, primeToPow := range primes {
			res := math.Pow(float64(primeToPow), float64(primePow-1))
			if res >= float64(l) && res <= float64(r) {
				counter++
			} else if res > float64(r) {
				break
			}
		}
	}

	fmt.Printf("%d\n", counter)
}
