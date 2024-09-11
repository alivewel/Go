package main

import (
	"fmt"
	"strings"
)

func main() {
	bigString := "abcdefg"
	// maxWindowSize := 4

	// Проходим по строке с конца
	for end := len(bigString); end > 0; end-- {
		for start := end - 1; start >= 0; start-- {
			// Получаем текущую подстроку
			substring := bigString[start:end]
			fmt.Println(substring)
			// Проверяем наличие подстроки в строке
			if strings.Contains(bigString, substring) {
				// fmt.Printf("Подстрока '%s' найдена в строке.\n", substring)
			}
		}
	}
}
