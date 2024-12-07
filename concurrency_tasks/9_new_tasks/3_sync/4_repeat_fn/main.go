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
			case ch <- in:
			}
		}
		// close(ch)
	}(ch)
	return ch
}

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
	if len(res) != 3 {
		panic("wrong code")
	}
}
