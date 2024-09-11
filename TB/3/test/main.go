package main

import (
	"fmt"
	"strings"
)

// Функция для проверки наличия всех символов в подстроке
func containsAllChars(substring string, containsChar string) bool {
	for _, char := range containsChar {
		if !strings.ContainsRune(substring, char) {
			return false
		}
	}
	return true
}

// Основная функция
func main() {
	// var lastCharInput string
	// var containsChar string
	// var lenPass int

	// // Ввод данных
	// fmt.Println("Введите большую строку:")
	// fmt.Scanln(&lastCharInput)
	// fmt.Println("Введите строку символов:")
	// fmt.Scanln(&containsChar)
	// fmt.Println("Введите количество символов для подстроки:")
	// fmt.Scanln(&lenPass)

	lastCharInput := "abacba"
	containsChar := "abc"
	lenPass := 4

	// Проверка длины строки
	if lenPass > len(lastCharInput) {
		fmt.Println(-1)
		return
	}

	// Проход по строке с конца
	for i := len(lastCharInput) - lenPass; i >= 0; i-- {
		substring := lastCharInput[i : i+lenPass]
		if containsAllChars(substring, containsChar) {
			fmt.Println(substring) // подстрока найдена
			return
		}
	}
	fmt.Println(-1) // подстрока не найдена
}

// 1. "Среди всех с одинаковым с ним началом - самый длинный" - что это значит?
// 2. Максимальная длина пароля - не означает, что пароль обязательно должен быть этой длины.
// Пароль может быть короче. При этом нам нужно взять самый длинный пароль.
// По сути можно сначала пройтись по всей строке и искать подстроку с длиной максимальной длины пароля.

// Выведите возможный пароль от ноутбука, удовлетворяющий указанным условиям.
// Если вариантов пароля несколько, выберете тот, который начинается с последовательности из первой строки правее
// (позже) других, а среди всех с ним началом - самый длинный
// "Среди всех с одинаковым с ним началом - самый длинный" - что это значит?

// Смущает формулировка - "Среди всех с одинаковым с ним началом - самый длинный"
// Смущает формулировка - максимальная длина пароля.
// Я использую подстроку длиной максимальной длины пароля.
// Кажется, что если в большей подстроке есть пароль, то нет смысла искать в меньшей подстроке.
