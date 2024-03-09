package main

import (
	"fmt"
	"time"
)

// print square of range 0...20 in random order
func main() {
	counter := 20
	for i := 0; i < counter; i++ {
		go func() {
			fmt.Println(i * i)
		}()
	}

	time.Sleep(time.Second)
}

// 1. Здесь не хватает примитивов синхронизации - Wait groups.
// 2. Необходимо сделать захват переменной в анонимной функции.
// Или внутри цикла создать новую переменную: i := i 
// Исправленный вариант лежит в папке solution