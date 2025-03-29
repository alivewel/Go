package main

func dfs(i, j, size int, grid [][]int) int {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
        return size
    }
    size++
    grid[i][j] = 0
    size = dfs(i-1, j, size, grid)
    size = dfs(i+1, j, size, grid)
    size = dfs(i, j-1, size, grid)
    size = dfs(i, j+1, size, grid)
    return size
}

func largestIsland(grid [][]int) int {
    maxSize := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                size := dfs(i, j, 0, grid)
                if size > maxSize {
                    maxSize = size
                }
            }
        }
    }
    return maxSize
}
