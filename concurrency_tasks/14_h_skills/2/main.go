package main

import (
	"context"
	"log"
)

// что выведет программа?

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	ch1 <- 1
	ch2 <- 2

	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

loop:
	for {
		select {
		case a := <-ch1:
			log.Println(a)
		case b := <-ch2:
			log.Println(b)
		case <-ctx.Done():
			log.Println("complete")
			break loop
		}
	}
}
