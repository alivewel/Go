package main

import "fmt"

func change(a *int) {
	t := *a * 2 // 6
	a = &t      // нужно было вот так сделать, чтобы результат изменился *a = t
	// обратились не по указателю
	fmt.Println(&a)
	fmt.Println(a)
}
func main() {
	a := 3
	change(&a)
	fmt.Println(a) // 3
}

// Что выведет код?
// a == 3
