package main

import (
	"fmt"
	"reflect"
)

// **Чемпионат по шагам.**  
// Недавно мы устроили чемпионат по шагам. И вот настало время подводить итоги! Необходимо определить `userids` участников, которые прошли наибольшее количество шагов `steps` за все дни, не пропустив ни одного дня соревнований.

// #Пример 1 # ввод
// statistics = [
//     [{ userId: 1, steps: 1000 }, { userId: 2, steps: 1500 }],
//     [{ userId: 2, steps: 1000 }],
// ]
// # вывод
// champions = { userIds: [2], steps: 2500 }

// #Пример 2
// statistics = [
//     [{ userId: 1, steps: 2000 }, { userId: 2, steps: 1500 }],
//     [{ userId: 2, steps: 4000 }, { userId: 1, steps: 3500 }],
// ]
// # вывод
// champions = { userIds: [1, 2], steps: 5500 }

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
	if len(Statistics) == 0 {
		return res
	}
	// 1 создаем мапу key - UserId, value - кол-во шагов и кол-во дней участия
	userStat := make(map[int][2]int, 0)
	for _, stat := range Statistics {
		for _, st := range stat {
			prevValue := userStat[st.UserId]
			userStat[st.UserId] = [2]int{prevValue[0] + st.Steps, prevValue[1] + 1}
		}
	}
	countDaysChemp := len(Statistics)
	maxSteps := 0
	for userId, stat := range userStat {
		if stat[1] == countDaysChemp {
			if stat[0] == maxSteps {
				res.UserIds = append(res.UserIds, userId)
			}
			if stat[0] > maxSteps {
				maxSteps = stat[0]
				res.Steps = stat[0]
				res.UserIds = []int{userId}
				// res.UserIds = append(res.UserIds, userId)
			}
		}
	}

	return res
}

func main() {
	// stats := [][]Statistics{
	// 	{{UserId: 3, Steps: 100}, {UserId: 7, Steps: 400}, {UserId: 2, Steps: 150}},
	// 	{{UserId: 2, Steps: 200}, {UserId: 3, Steps: 200}},
	// }
	stats := [][]Statistics{{}}
	result := getChampions(stats)
	fmt.Printf("Champions: %+v\n", result)

	// Тесты
	runTests()
}

func runTests() {
	tests := []struct {
		name     string
		input    [][]Statistics
		expected Result
	}{
		{
			name: "Basic test",
			input: [][]Statistics{
				{{UserId: 3, Steps: 100}, {UserId: 7, Steps: 400}, {UserId: 2, Steps: 150}},
				{{UserId: 2, Steps: 200}, {UserId: 3, Steps: 200}},
			},
			expected: Result{
				UserIds: []int{2},
				Steps:   350,
			},
		},
		{
			name: "Two winners",
			input: [][]Statistics{
				{{UserId: 1, Steps: 100}, {UserId: 2, Steps: 100}},
				{{UserId: 1, Steps: 200}, {UserId: 2, Steps: 200}},
			},
			expected: Result{
				UserIds: []int{1, 2},
				Steps:   300,
			},
		},
		{
			name: "Only one with full participation",
			input: [][]Statistics{
				{{UserId: 1, Steps: 100}},
				{{UserId: 1, Steps: 200}, {UserId: 2, Steps: 400}},
			},
			expected: Result{
				UserIds: []int{1},
				Steps:   300,
			},
		},
		{
			name: "No full participants",
			input: [][]Statistics{
				{{UserId: 1, Steps: 100}},
				{{UserId: 2, Steps: 200}},
			},
			expected: Result{
				UserIds: nil,
				Steps:   0,
			},
		},
	}

	for _, tt := range tests {
		got := getChampions(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			fmt.Printf("Test %s FAILED: expected %+v, got %+v\n", tt.name, tt.expected, got)
		} else {
			fmt.Printf("Test %s PASSED\n", tt.name)
		}
	}
}
