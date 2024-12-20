package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	// Имеется 2 входных канала in1 и in2 и один выходной out.
	// Требуется реализовать функцию merge, которая будет сливать данные из входных каналов в один выходной.
	for _, c := range ch {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for ch := range c {
				out <- ch
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func source(sourceFunc func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- sourceFunc(i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()

	return ch
}

func main() {
	rand.Seed(time.Now().UnixMilli())

	in1 := source(func(i int) int {
		return rand.Int()
	})

	in2 := source(func(i int) int {
		return 1
	})

	out := merge(in1, in2)

	for val := range out {
		fmt.Println("Value:", val)
	}
}
