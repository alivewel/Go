package main

import (
	"fmt"
	"math/rand"
)

// Что выведется при запуске программы?

func main() {
	fmt.Println("Start")

	res := make(chan int, 1)
	res <- rand.Int()

	fmt.Println("Finish", <-res)
}
