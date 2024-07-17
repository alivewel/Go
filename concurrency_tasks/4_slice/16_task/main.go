package main

import "fmt"

func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Hello!")
	testSlice = append(testSlice, "Hello!")
	test(testSlice)

	// fmt.Println(testSlice) // [Hello! Hello!]
	fmt.Println(testSlice[:3]) // [Hello! Hello! Goodbye!]
}

func test(testSlice []string) {
	testSlice = append(testSlice, "Goodbye!")
}
