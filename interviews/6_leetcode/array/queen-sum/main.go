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
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			row, col := rows[i], cols[j]
			d1 := d1Sum[i - j + m - 1]
			d2 := d2Sum[i + j]
			curSum := row + col + d1 + d2
			curSum -= 3 * board[i][j]
			if i == 0 && j == 0 { // костыль, чтобы отрицательные числа в начале проходили
				maxRes = curSum
			}
			if curSum > maxRes {
				maxRes = curSum
			}
		}
	}
	return maxRes
}

func main() {
	// board := [][]int{
    //     {2, 2, 1, 1},
    //     {1, 2, 1, 1},
    //     {1, 2, 2, 1},
    // }
	board := [][]int{
        {-5},
    }
	fmt.Println(maxQueenSum(board))
}
