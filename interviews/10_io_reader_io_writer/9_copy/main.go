package main

import (
	"io"
	"log"
	"os"
)

func main() {
	src, err := os.Open("data/source.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create("data/destination.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
