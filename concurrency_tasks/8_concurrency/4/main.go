package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	// Первая горутина для записи данных в канал
	go func() {
		for _, val := range []int{1, 2, 3, 4, 5} {
			ch <- val
		}
		close(ch) // Закрываем канал после записи
	}()

	// Вторая горутина для записи данных в канал
	go func() {
		for _, val := range []int{6, 7, 8, 9, 10} {
			ch <- val
		}
		close(ch) // Закрываем канал после записи
	}()

	// Чтение данных из канала
	for val := range ch {
		fmt.Printf("%d ", val)
	}
}