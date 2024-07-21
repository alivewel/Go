package main

import (
	"fmt"
	"sync"
)

// merge two channels
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

func merge[T any](chans ...chan T) chan T {
	ch := make(chan T) // создаем выходной канал
	wg := sync.WaitGroup{}
	wg.Add(len(chans)) // устанавливаем количество горутин для ожидания

	for _, in := range chans { // проходимся по всем каналам
		go func(in chan T) { // для каждого канала запускаем отдельную горутину
			defer wg.Done()
			for val := range in { // учитываем буф. и небуф. каналы
				ch <- val
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(ch) // не забываем закрывать канал
	}()

	return ch
}

// можно обойтись без отдельной горутины для вызова
// wg.Wait() и закрытия канала на строке 40.
// В текущем подходе это делается в отдельной горутине,
// чтобы функция merge могла сразу вернуть объединенный канал.
// Однако, если вы готовы дождаться завершения всех горутин
// в merge перед возвратом канала, можно сделать это без
// дополнительной горутины.
