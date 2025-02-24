package main

import (
	"net/http"
)

// эта функция лезет по сети в старый монолит и может тупить
// эту функцию изменять нельзя
func getDiscount() *http.Response {
	discount, _ := http.Get("http://discounts.com/my")
	return discount
}

// как сделать влияние функции getDiscount() на всю программу
// более контроллируемым
func main() {
	getDiscount()
}
