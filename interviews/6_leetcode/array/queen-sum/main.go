package main

import (
	"fmt"
)

func maxQueenSum(board [][]int) int {
	n := len(board) // количество строк
	m := len(board[0]) // количество столбцов
	rows := make([]int, n)
	cols := make([]int, m)
	d1Sum := make([]int, n+m-1) // пока не знаю какую длину использовать при создании
	d2Sum := make([]int, n+m-1) // пока не знаю какую длину использовать при создании
	maxRes := 0
	
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			rows[i] += board[i][j]
			cols[j] += board[i][j]
			d1Sum[i - j + m - 1] += board[i][j]
			d2Sum[i + j] += board[i][j]
		}
	}
	fmt.Println(rows)
	fmt.Println(cols)
	fmt.Println(d1Sum)
	fmt.Println(d2Sum)
	return maxRes
}

func main() {
	board := [][]int{
        {2, 2, 1, 1},
        {1, 2, 1, 1},
        {1, 2, 2, 1},
    }
	fmt.Println(maxQueenSum(board))
}
