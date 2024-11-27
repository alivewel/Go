package main

import (
	"fmt"
)

func main() {
	cancelCh := make(chan struct {})
	dataCh := make(chan int)
	go func(cancelCh chan struct {}, datach chan int) {
		val := 0
		for {
			select {
			case <- cancelCh:
				return
			case datach <- val:
				val++
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh { 
		fmt.Println("read", curVal) 
		if curVal > 3 {
			fmt.Println("send cancel") 
			cancelCh <- struct{}{} 
			break
		}
	}
}
