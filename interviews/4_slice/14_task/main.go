package main

func main() {
	tasks := make(chan int, 1)
	tasks <- 1
	tasks <- 2
}

// fatal error: all goroutines are asleep - deadlock!
// произошла паника при попытке записи второго элемента, потому что буфер канала только 1
