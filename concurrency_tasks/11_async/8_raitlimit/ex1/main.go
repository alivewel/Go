package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	iterationsNum = 6
	goroutinesNum = 5
	quotaLimit    = 2
)

func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{} // ratelim.go, беpëм свободный слот
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		// if j%2== 0 {
		//   <-quotaCh //ratelim.gо, возврашаем слот
		//   quotaCh <- struct{}} // ratelim.gо, берëм слот
		// }
		runtime.Gosched() // даём поработать другим горутинам
	}
	<-quotaCh //ratelim.go, Bозврашаем слот
}

func main() {
	wg := &sync.WaitGroup{} // wait_2.go инициализируем группу
	quotaCh := make(chan struct{}, quotaLimit)
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1) // wait_2.go добавляем воркер
		go startWorker(i, wg, quotaCh)
	}
	wg.Wait() // wait_2.go ожидаем, пока waiter.Done() не приведет счетчик к 0
}

func formatWork(workerNum, input int) string {
	return fmt.Sprintf("workerNum %d %d\n", workerNum, input)
}
