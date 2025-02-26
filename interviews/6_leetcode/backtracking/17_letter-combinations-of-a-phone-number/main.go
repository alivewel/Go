package main

import "fmt"

// Функция для получения всех комбинаций букв
func letterCombinations(digits string) []string {
    // Карта соответствия цифр и букв
    digitToLetters := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    // Если строка пуста, вернуть пустой результат
    if len(digits) == 0 {
        return []string{}
    }

    // Инициализация очереди
    queue := []string{""}

    // Итерация по каждой цифре        // digits := "23"
    for i := 0; i < len(digits); i++ { // до 2
        letters := digitToLetters[digits[i]] // digits[0] == 2, digits[1] == 3
        var newQueue []string // letters == "abc", letters == "def"

        // Удаление элементов из очереди и добавление новых комбинаций
        for _, combination := range queue {
			fmt.Println(i, combination)
            for j := 0; j < len(letters); j++ {
				// fmt.Println(i, combination)
                newQueue = append(newQueue, combination+string(letters[j]))
            }
        }
		fmt.Println(i, " - ", newQueue)
        // Обновление очереди
        queue = newQueue
    }

    return queue
}

func main() {
    // Пример использования
    digits := "23"
    combinations := letterCombinations(digits)
    fmt.Println(combinations) // Вывод: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
}