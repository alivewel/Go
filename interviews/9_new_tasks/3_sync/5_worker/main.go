package main

import (
	"fmt"
	// "sync"
)

func worker(f func(int) int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- f(job)
	}
}

const numJobs = 5
const numWorkers = 3

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// wg := sync.WaitGroup{}
	multiplier := func(x int) int {
		return x * 10
	}
	// wg.Add(1)
	// defer wg.Done()
	for i := 0; i < numWorkers; i++ {
		go worker(multiplier, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		// jobs <- multiplier(i)
		jobs <- i
	}
	// закрываем канал задач, чтобы рабочие поняли, что больше задач не будет
	close(jobs)

	for i := 0; i < numJobs; i++ {
		res := <-results
		fmt.Println(res)
	}
	// wg.Wait()
}
