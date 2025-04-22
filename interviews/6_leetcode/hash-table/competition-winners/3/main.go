package main

import "fmt"

type Result struct {
	UserIds []int
	Steps   int
}

type Statistics struct {
	UserId int
	Steps  int
}

func getChampions(Statistics [][]Statistics) Result {
	res := Result{}
	// ключ userId, значение - кол-во шагов и кол-во участия в днях
	userStat := make(map[int][2]int)
	for _, stat := range Statistics {
		for _, st := range stat {
			prevVal := userStat[st.UserId]
			userStat[st.UserId] = [2]int{prevVal[0] + st.Steps, prevVal[1] + 1}
		}
	}
	maxSteps := 0
	countDayStats := len(statistics)
	for user, resUser := range userStat {
		if countDayStats == resUser[1] { // если user участваовал во всех днях соревнований
			if resUser[0] == maxSteps {
				res = append(res, user)
			}
			if resUser[0] > maxSteps {
				res = res[:0]
				res = append(res, user)
				maxSteps = resUser[0]
			}
		}
	}
	return Result{}
}

func findCompetitionWinners(statistics [][][]int) []int {
	res := make([]int, 0)
	// ключ userId, значение - кол-во шагов и кол-во участия в днях
	userStat := make(map[int][2]int)
	for _, stat := range statistics {
		for _, st := range stat {
			prevVal := userStat[st[0]]
			userStat[st[0]] = [2]int{prevVal[0] + st[1], prevVal[1] + 1}
		}
	}
	// находим максимум
	maxSteps := 0
	countDayStats := len(statistics)
	for user, resUser := range userStat {
		if countDayStats == resUser[1] { // если user участваовал во всех днях соревнований
			if resUser[0] == maxSteps {
				res = append(res, user)
			}
			if resUser[0] > maxSteps {
				res = res[:0]
				res = append(res, user)
				maxSteps = resUser[0]
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
