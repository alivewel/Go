package main

import (
	"fmt"
	"sync"
)

// print square of range 0...20 in random order
func main() {
	counter := 20
	wg := sync.WaitGroup{}
	for i := 0; i < counter; i++ {
		wg.Add(1)
		// i := i // можно внутри цикла объявить новую переменную
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}
	wg.Wait()
}
