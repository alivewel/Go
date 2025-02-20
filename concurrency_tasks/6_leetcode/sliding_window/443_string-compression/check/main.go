package main

import "fmt"

func main() {
	countChars := 12345
	for countChars != 0 {
		fmt.Println(countChars % 10)
		countChars /= 10
	} 
}
