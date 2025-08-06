package main

import (
    "fmt"
)

func test() int {
    x := 5
    defer func() {
        x = x + 10
    }()

    return x
}

func main() {
    fmt.Println(test()) // 5
}
