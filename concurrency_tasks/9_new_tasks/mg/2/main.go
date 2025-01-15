package main

import (
	"fmt"
)

// Что выведется при запуске программы?

type Figure interface {
	Area() float64
}

type square struct {
	a float64
}

func (s *square) Area() float64 {
	return s.a * s.a
}

func main() {
	var a *square

	figure := Figure(a)

	fmt.Println(figure == nil)
}
