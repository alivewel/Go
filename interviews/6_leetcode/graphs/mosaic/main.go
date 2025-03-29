package main

func dfs(i, j, col int, grid, visited [][]int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != col || visited[i][j] == 1 {
        return
    }
    visited[i][j] = 1
    dfs(i-1, j, col, grid, visited)
    dfs(i+1, j, col, grid, visited)
    dfs(i, j-1, col, grid, visited)
    dfs(i, j+1, col, grid, visited)
}

// решение прошло с первого раза :)

func maxComponents(grid [][]int) int {
    visited := make([][]int, len(grid))
    for i := range visited {
        visited[i] = make([]int, len(grid[0]))
    }
    countColors := make(map[int]int)
    for i := range grid {
        for j := range grid[i] {
            if visited[i][j] != 1 {
                countColors[grid[i][j]]++
                dfs(i, j, grid[i][j], grid, visited)
            }
        }
    }
    maxCount, maxColor := -1, 0
    for color, count := range countColors {
        if count > maxCount {
            maxCount = count
            maxColor = color
        }
    }
    return maxColor
}