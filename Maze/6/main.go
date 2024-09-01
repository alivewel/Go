package main

import (
	"fmt"
)

// Координаты точки в лабиринте
type Point struct {
	X, Y int
}

// Проверяем, является ли точка внутри лабиринта и проходимой
func isValid(maze [][]rune, visited [][]bool, point Point) bool {
	rows, cols := len(maze), len(maze[0])
	if point.X >= 0 && point.X < rows && point.Y >= 0 && point.Y < cols &&
		maze[point.X][point.Y] == ' ' && !visited[point.X][point.Y] {
		return true
	}
	return false
}

// Алгоритм поиска пути BFS
func bfs(maze [][]rune, start, end Point) []Point {
	// Направления движения: вверх, вниз, влево, вправо
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Очередь для хранения текущего пути и посещённые ячейки
	queue := []Point{start}
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}
	visited[start.X][start.Y] = true

	// Карта для хранения родительской ячейки
	parent := make(map[Point]*Point)
	parent[start] = nil

	// Поиск в ширину
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Если достигли конца
		if current == end {
			break
		}

		// Проверяем всех соседей
		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}

			if isValid(maze, visited, next) {
				visited[next.X][next.Y] = true
				queue = append(queue, next)
				parent[next] = &current
			}
		}
	}

	// Восстанавливаем путь
	path := []Point{}
	for step := &end; step != nil; step = parent[*step] {
		path = append([]Point{*step}, path...)
	}

	return path
}

// Печать лабиринта с путём
func printMazeWithSolution(maze [][]rune, path []Point) {
	// Отмечаем путь в лабиринте
	for _, p := range path {
		if maze[p.X][p.Y] == ' ' {
			maze[p.X][p.Y] = 'x'
		}
	}

	// Печать лабиринта
	for _, row := range maze {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func main() {
	// Инициализация лабиринта
	maze := [][]rune{
		[]rune("+---+---+---+"),
		[]rune("|   |       |"),
		[]rune("+   +   +   +"),
		[]rune("|   |   |   |"),
		[]rune("+   +   +   +"),
		[]rune("|       |   |"),
		[]rune("+---+---+   +"),
	}

	start := Point{1, 1} // Начальная точка
	end := Point{5, 8}   // Конечная точка

	// Нахождение пути
	path := bfs(maze, start, end)

	path = []Point{
		Point{1, 2},
		Point{2, 2},
		Point{3, 2},
		Point{4, 2},
		Point{5, 2},
		Point{5, 4},
		Point{5, 6},
		Point{4, 6},
		Point{3, 6},
		Point{2, 6},
		Point{1, 6},
	}

	fmt.Println(path)

	// Печать лабиринта с отмеченным путём
	printMazeWithSolution(maze, path)
}
