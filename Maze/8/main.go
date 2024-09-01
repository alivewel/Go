package main

import (
	"fmt"
)

// Определение точки в лабиринте
type Point struct {
	X, Y int
}

// Проверка на валидность координат и проходимость
func isValid(maze [][]int, visited [][]bool, point Point) bool {
	rows, cols := len(maze), len(maze[0])
	return point.X >= 0 && point.X < rows && point.Y >= 0 && point.Y < cols &&
		maze[point.X][point.Y] == 0 && !visited[point.X][point.Y]
}

// Реализация поиска в ширину (BFS)
func bfs(maze [][]int, start, end Point) []Point {
	// Направления движения: вверх, вниз, влево, вправо
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Инициализация очереди для хранения текущего пути и посещённых ячеек
	queue := []Point{start}
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}
	visited[start.X][start.Y] = true

	// Карта для хранения родительских ячеек для восстановления пути
	parent := make(map[Point]*Point)
	parent[start] = nil

	// Поиск в ширину
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Если достигли цели
		if current == end {
			break
		}

		// Исследование соседей
		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}

			if isValid(maze, visited, next) {
				visited[next.X][next.Y] = true
				queue = append(queue, next)
				parent[next] = &current
			}
		}
	}

	// Восстановление пути
	path := []Point{}
	for step := &end; step != nil; step = parent[*step] {
		path = append([]Point{*step}, path...)
	}

	// Если путь пустой, значит нет решения
	if len(path) == 1 && path[0] != start {
		return []Point{}
	}

	return path
}

// Печать лабиринта с путём
func printMazeWithPath(maze [][]int, path []Point) {
	// Отметка пути в лабиринте
	for _, p := range path {
		if maze[p.X][p.Y] == 0 {
			maze[p.X][p.Y] = 2 // Используем 2 для обозначения пути
		}
	}

	// Печать лабиринта
	for _, row := range maze {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("█ ") // Стена
			} else if cell == 2 {
				fmt.Print("x ") // Путь
			} else {
				fmt.Print("  ") // Пустая ячейка
			}
		}
		fmt.Println()
	}
}

func main() {
	// Инициализация лабиринта (0 - пустая ячейка, 1 - стена)
	maze := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	start := Point{0, 0} // Начальная точка
	end := Point{4, 4}   // Конечная точка

	// Нахождение пути
	path := bfs(maze, start, end)

	if len(path) == 0 {
		fmt.Println("Путь не найден.")
	} else {
		fmt.Println("Путь найден:")
		printMazeWithPath(maze, path)
	}
}
