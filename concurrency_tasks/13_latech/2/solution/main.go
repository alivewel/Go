package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()

	// time.Sleep(time.Millisecond * 500)

	for i := range ch {
		fmt.Println(i)
	}
	// time.Sleep(time.Millisecond * 100)
}

// при запуске начального кода мы получаем
// panic: send on closed channel
// попытка записи с закрытый канал

// Как исправить?

// Есть правило: Кто пишет в канал, тот его и закрывает.
// мы добавляем закрытие каналу в горутину, которая пишет в канал
// при запуске main-горутины мы дойдем до цикла for range,
// который будет ждать, пока кто-то запишет в канал
// от time.Sleep при этом можно избавиться
// изначально мы в main-горутине закрывали канал, прежде чем писали в него.

// при этом если убрать закрытие канала, то получим:
// all goroutines are asleep - deadlock!
// это происходит потому что цикл for range продолжает ждать пока кто-то запишет данные в канал