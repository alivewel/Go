package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	fmt.Println("service1() started", time.Since(start))
	c <- "Hello from service 1"
}

func service2(c chan string) {
	fmt.Println("service2() started", time.Since(start))
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	time.Sleep(3 * time.Second)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	default:
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}

// без задержки сработает default
// с задержкой один из кейсов (рандомно)
// это происходит потому что данные не успевают записаться в канал
// потому что main горутина стартанула быстрее
// time.Sleep блокирует main горутину

// Оператор default является неблокируемым, но это еще не все,
// оператор default делает блок select всегда неблокируемым.
// Это означает, что операции отправки и чтение на любом канале
// (не имеет значения будет ли канал с буфером или без) всегда
// будут неблокируемыми.
