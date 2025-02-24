package main

import "fmt"

func service() {
	fmt.Println("Hello from service!")
}

// что выведет программа?

func main() {
	fmt.Println("main() started")

	go service()

	select {}

	fmt.Println("main() stopped")
}

// Подобно пустому for{}, пустой select{} так же является валидным, 
// но есть подвох. Как мы уже знаем select блокируется до тех пор, 
// пока один из блоков case не будет выполнен, но так как в пустом select 
// отсутствуют блоки case, горутина не будет разблокирована, и как результат, 
// мы получим deadlock.
