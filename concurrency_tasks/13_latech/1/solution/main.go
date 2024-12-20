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
		// url := url
		wg.Add(1)
		go func(url string) {
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
		}(url)
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

// 1) Добавить WaitGroup
// 2) Исправь представленную программу, чтобы как только какая-нибудь горутина ответила с ошибкой, то программа завершилась (добавить context)
// Почему нельзя добавлять `wg.Add()` в горутине?

// Проблемы с конкурентным доступом
// Если вы вызываете `wg.Add()` в горутине, это может привести к ситуации, когда одна горутина добавляет счетчик, в то время как другая горутина ожидает зeавершения. Это создает риск гонки данных, так как `wg.Add()` не является атомарной операцией в контексте нескольких горутин.
// Неопределенное поведение
// Если `wg.Add()` вызывается после того, как горутина уже начала выполнение, это может привести к неопределенному поведению. Например, если вы добавляете счетчик после того, как горутина уже завершила свою работу, это может привести к тому, что `wg.Wait()` завершится раньше, чем вы ожидаете.
