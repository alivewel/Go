package main

import (
	"context"
	"errors"
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
	ctx, cancel := context.WithCancel(context.Background())
	for _, url := range urls {
		url := url
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
			}

			fmt.Printf("Fetching %s...\n", url)
			err := fetchUrl(ctx, url)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				fmt.Printf("Error fetching %s: %v\n", url, err)
				cancel()
				return
			}
			fmt.Printf("Fetched %s\n", url)
		}()
	}

	fmt.Println("All requests launched!")
	wg.Wait()
	fmt.Println("Program finished.")
}

func fetchUrl(ctx context.Context, url string) error {
	// Подробная реализация опущена и не относится к теме задачи
	_, err := http.Get(url)
	_ = ctx
	return err
}
