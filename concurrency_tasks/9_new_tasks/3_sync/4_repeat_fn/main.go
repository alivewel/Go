package main

import (
	"context"
	"math/rand"
)

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	// напишите ваш код здесь
	return nil
}

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	// напишите ваш код здесь
	return nil
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
	if len(res) != 3 {
		panic("wrong code")
	}
}
