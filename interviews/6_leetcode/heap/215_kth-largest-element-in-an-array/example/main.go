package main

import (
	"container/heap"
	"fmt"
)

// Определение мин-кучи
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // Мин-куча: меньшее значение наверху
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func main() {
	// Исходные данные
	numbers := []int{5, 3, 8, 1, 2}

	// Инициализация кучи
	h := &IntHeap{}
	heap.Init(h)

	// Добавляем элементы в кучу
	for _, num := range numbers {
		if h.Len() < 3 {
			heap.Push(h, num)
		} else if (*h)[0] < num {
			heap.Pop(h)
			heap.Push(h, num)
		}
		
	}

	fmt.Println("Извлечение элементов по возрастанию:")
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h)) // будет выводить: 1, 2, 3, 5, 8
	}
}
