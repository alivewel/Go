package main

import (
	"fmt"
	"sync"
)

// Что выведет программа?

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()

		defer func() {
			if v := recover(); v != nil {
				fmt.Println("got panic", v)
			}
		}()

	}()
	wg.Wait()
}

func getByIndex(index int) {
	var a []int
	fmt.Println(a[index])
}

// Что такое паника?
// Критическая ошибка, пналог исключения в других языках

// Как перехватить панику? Какой синтаксис?
// Мы не можем перехватить панику из другой горутины, только из своей

// Почему в main не отлавливаем панику?
// Потому что мы не может отловить панику в другой горутине
// В main можем отловить только панику в main горутине