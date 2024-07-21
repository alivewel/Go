package main

import (
	"fmt"
	// "sync"
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

	ch3 := syncMerge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

func syncMerge[T any](chans ...chan T) chan T {
	totalLen := 0              // общая длина всех каналов
	for _, ch := range chans { // проходимся по всем каналам и суммируем длину
		totalLen += len(ch) // учитываем буф. и небуф. каналы
	}
	out := make(chan T, totalLen) // новый объединенный канал
	for _, ch := range chans {    // проходимся во второй раз
		for val := range ch { // учитываем буф. и небуф. каналы
			out <- val
		}
	}
	close(out) // не забываем закрывать канал
	return out
}

// написать функцию для синхронного мерджа каналов из видео
// подробно описать ее и понять работу
// понять функцию для мерджа асинхронную и написать ее самому

// func mergeChannels(channels ...<-chan int) <-chan int {
// 	var wg sync.WaitGroup
// 	out := make(chan int)
// 	output := func(c <-chan int) {
// 		defer wg.Done()
// 		for val := range c {
// 			out <- val
// 		}
// 	}
// 	wg.Add(len(channels))
// 	for _, c := range channels {
// 		go output(c)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 		}()
// 		return out
// 	}
