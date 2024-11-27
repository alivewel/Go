package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello!")
}

func main() {
	timer := time.AfterFunc(1*time.Second, sayHello)

	fmt.Scanln()
	timer.Stop()
	fmt.Scanln()
}
