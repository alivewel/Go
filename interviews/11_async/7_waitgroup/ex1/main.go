package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	goroutinesNum = 3
	iterationsNum = 3
)

func startWorker(workerNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(workerNum, j))
		runtime.Gosched()
	}
}

func main() {
	wg := &sync.WaitGroup{} // wait_2.go инициализируем группу
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1) // wait_2.go дбавляем воркер
		go startWorker(i, wg)
	}
	wg.Wait() // wait_2.go ожидаем, пока waiter.Done() не приведет счетчик к 0
}

func formatWork(workerNum, input int) string {
	return fmt.Sprintf("workerNum %d %d\n", workerNum, input)
}
