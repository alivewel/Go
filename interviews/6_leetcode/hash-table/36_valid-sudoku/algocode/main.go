package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make(map[[2]int]struct{}) // ключ: номер строки и элемент строки
	cols := make(map[[2]int]struct{}) // ключ: номер столбца и элемент столбца
	blocks := make(map[[2]int]struct{})
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] == '.' { // забыл про эту штуку
				continue
			}
            val := int(board[i][j])
            if _, found := rows[[2]int{i, val}]; found {
                return false
            }
            rows[[2]int{i, val}] = struct{}{}
            if _, found := cols[[2]int{j, val}]; found {
                return false
            }
            cols[[2]int{j, val}] = struct{}{}
            //          2  
            blockIdx := i / 3 * 3 + j  / 3
            if _, found := blocks[[2]int{blockIdx, val}]; found {
                return false
            }
            blocks[[2]int{blockIdx, val}] = struct{}{}
        }
    }
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	// board := [][]byte{
	// 	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// }
	fmt.Println(isValidSudoku(board))
}
