package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.com",
	}

	wg := &sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			fmt.Printf("Fetching %s...\n", url)

			err := fetchUrl(url)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}

			fmt.Printf("Fetched %s\n", url)
		}(url)
	}

	fmt.Println("All requests launched!")
	wg.Wait()
	fmt.Println("Program finished.")
}

func fetchUrl(url string) error {
	// Подробная реализация опущена и не относится к теме задачи
	_, err := http.Get(url)
	return err
}

// почему нельзя добавить wg.Add(1) на уровне горутины (внутри нее)?
// мы можем не успеть добавить счетчик.  wg.Wait() отработается быстрее и ни одна горутина не успеет завершить свою работу.

// что будет если будет 100к урлов в срезе?
// 1) у нас могут закончиться порты куда мы ходим
// 2) в юникс системах до 10к соединений на одном хосте могут держать
// 3) также у нас есть ограничение дискриптеров файловой системы.

// Как это можно решить?
// Нужно ограничить количество запросов
// Для этого нужно добавить воркер