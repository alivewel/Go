package main

import (
	"fmt"
)

// написать функцию для возведения чисел 1 до 9 в квадрат через каналы

func squares(ch chan int) {
	for i := 0; i < 9; i++ {
		ch <- i * i
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go squares(ch)

	// periodic block/unblock of main goroutine until chanel closes
	for i := range ch {
		fmt.Println(i)
	}

}