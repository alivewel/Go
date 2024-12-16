package main

import (
	"sync"
)

var (
	c int
	m sync.Mutex // Mutex для защиты доступа к переменной c
)

type Trace struct {
	// Some fields here
}

type Sender interface {
	Send(Trace)
}

func Do(sender Sender, tr Trace) {
	m.Lock() // Блокируем доступ к переменной c
	defer m.Unlock() // Разблокируем в конце функции

	c++
	if c == 100 {
		sender.Send(tr)
		c = 0
	}
}