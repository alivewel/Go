package main

import (
	"fmt"
	"sync"
	"time"
)

// Что выведет программа?

var m = map[string]int{"a": 1}
var mtx sync.RWMutex

func main() {
	go read()
	time.Sleep(1 * time.Second)
	go write()
	time.Sleep(1 * time.Second)
}

func read() {
	mtx.Lock()
	defer mtx.Unlock()
	for {
		fmt.Println(m["a"])
	}
}

func write() {
	mtx.RLock()
	defer mtx.RUnlock()
	for {
		m["a"] = 2
	}
}

// отличие data race and data condition
