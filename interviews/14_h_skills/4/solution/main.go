package main

import (
	"fmt"
	"time"
)

// Реализовать функцию, выполняющую батчинг значений из канала с.
// Для возврата использовать выходной канал chan []any
// * с - входной канал значений. Из этих значений должны формироваться батчи
// * batchSize - размер канала
// func doBatching(c chan any, batchSize int) chan []any {
// 	out := make(chan []any)
// 	go func() {
// 		arr := make([]any, batchSize)
// 		for value := range c {
// 			out <- value
// 		}
// 	}()
	
// 	return out
// }

func doBatching(c chan any, batchSize int) chan []any {
	cOut := make(chan []any)

	go func() {
		defer close(cOut) // Закрываем выходной канал после завершения работы горутины
		batch := make([]any, 0, batchSize)

		for value := range c {
			batch = append(batch, value) // Добавляем значение в текущий батч
			if len(batch) == batchSize {  // Если батч достиг нужного размера
				cOut <- batch              // Отправляем батч в выходной канал
				batch = make([]any, 0, batchSize) // Сбрасываем батч
			}
		}

		// Отправляем оставшиеся элементы, если они есть
		if len(batch) > 0 {
			cOut <- batch
		}
	}()

	return cOut
}

func main() {
	c := make(chan any, 10)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	close(c)

	cOut := doBatching(c, 2)

	go func() {
		for value := range cOut {
			fmt.Println(value)
		}
	}()

	time.Sleep(time.Second)
}
