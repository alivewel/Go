package main

import (
	"fmt"
	"unsafe"
)

func main() {
	nums := make([]int, 1, 2)
	fmt.Println(nums) // <- what's the output? []

	appendSlice(&nums, 1024)
	fmt.Println(nums) // <- what's the output? []

	mutateSlice(nums, 1, 512)
	fmt.Println(nums)                // <- what's the output? []
	fmt.Println(unsafe.Sizeof(nums)) // 4 + 4 + 4 = 12
}

func appendSlice(sl *[]int, val int) {
	*sl = append(*sl, val)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}

// При попытке запустить изначальную программу мы получаем панику:
// [0]
// [0]
// panic: runtime error: index out of range [1] with length 1

// Паника получается из-за того что все таки при попытке обратиться по индексу
// нужно учитывать length, а не capacity.
// Это не интуитивно понятно.

// Что нужно сделать, чтобы наши методы appendSlice и mutateSlice заработали и избавиться от паники?
// Передавать слайс в appendSlice по указателю. Таким образом у нас length увеличиться до 2
// и обращаться по индексу станет безопасно.

// [0]
// [0 1024]
// [0 512]
// 24

// Также стоит запомнить что структура слайса занимает 24 байта
// указатель 8 байт + int 8 байт + int 8 байт

// Для того чтобы узнать сколько памяти занимает элемент
// нужно использовать метод Sizeof из библиотеки unsafe

// https://go.dev/play/p/sFiz8bJutcS
