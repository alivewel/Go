package main

import "sync"

// var c int

type Trace struct {
	// Some fields here
}

type Sender interface {
	Send(Trace)
}

type TraceHandler struct {
	Counter int
	Rate    int
	M       sync.Mutex
}

func (th *TraceHandler) Reset() {
	th.Counter = 0
}

func (th *TraceHandler) SendIfNeeded(sender Sender, tr Trace) {
	th.M.Lock()
	th.Counter++
	if th.Counter == th.Rate {
		go sender.Send(tr)
		th.Reset()
	}
	th.M.Unlock()
}

func Do(sender Sender, tr Trace, th *TraceHandler) {
	th.SendIfNeeded(sender, tr)
}

// TraceHandler, содержащая sync.Mutex нельзя передавать по значению в функцию Do
// Это может привести к неожиданному поведению, так как при передаче структуры по значению создается копия, и блокировка может не работать должным образом.