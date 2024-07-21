package main

import "fmt"

func sender(c chan int) {
	c <- 1 // len 1, cap 3
	c <- 2 // len 2, cap 3
	c <- 3 // len 3, cap 3
	c <- 4 // <- goroutine blocks here
	close(c)
}

// пример для ознакомления, можно отследить как изменяется длина буф. канала при чтении из него

func main() {
	c := make(chan int, 3)

	go sender(c)

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

	// read values from c (blocked here)
	for val := range c {
		fmt.Printf("Length of channel c after value '%v' read is %v\n", val, len(c))
	}
}

// Вывод программы:

// Length of channel c is 0 and capacity of channel c is 3
// Length of channel c after value '1' read is 3
// Length of channel c after value '2' read is 2
// Length of channel c after value '3' read is 1
// Length of channel c after value '4' read is 0
