package main

import "fmt"

// получить количество символов в строке
func main() {
	s := "word"

	// fmt.Println(len([]byte(s)))
	// fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
}
