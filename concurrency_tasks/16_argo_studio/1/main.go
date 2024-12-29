package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

// const N = 10_000_000
const N = 5

func merge(ctx context.Context, ch ...<-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	// Имеется 2 входных канала in1 и in2 и один выходной out.
	// Требуется реализовать функцию merge, которая будет сливать данные из входных каналов в один выходной.
	for _, c := range ch {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				for ch := range c {
					out <- ch
				}
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// в select я проверяю на ctx.Done() при старте горутины,
// но у нас может быть ситуация при которой цикл может заблокироваться
// как это исправить?

// ctrl + D interrupt

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure that the context is cancelled when Run exits

	resL := make([]chan int, N)
	for i := range resL {
		resL[i] = make(chan int)
		go func(i int) {
			// Simulate some work and send data to the channel
			for j := 0; j < 5; j++ { // Example: send 10 integers
				select {
				case resL[i] <- rand.Int():
				case <-ctx.Done():
					return
				}
			}
			close(resL[i]) // Close the channel when done
		}(i)
	}

	// Convert resL to a slice of <-chan int
	inChannels := make([]<-chan int, N)
	for i := range resL {
		inChannels[i] = resL[i]
	}

	rr := merge(ctx, inChannels...)

	// Read from the merged channel
	for val := range rr {
		fmt.Println("Merged Value:", val)
	}
}

func main() {
	Run()
}
