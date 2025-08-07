package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// что не нравится в коде?
var counter int64

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // counter++
            atomic.AddInt64(&counter, 1)
        }()
    }
    wg.Wait()
    fmt.Println(counter)
}

// После просят рассказать про атомики, rwmutex, что когда лучше применять и так далее