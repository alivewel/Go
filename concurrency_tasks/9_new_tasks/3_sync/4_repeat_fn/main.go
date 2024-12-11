package main

import (
	"context"
	"fmt"
	"math/rand"
)

// Функция repeatFn бесконечно вызывает функцию fn и пишет результат ее работы в возвращаемый канал.
// Прекращает работу раньше, если контекст отменен.

// Функция take читает не более чем num из канала in, пока in открыт, и пишет значение в возвращаемый канал.
// Прекращает работу раньше, если контекст отменен.

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func(chan interface{}) {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- fn():
			}
		}
	}(ch)
	return ch
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	ch := make(chan interface{})
	go func(chan interface{}) {
		defer close(ch)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			// case ch <- in:
			case v, ok := <-in:
				if !ok {
					return
				}
				ch <- v
			}
		}
		// close(ch)
	}(ch)
	return ch
}

// из задания (не мое):

// // repeatFn создает канал, который будет повторять результат вызова функции fn
// func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
// 	out := make(chan interface{}) // Создаем выходной канал

// 	go func() {
// 		defer close(out) // Закрываем канал при выходе из горутины
// 		for {
// 			select {
// 			case <-ctx.Done(): // Если контекст отменен, выходим
// 				return
// 			case out <- fn(): // Отправляем результат вызова функции fn в канал
// 			}
// 		}
// 	}()
// 	return out // Возвращаем выходной канал
// }

// // take создает канал, который будет принимать num значений из входного канала in
// func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
// 	out := make(chan interface{}) // Создаем выходной канал

// 	go func() {
// 		defer close(out) // Закрываем канал при выходе из горутины
// 		for i := 0; i < num; i++ {
// 			select {
// 			case <-ctx.Done(): // Если контекст отменен, выходим
// 				return
// 			case v, ok := <-in: // Получаем значение из входного канала
// 				if !ok { // Если входной канал закрыт, выходим
// 					return
// 				}
// 				out <- v // Отправляем полученное значение в выходной канал
// 			}
// 		}
// 	}()
// 	return out // Возвращаем выходной канал
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rand := func() interface{} {
		return rand.Int()
	}
	var res []interface{}
	for num := range take(ctx, repeatFn(ctx, rand), 3) {
		res = append(res, num)
	}
	fmt.Println(res)
	for _, v := range res {
		if v != nil {
			fmt.Printf("%T %v", v, v)
			fmt.Println(v) // Разыменовываем указатель
		}
	}
	// res2 := []interface{}{1, 2, 3}
	// fmt.Println(res2) // Вывод: [1 2 3]
	if len(res) != 3 {
		panic("wrong code")
	}
}
