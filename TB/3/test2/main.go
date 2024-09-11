package main

import (
	"fmt"
)

func containsAllChars(window map[rune]int, charsMap map[rune]int) bool {
    for char, count := range charsMap {
        if window[char] < count {
            return false
        }
    }
    return true
}

func main() {
    bigString := "abcdefg"
    chars := "ce"
    length := 3

    // Создаем карту для символов, которые нужно найти
    charsMap := make(map[rune]int)
    for _, char := range chars {
        charsMap[char]++
    }

    // Создаем карту для текущего окна
    window := make(map[rune]int)

    // Инициализируем первое окно
    for i := 0; i < length && i < len(bigString); i++ {
        window[rune(bigString[i])]++
    }

    // Проверяем первое окно
    if containsAllChars(window, charsMap) {
        fmt.Printf("Найдена подстрока: %s\n", bigString[:length])
        return
    }

    // Сдвигаем окно по строке
    for i := length; i < len(bigString); i++ {
        // Добавляем новый символ в окно
        window[rune(bigString[i])]++
        // Удаляем старый символ из окна
        window[rune(bigString[i-length])]--
		fmt.Println(window)
        // Проверяем текущее окно
        if containsAllChars(window, charsMap) {
            fmt.Printf("Найдена подстрока: %s\n", bigString[i-length+1:i+1])
            return
        }
    }

    fmt.Println("Подстрока не найдена.")
}
