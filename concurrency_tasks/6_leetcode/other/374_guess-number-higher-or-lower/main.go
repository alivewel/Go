package main

import "fmt"

// Предполагается, что API guess уже определен
func guess(num int) int {
	// Эта функция будет предоставлена LeetCode
	// Например, если загаданное число 6, то:
	if num < 6 {
		return 1
	} else if num > 6 {
		return -1
	}
	return 0
}

func guessNumber(n int) int {
	low, high := 1, n

	for low <= high {
		mid := (low + high) / 2
		result := guess(mid)
		fmt.Println(mid, low, high)
		if result == 0 {
			return mid // Угадали число
		} else if result == -1 {
			high = mid - 1 // Загаданное число меньше mid
		} else {
			low = mid + 1 // Загаданное число больше mid
		}
	}

	return -1 // На всякий случай, хотя это не должно произойти
}

func main() {
	n := 10 // Пример верхней границы диапазона
	result := guessNumber(n)
	fmt.Println("Загаданное число:", result)
}
