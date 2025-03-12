package main

import "fmt"

func dfs(grid [][]byte, x, y int) {
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
        return
    }
    grid[x][y] = '0'
    dfs(grid, x + 1, y)
    dfs(grid, x - 1, y)
    dfs(grid, x, y + 1)
    dfs(grid, x, y - 1)
}

func numIslands(grid [][]byte) int {
    count := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {
                count++
                dfs(grid, i, j)
            }
        }
    }
	return count
}

func main() {
	grid := [][]byte{
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'1', '1', '0', '0', '0'},
        {'0', '0', '0', '0', '0'},
    }
	// grid := [][]byte{
    //     {'1', '0', '1', '0', '1'},
    //     {'1', '0', '1', '0', '1'},
    //     {'1', '1', '1', '1', '1'},
    //     {'0', '0', '0', '0', '0'},
    // }
	res := numIslands(grid)
	fmt.Println(res)
}
