package main

import (
	"fmt"
	"sync"
)

var i int // i == 0

// goroutine increment global variable i
func worker(wg *sync.WaitGroup, ch chan int) {
    i = <-ch 
	// i = i + 1
	wg.Done()
}
// попробовать с каналами
func main() {
	var wg sync.WaitGroup
    ch := make(chan int)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
        ch <- i
		go worker(&wg, ch)
	}

	// wait until all 1000 goroutines are done
	wg.Wait()

	// value of i should be 1000
	fmt.Println("value of i after 1000 operations is", i)
}
