package main

import "fmt"

func numIslands(grid [][]byte) int {
	counter := int(0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// Случай для верхнего левого угла (i == 0 && j == 0)
			if i == 0 && j == 0 && grid[i][j] == '1' {
				counter++
				fmt.Println(i, j)
			}
			// Случай для первой строки (i == 0)
			if i == 0 && j > 0 && grid[i][j-1] == '0' && grid[i][j] == '1' {
				counter++
				fmt.Println(i, j)
			}
			// Случай для первого столбца (j == 0)
			if j == 0 && i > 0 && grid[i-1][j] == '0' && grid[i][j] == '1' {
				counter++
				fmt.Println(i, j)
			}
			// Общий случай для остальных ячеек
			if i > 0 && j > 0 && grid[i-1][j] == '0' && grid[i][j-1] == '0' && grid[i][j] == '1' {
				counter++
				fmt.Println(i, j)
			}
		}
	}
    return counter
}

func main() {
	// grid := [][]byte{
    //     {'1', '1', '1', '1', '0'},
    //     {'1', '1', '0', '1', '0'},
    //     {'1', '1', '0', '0', '0'},
    //     {'0', '0', '0', '0', '0'},
    // }
	grid := [][]byte{
        {'1', '0', '1', '0', '1'},
        {'1', '0', '1', '0', '1'},
        {'1', '1', '1', '1', '1'},
        {'0', '0', '0', '0', '0'},
    }
	res := numIslands(grid)
	fmt.Println(res)
}
