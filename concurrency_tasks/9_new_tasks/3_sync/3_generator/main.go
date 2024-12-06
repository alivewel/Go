package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))
	for x := range pipeline {
		fmt.Println(x)
	}
}

func generator(ctx context.Context, in ...int) <-chan int {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for _, num := range in {
		wg.Add(1)
		go func(ch chan int, num int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				ch <- num
			}
		}(ch, num)
	}
	go func(chan int) {
		wg.Wait()
		close(ch)
	}(ch)
	return ch
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	for num := range in {
		wg.Add(1)
		go func(ch chan int, num int) {
			defer wg.Done()
			// ch <- num * num
			select {
			case <-ctx.Done():
				return
			default:
				ch <- num * num
			}
		}(ch, num)
	}
	go func(chan int) {
		wg.Wait()
		close(ch)
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
