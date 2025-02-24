package main

import "fmt"

// вернуть из функции ошибку, не подключая доп. пакетов
func main() {
	fmt.Println(handle())
}

func handle() error {
	return &myError{}
}

type myError struct{}

func (m *myError) Error() string {
	return "Error!"
}

// для того, чтобы вернуть ошибку нам нужна создать структуру, которая удовлетворяет интерфейсу Error
// для этого нужно имплементировать метод Error(), который возвращает string

// type Error interface {
// 	Error() string
// }