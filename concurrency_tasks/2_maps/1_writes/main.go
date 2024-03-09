package main

import (
	"fmt"
	"sync"
)

// find and fix 2 bugs
func main() {
	var storage map[int]int // Некорректное создание map

	wg := sync.WaitGroup{}
	writes := 1000

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() { // нет мьютексов
			defer wg.Done()
			storage[i] = i
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}

// 1. Некорректное создание map, нужно использовать make
// С помощью записи var storage map[int]int мы только объявляем мапу,
// но ничего нельзя присвоить туда
// 2. Хэш-таблица не concurrency save.
// Мы не можем ее использовать в разных горутинах конкруентно без блокировок.
// Необходимо использование мьютекса из библиотеки sync.
