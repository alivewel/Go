package main

import (
	"fmt"
)

func main() {
	bigString := "abcdefg"
	maxWindowSize := 4

	// // Проходим по строке справа налево
	for windowSize := 1; windowSize <= maxWindowSize; windowSize++ {
		for i := len(bigString) - windowSize; i >= 0; i-- {
			substring := bigString[i : i+windowSize]
			fmt.Println(substring)
		}
	}

	// После достижения максимального размера окна, продолжаем с левой стороны
	for windowSize := maxWindowSize; windowSize < len(bigString); windowSize++ {
		for i := 0; i <= len(bigString)-windowSize; i++ {
			substring := bigString[i : i+windowSize]
			fmt.Println(substring)
		}
	}
}
