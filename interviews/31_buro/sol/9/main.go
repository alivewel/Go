package main

import (
    "context"
    "fmt"
    "time"
)

const processingMaxTimeout = 3 * time.Second

func main() {
    var result int

    ctx, cancel := context.WithTimeout(context.Background(), processingMaxTimeout)
    defer cancel()

    done := make(chan int, 1)

    go func() {
        result := process()
        done <- result
    }()

    select {
    case res := <-done:
        result = res
    case <-ctx.Done():
        result = 0
    }

    fmt.Println(result)
}

func process() int {
    // long processing job
    time.Sleep(time.Second * 2)

    return 1
}
