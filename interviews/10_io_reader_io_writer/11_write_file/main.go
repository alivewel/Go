package main

import (
	"os"
)

func main() {
	// f, err := os.Open("data/file.txt") // panic: write file.txt: bad file descriptor
	f, err := os.OpenFile("data/file.txt", os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close() // no error handling

	_, err = f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}
}
