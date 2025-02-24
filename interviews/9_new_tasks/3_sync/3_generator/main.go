package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))
	for x := range pipeline {
		fmt.Println(x)
	}
}

func generator(ctx context.Context, in ...int) <-chan int {
	ch := make(chan int)
	go func(ch chan int) {
		defer close(ch)
		for _, num := range in {
			select {
			case <-ctx.Done():
				return
			case ch <- num: // если канал заблокирован, выполнение не будет продолжаться, пока не освободится место в канале.
				// default: // если канал ch заблокирован (например, если никто не читает из него), это приведет к блокировке горутины.
				// 	ch <- num
			}
		}
	}(ch)
	return ch
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	ch := make(chan int)

	go func(ch chan int) {
		defer close(ch)
		for num := range in {
			select {
			case <-ctx.Done():
				return
			case ch <- num * num:
			}
		}
	}(ch)

	return ch
}

// Функция generator принимает на вход контекст и слайс целых чисел,
// элементы которого последовательно записываются в
// возвращаемый канал.

// Функция squarer принимает на вход контекст и канал целых чисел.
// Функция последовательно читает из канал числа, возводит их в квадрат
// и пишет в возвращаемый канал.

// Обе функции должны уметь завершаться по отмене контекста.
