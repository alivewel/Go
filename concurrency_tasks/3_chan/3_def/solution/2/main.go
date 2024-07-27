package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// avoid deadlock
func main() {
	ch := make(chan int)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-ctx.Done():
			fmt.Println("Done")
		}
	}()
	ch <- 42
	cancel()
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}

// здесь мы добавили контекст
// и записали значение в канал

// по сути можно делать что-то одно
// отменять контекст
// или
// записывать значение в канал
