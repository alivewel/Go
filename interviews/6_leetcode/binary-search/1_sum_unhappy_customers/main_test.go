package main

import (
	"testing"
)

func TestFindGood(t *testing.T) {
	tests := []struct {
		name   string
		goods  []int
		needs  []int
		want   int
	}{
		{
			name:  "Пример из условия",
			goods: []int{8, 3, 5},
			needs: []int{5, 14, 12, 44, 55},
			want:  93,
		},
		{
			name:  "Простой случай",
			goods: []int{1, 4, 6},
			needs: []int{2, 5, 10},
			want:  6, // 1 + 1 + 4
		},
		{
			name:  "Все нужды между товарами",
			goods: []int{10, 20, 30},
			needs: []int{15, 25, 35},
			want:  15, // 5 + 5 + 5
		},
		{
			name:  "Один товар",
			goods: []int{100},
			needs: []int{50, 150, 100},
			want:  100, // 50 + 50 + 0
		},
		{
			name:  "Нет покупателей",
			goods: []int{5, 10, 15},
			needs: []int{},
			want:  0,
		},
		{
			name:  "Одинаковые товары",
			goods: []int{10, 10, 10},
			needs: []int{5, 10, 15},
			want:  5 + 0 + 5, // 10
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindGood(tt.goods, tt.needs)
			if got != tt.want {
				t.Errorf("FindGood(%v, %v) = %d; want %d", tt.goods, tt.needs, got, tt.want)
			}
		})
	}
}
