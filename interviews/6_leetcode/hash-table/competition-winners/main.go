package main

import "fmt"

func findCompetitionWinners(statistics [][][]int) []int {
	res := make([]int, 0)
	// Ключ - идентификатор участника
	// Значение - количество дней участия в соревновании и количество пройденных шагов
	ustats := make(map[int][2]int, 0)
	countDay := len(statistics)
	for _, stat := range statistics {
		for _, st := range stat {
			prevVal := ustats[st[0]]
			ustats[st[0]] = [2]int{prevVal[0] + st[1], prevVal[1] + 1}
		}
	}
	max := 0
	for k, v := range ustats {
		if v[1] == countDay {
			if v[0] == max {
				res = append(res, k)
			} else if v[0] > max {
				max = v[0]
				res = res[:0] // обнулить длину
				res = append(res, k)
			}
		}
	}
	return res
}

// Сделать слайс пустым (обнулить длину): slice = slice[:0] 

func main() {
	// nums := []int{10, 1, 8, 0, 3}
	stats := [][][]int{
		{{3, 100}, {7, 400}, {2, 150}},
		{{2, 200}, {3, 200}},
	}
	fmt.Println(findCompetitionWinners(stats))
}
