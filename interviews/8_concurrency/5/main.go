package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Первая горутина отправляет сообщение в ch1
	go func() {
		ch1 <- "one"
	}()

	// Вторая горутина отправляет сообщение в ch2
	go func() {
		ch2 <- "two"
	}()

	// Используем select для чтения из каналов
	for i := 0; i < 2; i++ {
		select {
		case one := <-ch1:
			fmt.Println(one)
		case two := <-ch2:
			fmt.Println(two)
		}
	}
}