package main

import (
	"fmt"
	"strings"
)

func containsAllChars(charsMap map[rune]int, substring string) bool {
	window := make(map[rune]int)

	// Заполняем карту символов для подстроки
	for _, char := range substring {
		window[char]++
	}

	// Проверяем, содержатся ли все символы из charsMap в substring
	for char, count := range charsMap {
		if window[char] < count {
			return false
		}
	}
	return true
}

func main() {
	// Чтение количества записей
	var bigString, containsChar string
	var lenPass int
	fmt.Scan(&bigString)
	fmt.Scan(&containsChar)
	fmt.Scan(&lenPass)

	// Создаем карту для символов, которые нужно найти
	charsMap := make(map[rune]int)
	for _, char := range containsChar {
		charsMap[char]++
	}

	end := len(bigString)
	start := end - len(containsChar)
	success := false
	// Проходим по строке с конца
	for start >= 0 {
		if !strings.Contains(containsChar, bigString[start:start+1]) {
			end = start
			start = end - len(containsChar)
			if start < 0 {
				break
			}
		}

		substring := bigString[start:end]
		if len(substring) >= lenPass {
			end--
		}
		start--
		if containsAllChars(charsMap, substring) {
			fmt.Printf(substring)
			fmt.Println()
			success = true
			break
		}
	}
	if !success {
		fmt.Println("-1")
	}

}
