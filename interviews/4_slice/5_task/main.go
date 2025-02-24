package main

import (
	"fmt"
)

type account struct {
	value int
}

// есть ли отличия в прохождения в цикле range между первым и вторым случаем
func main() {
	s1 := make([]account, 0, 10)
	for i, j := range s1 {
		fmt.Println(i, j)
	}
	fmt.Println("____")
	s2 := make([]account, 10, 10)
	for i, j := range s2 {
		fmt.Println(i, j)
	}
}
