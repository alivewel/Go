package main

import "fmt"

// Основная функция для получения всех комбинаций
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// Карта соответствия цифр и букв
	lettersDigits := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	// Вызов вынесенной функции backtrack
	backtrack(digits, lettersDigits, 0, "", &result)

	return result
}

// Вынесенная функция backtrack
func backtrack(digits string, lettersDigits map[byte]string, index int, combination string, result *[]string) {
	// Базовый случай: если длина комбинации равна длине digits
	if index == len(digits) {
		*result = append(*result, combination)
		return
	}

	// Получаем буквы, соответствующие текущей цифре
	letters := lettersDigits[digits[index]]

	// Рекурсивно перебираем каждую букву
	for _, letter := range letters {
		backtrack(digits, lettersDigits, index+1, combination+string(letter), result)
	}
}

func main() {
	// Пример использования
	digits := "23"
	fmt.Println(letterCombinations(digits)) // Вывод: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
}