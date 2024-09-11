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
	// bigString := "abcdefg"
	// bigString := "abacba"
	// bigString := "dacbacaba"
	// bigString := "dabacaba"
	// bigString := "abacaba"
	// containsChar := "dac"
	// containsChar := "abc"
	// lenPass := 4

	// bigString := "aacbabc"
	// containsChar := "abc"
	// lenPass := 4

	// bigString := "xyzabcabcxyz"
	// containsChar := "abc"
	// lenPass := 4

	bigString := "abababab"
	containsChar := "abc"
	lenPass := 4

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
			// fmt.Println("отсуп", bigString[start:end], start, end)
		}

		substring := bigString[start:end]
		// fmt.Println(substring)
		// fmt.Println(bigString[start:start+1])
		if len(substring) >= lenPass {
			end--
		}
		start--
		if containsAllChars(charsMap, substring) {
			fmt.Printf(substring)
			success = true
			break
		}
	}
	if !success {
		fmt.Println("-1")
	}

}

// В случае
// bigString := "dcabacaba"
// containsChar := "dac"
// Выход
// dcab
// возможно должно выходить только dca
// буквы b точно не должно быть
