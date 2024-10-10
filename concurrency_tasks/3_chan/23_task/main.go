package main

import (
	"fmt"
	"time"
)

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

// concurrent map read and map write

// чтобы отловить проблему с конкурентной записью
// можно запустить код с флагом data race
