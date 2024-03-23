package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// эта функция лезет по сети в старый монолит и может тупить
// эту функцию изменять нельзя
func getDiscount() int {
	// sleep 0-4 sec
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	return rand.Int()
}

// как сделать влияние функции getDiscount() на всю программу
// более контроллируемым
func main() {
	rand.New(rand.NewSource(time.Now().Unix()))
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
	getDiscountWithContext(ctx)
	fmt.Println(getDiscount())
	cancel()
}

func getDiscountWithContext(ctx context.Context) int {
	// ctx := 
	// ctx.Err() 
	return getDiscount()
}
