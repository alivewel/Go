package main

import "fmt"

func numIslands(grid [][]byte) int {
	counter := int(0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i > 0 && j > 0 && grid[i-1][j] == '0' && grid[i][j-1] == '0'  && grid[i][j] == '1' {
				counter++
				// fmt.Print(i, j)
				fmt.Println(i, j)
			}
			if i == 0 && j == 0 && grid[i-1][j] == '0' && grid[i][j-1] == '0'  && grid[i][j] == '1' {
				counter++
				// fmt.Print(i, j)
				fmt.Println(i, j)
			}
			// fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
    return counter
}

func main() {
	grid := [][]byte{
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'1', '1', '0', '0', '0'},
        {'0', '0', '0', '0', '0'},
    }
	res := numIslands(grid)
	fmt.Println(res)
}
