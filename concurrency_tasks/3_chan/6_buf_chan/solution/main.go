package main

import (
	"fmt"
	// "time"
)

func squares(c chan int) {
	// for i := 0; i <= 10; i++ {
	for i := range c {
		num := <-c
		// fmt.Println(i, num)
		fmt.Println(i, num * num)

	}
	
}

// что мы увидим при запуске программы?

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)
	go squares(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here
	c <- 5
	c <- 6
	c <- 7
	c <- 8
	// time.Sleep(time.Second)
	fmt.Println("main() stopped")
}

// Когда размер буфера больше 0, горутина не блокируется до тех пор, пока буфер не будет заполнен. 
// Текущая горутина не будет заблокирована, пока в канал не будет передано n+1 данных.
