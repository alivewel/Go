package main

import (
	"fmt"
	"sync"
)

// Что выведет программа?

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()

		defer func() {
			if v := recover(); v != nil {
				fmt.Println("got panic", v)
			}
		}()

	}()
	wg.Wait()
}

func getByIndex(index int) {
	var a []int
	fmt.Println(a[index])
}

// 
