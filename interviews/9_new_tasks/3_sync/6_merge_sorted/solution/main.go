package main

import "fmt"

func mergeSorted(a, b <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		val1, ok1 := <-a
		val2, ok2 := <-b
		for ok1 && ok2 {
			// отправляем меньшее значение в канал и читаем из этого же канала еще одно значение
			if val1 < val2 {
				out <- val1
				val1, ok1 = <-a
			} else {
				out <- val2
				val2, ok2 = <-b
			}
		}
		// если канал a длинее чем канал b
		for ok1 {
			out <- val1
			val1, ok1 = <-a
		}
		// если канал b длинее чем канал a
		for ok2 {
			out <- val2
			val2, ok2 = <-b
		}
	}()
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
