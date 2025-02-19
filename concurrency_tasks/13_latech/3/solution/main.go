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
		getByIndex(1)
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

// Чем может быть полезно отлавливать панику в main горутине?
// При запуске сервиса при создании grpc-сервера, http-сервера, подключении к БД
// при неуспешном подключении можно отловить панику и как-то прологировать и в метрики это записать.

// Что нельзя отловить в панике?
// stack overflow


// Что нам помогает управлять памятью?
// Сборщик мусора + аллокатор

// Какая модель у сборщик мусора?
// mark and sweap, stop the world и помечает разными цветами
// Можно регулировать работу сборщика мусора через переменную GOGC

// Могут ли быть утечки памяти в Go?
// Может быть ситуация с разрастающейся глобальной переменной, которую GC не может очистить
