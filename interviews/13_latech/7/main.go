package main

import "fmt"

type MyError struct {
	data string
}

func (e MyError) Error() string {
	return e.data
}

func main() {
	err := foo(4)
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok")
	}
}

func foo(i int) error {
	if i > 5 {
		return MyError{data: "i > 5"}
	}
	return nil
}
