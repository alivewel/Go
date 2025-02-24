package main

import (
	"fmt"
	"time"
)

// Что выведет программа?

var m = map[string]int{"a": 1}

func main() {
	go read()
	time.Sleep(1 * time.Second)
	go write()
	time.Sleep(1 * time.Second)
}

func read() {
	for {
		fmt.Println(m["a"])
	}
}

func write() {
	for {
		m["a"] = 2
	}
}
