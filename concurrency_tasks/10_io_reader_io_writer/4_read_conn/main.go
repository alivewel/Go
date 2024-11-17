package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// Establish a TCP connection to the server
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close() // Ensure the connection is closed when done

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	// read 1024 bytes at a time
	buf := make([]byte, 10)

	// Make a buffer to store the data
	for {
		// Read from file into buffer
		n, err := conn.Read(buf)

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
