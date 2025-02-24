package main

import "fmt"

func main() {
	bufChan := make(chan string, 1)
	bufChan <- "first"

	select {
	case str := <-bufChan:
		fmt.Println("read", str)
	case bufChan <- "second":
		fmt.Println("write", <-bufChan, bufChan)
	}
}
