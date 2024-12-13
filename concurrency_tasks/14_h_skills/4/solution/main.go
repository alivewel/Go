package main

import (
	"fmt"
	"time"
)

// Реализовать функцию, выполняющую батчинг значений из канала с.
// Для возврата использовать выходной канал chan []any
// * с - входной канал значений. Из этих значений должны формироваться батчи
// * batchSize - размер канала
func doBatching(c chan any, batchSize int) chan []any {
	out := make(chan []any)
	go func() {
		arr := make([]any, batchSize)
		for value := range c {
			out <- value
		}
	}()
	
	return out
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
