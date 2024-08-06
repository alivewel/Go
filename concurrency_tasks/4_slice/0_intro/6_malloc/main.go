package main

import "fmt"

// У нас есть функция для удвоения элементов массива.
// Корректно ли она работает? Можно ли ее как-то оптимизировать?

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(double(arr))
}

func double(nums []int) []int {
	var res []int
	for _, num := range nums {
		// fmt.Println(len(res), cap(res))
		res = append(res, num*2)
	}
	return res
}
