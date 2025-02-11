package main

import "fmt"

func hIndex(citations []int) int {
	length := len(citations)
	// count := make([]int, 0, length + 1)
	count := make([]int, length+1) // правильно создавать таким образом, мы обращаемся по индексу
	for _, cit := range citations { 
		if cit > length {
			count[length]++
		} else {
			count[cit]++
		}
	}
	fmt.Println(count)
	cnt := 0
	for i := len(count) - 1; i >= 0; i-- {
		cnt += count[i]
		if cnt >= i {
			return i // Возвращаем i, так как это h-индекс
		}
	}
	return 0
}

func main() {
	nums := []int{10, 1, 8, 0, 3}
	fmt.Println(hIndex(nums))
}
