package main

import (
	"fmt"
	"sync"
	"time"
)

func getComments() chan string {
	// надо использовать буферизированный канал
	result := make(chan string, 1)
	go func(out chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("async operation ready, return comments")
		out <- "32 комментария"
	}(result)
	return result
}

func getPage(wg *sync.WaitGroup) {
	defer wg.Done()
	resultCh := getComments()

	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")

	commentsData := <-resultCh
	fmt.Println("main goroutine:", commentsData)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go getPage(&wg)
	}
	wg.Wait()
}
