package main

import "fmt"

func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Hello!")
	testSlice = append(testSlice, "Hello!")
	test(testSlice)

	// fmt.Println(testSlice) // [Hello! Hello!]
	fmt.Println(len(testSlice))
	fmt.Println(testSlice[:3]) // [Hello! Hello! Goodbye!]
}

func test(testSlice []string) {
	testSlice = append(testSlice, "Goodbye!")
	// fmt.Println(len(testSlice))
}

// слайс в функцию передается по указателю
// данные внутри функции в слайс в итоге записываются
// при выходе из функции длина остается прежней
// для того, чтобы данные распечатались нужно явно передавать границы до третьего индекса массива testSlice[:3]
