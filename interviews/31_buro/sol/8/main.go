package main

import (
    "fmt"
    "sync"
    "time"
)

func doWork(i int) {
    fmt.Printf("Start: %d\n", i)
    time.Sleep(100 * time.Millisecond) // имитация работы
    fmt.Printf("End:   %d\n", i)
}

func main() {
    var wg sync.WaitGroup
    sem := make(chan struct{}, 5) // семафор на 5 "пропусков"

    for i := 0; i < 100; i++ {
        wg.Add(1)

        go func(i int) {
            defer wg.Done()

            sem <- struct{}{}   // захватываем семафор
            defer func() { <-sem }() // освобождаем

            doWork(i)
        }(i)
    }

    wg.Wait()
}
