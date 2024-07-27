package main

import "fmt"

// Объясните работу программы
// Что здесь не так? Почему это происходит?

func main() {
	fmt.Println("main() started")

	c := make(chan string)
	c <- "John"

	fmt.Println("main() stopped")
}
