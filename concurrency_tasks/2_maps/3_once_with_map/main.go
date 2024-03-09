package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// selected only unique values
// find bug with reads-writes
func main() {
	alreadyStored := make(map[int]struct{}) // создаем множество
	mu := sync.Mutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0...9
	}
	// 3, 4, 5, 0, 4, 9, 8, 6, 6, 5, 5, 4, 4, 4, 2, 1, 2, 3, 1 ...

	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if _, ok := alreadyStored[doubles[i]]; !ok {
				alreadyStored[doubles[i]] = struct{}{}

				uniqueIDs <- doubles[i]
			}
		}()
	}

	wg.Wait()
	close(uniqueIDs)
	for val := range uniqueIDs {
		fmt.Println(val)
	}
	// close(uniqueIDs) // здесь неправильно
	fmt.Printf("len of ids: %d\n", len(uniqueIDs)) // 0, 1, 2, 3, 4 ...
	fmt.Println(uniqueIDs)
}

// в этом примере слайс интов размером 1000 с дублями
// нам нужно провести дедупликацию
// для этого мы создаем буфферизированный канал uniqueIDs и множество alreadyStored
// мы запускаем горутинки и в каждой из них проверяем, есть ли у нас это значение

// Данный код запускается не с первого раза, а при "успешном" запуске мы получаем deadlock!
// Это происходит из-за того, что наш канал не закрыт
// В цикле for range мы прочли все значения, которые есть в канале,
// планировщик видит, что у нас нет других горутин, только main
// и понимает, что никто не отправит закрытие канала

// после закрытия канала мы все равно видим, что код запускается не всегда
// go run -race main.go
// если запустить с флагом -race, то вообще не запустится

// в логе ошибок мы видим: concurrent map read and map write
// в этом примере происходит сначала чтение, а потом запись
// на 30 строке происходит чтение, а на 32 запись
// запись у нас защищено, а чтение нет
// чтобы решить проблему нужно вынести до чтения