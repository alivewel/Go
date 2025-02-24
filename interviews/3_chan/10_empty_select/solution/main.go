package main

import (
	"fmt"
	"time"
)

// что выведет программа?

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}

// В этом примере мы используем оператор select как switch, но вместо булевых операций, 
// мы используем операции для чтения данных из канала. Оператор select также является 
// блокируемым, за исключением использования default(позже вы увидите пример с его использованием). 
// После выполнения одного из блоков case, горутина main будет разблокирована. 
// Задались вопросом когда case условие выполнится?


// Если все блоки case являются блокируемыми, тогда select будет ждать до момента, 
// пока один из блоков case разблокируется и будет выполнен. Если несколько или все
// канальные операции не блокируемы, тогда один из неблокируемых case будет выбран
// случайным образом (Примечание переводчика: имеется ввиду случай, когда пришли 
// одновременно данные из двух и более каналов).

// Вышеприведенная программа имитирует реальный веб-сервис, в котором балансировщик
// нагрузки получает миллионы запросов и должен возвращать ответ от одной из доступных служб. 
// Используя стандартные горутины, каналы и select, мы можем запросить ответ у нескольких сервисов,
// и тот, который ответит раньше всех, может быть использован.