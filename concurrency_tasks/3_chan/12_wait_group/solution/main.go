package main

import (
	"fmt"
	"sync"
	"time"
)

// что такое WaitGroup? для чего он используется?
// для чего нужны методы Add, Done, Wait?
// что выведет данная программа?

func service(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() // decrement counter
}

func main() {
	fmt.Println("main() started")
	var wg sync.WaitGroup // create waitgroup (empty struct)

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment counter
		go service(&wg, i)
	}

	wg.Wait() // blocks here
	fmt.Println("main() stopped")
}

// Это структура со счетчиком, которая отслеживает сколько горутин вами было создано, и сколько из них было завершено.
// Достижение счетчиком нуля будет означать, что все горутины были выполнены.
// WaitGroup дожидается полного завершения всех горутин.

// Метод Add принимает int аргумент, который является delta (дельтой) для счетчика WaitGroup.
// Где счетчика — это число со значением, по умолчанию равным 0.
// Он хранит число запущенных горутин. Когда WaitGroup создана, значение счетчика будет равно 0,
// и мы можем увеличивать его, передавая delta как параметр метода Add. Счетчика не понимает
// автоматически, когда была запущена программа, поэтому нам нужно вручную увеличивать его,
// используя функцию Add.

// Метод Wait используется для блокировки текущей горутины, когда мы его вызываем.
// Как только счетчик достигнет 0, горутина будет разблокирована.
// Поэтому нам необходимо как-то уменьшать значение счетчика.

// Метод Done уменьшает значение счетчика. Он не принимает никаких параметров.
// Eсли посмотреть исходники пакета sync, то можно увидеть, что внутри себя он просто вызывает Add(-1).

// Вывод программы:
// main() started
// Service called on instance 1
// Service called on instance 3
// Service called on instance 2
// main() stopped
