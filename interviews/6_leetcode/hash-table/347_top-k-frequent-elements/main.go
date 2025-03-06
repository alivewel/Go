package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}
	freqArr := make([][]int, len(nums)+1) // ! freqArr := make([][]int, len(nums)+1)
	for num, count := range freq {
		freqArr[count] = append(freqArr[count], num) 
	}
	res := make([]int, 0, k)
	for i := len(freqArr) - 1; i >= 0; i-- {
		for _, num := range freqArr[i] { // !
			// if num > 0 {
				res = append(res, num) // res := append(res, num)
				if len(res) >= k {
					return res // break - забыл, что вложенный цикл
				}
			// }
		}
	}
	return res
}

func main() {
	nums := []int{5,7,5,6,6,5} // Вывод: [5,6]
	k := 2
	res := topKFrequent(nums, k)
	fmt.Println(res)
}
