package main

import (
	"bufio"
	"os"
)

func main() {
	// Открываем файл для чтения и записи, создаем его, если он не существует
	f, err := os.OpenFile("data/file.txt", os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close() // no error handling

	// Создаем новый буферизованный писатель
	w := bufio.NewWriter(f)

	// Пишем данные в буфер
	_, err = w.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}
	// Сбрасываем буфер, чтобы записать данные в файл
	if err = w.Flush(); err != nil {
		panic(err)
	}
}
