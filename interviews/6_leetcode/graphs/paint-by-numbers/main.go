package main

// создадим массив visited
// проходимся по циклу

func dfs(i, j, col int, visited, grid [][]int) {
    if i >= len(visited) || i < 0 || j >= len(visited[0]) || j < 0 || visited[i][j] == 1 || grid[i][j] != col  {
        return
    }
    visited[i][j] = 1

    dfs(i, j-1, col, visited, grid)
    dfs(i, j+1, col, visited, grid)
    dfs(i-1, j, col, visited, grid)
    dfs(i+1, j, col, visited, grid)
}

func draw(grid [][]int) int {
    visited := make([][]int, len(grid))
    for i := range visited {
        visited[i] = make([]int, len(grid[0]))
    }
    res := 0
    for i := range grid {
        for j := range grid[i] {
            if visited[i][j] != 1 {
                res++
                dfs(i, j, grid[i][j], visited, grid)
            }
        }
    }
    return res
}
