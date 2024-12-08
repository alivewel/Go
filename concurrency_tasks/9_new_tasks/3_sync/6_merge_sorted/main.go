package main

import "fmt"

func mergeSorted(a, b <-chan int) <-chan int {
	out := make(chan int)
	go func(out chan int) {
		defer close(out)
		res1, ok1 := <-a
		res2, ok2 := <-b
		if ok1 && ok2 {
			if res1 > res2 {
				out <- res2
				out <- res1
			} else {
				out <- res1
				out <- res2
			}
		}
		if ok1 {
			out <- res1
		}
		if ok2 {
			out <- res2
		}
	}(out)
	return out
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)
	c := mergeSorted(a, b)
	for val := range c {
		fmt.Printf("%d ", val)
	}
}
