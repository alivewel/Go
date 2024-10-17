package main

import "fmt"

// Объясните работу программы

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)
	fmt.Println("main()")
	c <- "John"
	fmt.Println("main() stopped")
}
