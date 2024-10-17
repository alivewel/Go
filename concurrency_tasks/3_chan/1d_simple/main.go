package main

import "fmt"

func main() {
	fmt.Println("main() started")
	c := make(chan string)
	var num int = 12
	// launch anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!" + " I am " + fmt.Sprint(num))
	}(c)

	c <- "John"
	fmt.Println("main() stopped")
}
