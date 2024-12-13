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

// в данном случае у нас 3 кейса для выполнения
// в select порядок выполнения не определен (недерменирован).
// В случае, когда будет несколько кейсов для выполнения кейс для выполнения, будет выбран случайно.

// результат выполнения будет разный:

// 2024/12/13 15:34:33 1
// 2024/12/13 15:34:33 2
// 2024/12/13 15:34:33 complete

// 2024/12/13 15:34:33 2
// 2024/12/13 15:34:33 1

// 2024/12/13 15:34:33 complete

// 2024/12/13 15:34:33 1
// 2024/12/13 15:34:33 complete

// 2024/12/13 15:34:33 2
// 2024/12/13 15:34:33 complete

// 2024/12/13 15:34:33 complete
