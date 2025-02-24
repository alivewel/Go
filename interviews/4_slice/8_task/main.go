package main

import "fmt"

type I interface {
	Foo()
}

type S struct{}

func (s *S) Foo() {
	fmt.Println("Foo")
}

func Build() I {
	// return &S{}
	var res *S
	return res
}

func main() {
	i := Build()
	if i != nil {
		i.Foo()
	} else {
		fmt.Println("nil")
	}
}
