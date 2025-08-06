package main

import (
    "fmt"
    "sync"
)

// что не нравится в коде?
var counter int

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++
        }()
    }
    wg.Wait()
    fmt.Println(counter)
}

// После просят рассказать про атомики, rwmutex, что когда лучше применять и так далее