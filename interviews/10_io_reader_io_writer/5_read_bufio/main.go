package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Открываем файл для чтения
	f, err := os.Open("data/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // Закрываем файл в конце работы функции

	// Создаем новый буферизованный ридер
	reader := bufio.NewReader(f)

	// Читаем файл построчно
	for {
		line, err := reader.ReadString('\n') // Читаем строку до символа новой строки
		if err != nil {
			if err.Error() == "EOF" {
				break // Конец файла
			}
			log.Fatal(err)
		}
		fmt.Print(line) // Печатаем строку
	}
}
