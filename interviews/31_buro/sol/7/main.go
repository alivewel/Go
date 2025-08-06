package main

import (
    "fmt"
)

func main() {
    // ch := make(chan int) // unbuffered channel
    // ch <- 42             // ❌ deadlock!
    // val := <-ch
    // fmt.Println(val)

    ch := make(chan int)
    go func() {
        ch <- 42
    }()

    val := <-ch
    fmt.Println(val) // ✅ 42
}
