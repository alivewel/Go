package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanNums() [][]int {
	var n int
	fmt.Scan(&n) // Считываем количество массивов

	arrays := make([][]int, n) // Создаем срез для массивов

	scanner := bufio.NewScanner(os.Stdin) // Создаем новый сканер для ввода

	for i := 0; i < n; i++ {
		scanner.Scan() // Считываем строку

		line := scanner.Text()          // Получаем текст строки
		numbers := strings.Fields(line) // Разбиваем строку на элементы

		for _, num := range numbers {
			value, err := strconv.Atoi(num) // Преобразуем строку в число
			if err == nil {
				arrays[i] = append(arrays[i], value) // Добавляем элемент в массив
			}
		}
	}
	return arrays
}

func main() {

	arrays := scanNums()
	// Выводим массивы для проверки
	for _, array := range arrays {
		fmt.Println(array)
	}
	fmt.Println(len(arrays))
}
