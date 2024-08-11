package main

import "fmt"

// avoid deadlock
func main() {
	ch := make(chan int)
	fmt.Println(len(ch))
	select {
	case val := <-ch:
		fmt.Println(val)
	}
}
