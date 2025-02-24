package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = map[string]int{"a": 1}
	lock sync.RWMutex
)

func main() {
	go read()
	time.Sleep(1 * time.Second)
	go write()
	time.Sleep(1 * time.Second)
}

func read() {
	for {
		lock.RLock()
		fmt.Println(m["a"])
		lock.RUnlock()
	}
}

func write() {
	for {
		lock.Lock()
		m["a"] = 2
		lock.Unlock()
	}
}