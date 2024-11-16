package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("data/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	body, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
