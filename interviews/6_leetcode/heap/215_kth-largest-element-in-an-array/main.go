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


func findKthLargest(nums []int, k int) int {
	// Инициализация кучи
	h := &IntHeap{}
	heap.Init(h)
	
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else if (*h)[0] < num {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	return (*h)[0]
}

func main() {
	// Исходные данные
	numbers := []int{5, 3, 8, 1, 2}
	k := 3
	fmt.Println(findKthLargest(numbers, k))

}

