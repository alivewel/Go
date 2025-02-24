package main

import "fmt"

// Что если мы хотим изменить внутренний массив и не хотим менять внешний?

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("before", arr)
	handle(arr[:1])
	fmt.Println("after", arr)
}

func handle(arr []int) {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	newArr = append(newArr, 5)
	fmt.Println("append", newArr)
}

// before [1 2 3 4]
// append [1 5]
// after [1 2 3 4]

// Вывод
// Если хотим изменить переданный слайс, создаем его копию

// Иногда действительно нужно изменять переданный в функцию слайс, например если мы хотим отсортировать слайс
