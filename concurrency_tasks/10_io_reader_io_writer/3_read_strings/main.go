package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("some string: hello world")

	// read 1024 bytes at a time
	buf := make([]byte, 10)

	// Make a buffer to store the data
	for {
		// Read from file into buffer
		n, err := r.Read(buf)

		// If bytes were read (n > 0), process them
		if n > 0 {
			fmt.Println("Received", n, "bytes:", string(buf[:n]))
		}

		// Check for errors, but handle EOF properly
		if err != nil {
			if err == io.EOF {
				break // End of file, stop reading
			}
			panic(err) // Handle other potential errors
		}
	}
}
