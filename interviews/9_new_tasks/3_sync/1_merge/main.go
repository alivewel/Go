package main

import (
	"fmt"
	"sync"
)

func fillChan(n int) <-chan int {
	ch := make(chan int)
	go func(chan int) {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}

// merge - соединяет каналы в один
func merge(cs ...<-chan int) <-chan int {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for _, val := range cs {
		wg.Add(1)
		go func(<-chan int) {
			defer wg.Done()
			for in := range val {
				ch <- in
			}
		}(val)
	}

	go func(chan int) {
		wg.Wait()
		close(ch)
	}(ch)

	return ch
}

func main() {
	a := fillChan(2)
	b := fillChan(3)

	in := merge(a, b)

	for val := range in {
		fmt.Println(val)
	}
}
