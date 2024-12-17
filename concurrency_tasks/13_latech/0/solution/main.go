package main

import (
	"fmt"
)

// что выведет данная программа?

func foo(src *[]int) {
	*src = append(*src, 5)
}

func main() {
	arr := []int{1, 2, 3}
	src := make([]int, 1)
	copy(src, arr[:1])
	foo(&src)

	fmt.Println(src) // [1]
	fmt.Println(arr) // [1, 5, 3]
}

// исправить программу так, чтобы исходный слайс не менялся
// теперь вывод будет таким
// [1 5]
// [1 2 3]
