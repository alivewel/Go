package main

var c int

type Trace struct {
	// Some fields here
}

type Sender interface {
	Send(Trace)
}

func Do(sender Sender, tr Trace) {
	c++
	if c == 100 {
		sender.Send(tr)
		c = 0
	}
}
