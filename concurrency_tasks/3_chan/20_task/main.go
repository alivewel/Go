package main

import "fmt"

// через канал приходят ссылки на скачивание
// нам надо с помощью опередленного количества воркеров их скачать

func main() {
	urlChan := make(chan string)

	go func() {
		urlChan <- "http://www.google.com"
		urlChan <- "http://www.bing.com"
		urlChan <- "http://www.duckduckgo.com"
		urlChan <- "http://www.yandex.ru"
		urlChan <- "http://www.somethingelse.com"
		urlChan <- "http://www.iam_lazy_to_think.net"
		close(urlChan)
	}()

	res := download(urlChan, 3)
	for url := range res {
		fmt.Println(url)
	}
}
func download(in <-chan string, workersNum int) chan string {
	return nil
}
