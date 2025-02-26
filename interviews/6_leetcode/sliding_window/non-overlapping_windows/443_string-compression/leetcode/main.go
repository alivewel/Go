package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	writeIdx := 0  // Индекс записи в изменяемом массиве
	readIdx := 0   // Индекс чтения

	for readIdx < len(chars) {
		char := chars[readIdx] // Текущий символ
		start := readIdx       // Запоминаем начальный индекс группы

		// Продвигаемся, пока символы совпадают
		for readIdx < len(chars) && chars[readIdx] == char {
			readIdx++
		}

		// Записываем сам символ
		chars[writeIdx] = char
		writeIdx++

		// Записываем количество повторений, если их больше 1
		count := readIdx - start
		if count > 1 {
			for _, c := range strconv.Itoa(count) {
				chars[writeIdx] = byte(c)
				writeIdx++
			}
		}
	}

	return writeIdx
}

func main() {
	input := []byte("aaabbcdddde")
	newLength := compress(input)
	fmt.Println(string(input[:newLength])) // Выведет: "a3b2cd4e"
	fmt.Println(newLength)                 // Выведет: 7
}
