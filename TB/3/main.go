package main

import (
	"fmt"
)

func containsAllChars(charsMap map[rune]int, substring string) bool {
	window := make(map[rune]int)

	// Заполняем карту символов для подстроки
	for _, char := range substring {
		window[char]++
	}

	// Проверяем, содержатся ли все символы из charsMap в window
	for char, count := range charsMap {
		if window[char] < count {
			return false
		}
	}
	return true
}

func main() {
	// bigString := "abcdefg"
	bigString := "abacba"
	containsChar := "abc"
	lenPass := 4

	// Создаем карту для символов, которые нужно найти
	charsMap := make(map[rune]int)
	for _, char := range containsChar {
		charsMap[char]++
	}

	end := len(bigString)
	start := end - len(containsChar)

	// Получаем подстроку с индексами от start до end
	// substring := bigString[start:end] // end+1, чтобы включить символ на позиции `end`


	// strings.Contains косячная, заменить ее
	// if containsAllChars(charsMap, substring) {
	// 	fmt.Printf("Подстрока '%s' найдена в строке.\n", substring)
	// 	return
	// }

	// Проходим по строке с конца
	for start >= 0 {
		substring := bigString[start:end]
		fmt.Println(substring)
		if len(substring) >= lenPass {
			end--
		}
		start--
		if containsAllChars(charsMap, substring) {
			fmt.Printf(substring)
			break
		}
	}
}
