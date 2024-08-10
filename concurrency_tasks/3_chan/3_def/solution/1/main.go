package main

import (
	"fmt"
	"runtime"
	"time"
)

// avoid deadlock
func main() {
	ch := make(chan int)

	go func() {
		select {
		case val := <-ch:
			fmt.Println(val)
		default:
			fmt.Println("default")
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}

// решение - добавление default сценария

// добавление горутины, чтобы увидеть количество активных горутин
// при этом, если мы уберем default на 16 и 17 строке, у нас произойдет утечка горутины
// main завершит свою работу, а созданная нами горутина нет

// также можно решить задачу, используя контекст
