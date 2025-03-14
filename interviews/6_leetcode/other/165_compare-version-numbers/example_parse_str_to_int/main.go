package main

import (
	"fmt"
)

func main() {
	s := "12345"
	curr := 0
	for i := range s {
		curr = curr * 10 + int(s[i]) - int('0') // Вычитаем '0' для получения числового значения
	}
	fmt.Println(curr)
}
