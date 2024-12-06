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
	return nil
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	return nil
}
