package main

import (
	"testing"
	"math"
)

func TestMinProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"All positive", []int{1, 2, 3, 4, 5}, 2},
		{"Contains zero", []int{5, 2, -1, 0}, -5},
		{"Two large positives", []int{100, 200, 300}, 100 * 200},
		{"Two negatives", []int{-5, -4, -3, -2}, 6},
		{"Negative and positive", []int{-10, -5, 1, 2}, -20},
		{"With duplicates", []int{2, 2, 3, 3}, 4},
		{"Only two elements", []int{-2, 3}, -6},
		{"All negative", []int{-1, -2, -3}, 2},
		{"Contains max int", []int{math.MaxInt32, 1, 2}, 2},
		{"Large and small mix", []int{-1000000, 1000000, 0}, -1000000000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinProduct(tt.nums)
			if result != tt.expected {
				t.Errorf("got %d, expected %d", result, tt.expected)
			}
		})
	}
}
