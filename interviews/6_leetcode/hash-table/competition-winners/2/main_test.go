package main

import (
	"reflect"
	"testing"
)

// go test -v

func TestFindCompetitionWinners(t *testing.T) {
	tests := []struct {
		name       string
		statistics [][][]int
		want       []int
	}{
		{
			name: "Один победитель участвовал все дни",
			statistics: [][][]int{
				{{1, 100}, {2, 150}},
				{{1, 200}, {3, 50}},
			},
			want: []int{1},
		},
		{
			name: "Несколько победителей с равным количеством шагов",
			statistics: [][][]int{
				{{1, 100}, {2, 200}},
				{{1, 150}, {2, 50}},
			},
			want: []int{1, 2}, // 250 шагов у каждого, участвовали оба дня
		},
		{
			name: "Только один участник прошёл все дни",
			statistics: [][][]int{
				{{1, 100}, {2, 150}},
				{{1, 200}},
			},
			want: []int{1},
		},
		{
			name: "Никто не прошёл все дни",
			statistics: [][][]int{
				{{1, 100}, {2, 150}},
				{{3, 200}},
			},
			want: []int{},
		},
		{
			name: "Пустая статистика",
			statistics: [][][]int{},
			want:       []int{},
		},
		{
			name: "Все пользователи участвуют каждый день, разное количество шагов",
			statistics: [][][]int{
				{{1, 100}, {2, 200}},
				{{1, 150}, {2, 250}},
				{{1, 100}, {2, 50}},
			},
			want: []int{2}, // 500 vs 350
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCompetitionWinners(tt.statistics)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCompetitionWinners() = %v, want %v", got, tt.want)
			}
		})
	}
}
