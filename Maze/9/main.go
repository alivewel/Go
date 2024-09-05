package main

import (
	"fmt"
	"math"
)

// Point представляет координаты точки в лабиринте
type Point struct {
	X, Y int
}

// MazeWrapper представляет лабиринт с его границами и размерами
type MazeWrapper struct {
	Vertical   []int
	Horizontal []int
	Rows       int
	Cols       int
}

// IsGood проверяет, является ли лабиринт корректным
func (m *MazeWrapper) IsGood() bool {
	return len(m.Vertical) == m.Rows*m.Cols && len(m.Horizontal) == m.Rows*m.Cols
}

// At возвращает значение ячейки в вертикальной или горизонтальной линии
func (m *MazeWrapper) At(x, y int, vertical bool) int {
	// index := (x-1)*m.Cols + y - 1
	// index := (x)*m.Cols + y - 1
	index := (x)*m.Cols + y
	if index < 0 {
		// fmt.Println("")
		fmt.Println(index)
		panic("index out of range")
		// return -1
	}
	if vertical {
		// index = (x-1)*m.Cols + y - 1
		// fmt.Println("index:", index, "val:", m.Vertical[index], "| x, y:", x, y, "Vertical")
		return m.Vertical[index]
		// return m.Horizontal[index]
	}
	// fmt.Println("index:", index, "val:", m.Horizontal[index], "| x, y:", x, y, "Horizontal")

	// index = (x-2)*m.Cols + y- 1
	return m.Horizontal[index]
	// return m.Vertical[index]
}

// CaveWrapper представляет матрицу для хранения длины пути
type CaveWrapper struct {
	Data       []int
	Rows       int
	Cols       int
	EmptyValue int
}

// NewCaveWrapper создает новый CaveWrapper
func NewCaveWrapper(rows, cols, emptyValue int) CaveWrapper {
	// Создаем срез длиной rows*cols и заполняем его значениями emptyValue
	data := make([]int, rows*cols)
	for i := range data {
		data[i] = emptyValue
	}

	return CaveWrapper{
		Data:       data,
		Rows:       rows,
		Cols:       cols,
		EmptyValue: emptyValue,
	}
}

// Index возвращает индекс ячейки в линейном массиве
func (cw *CaveWrapper) Index(x, y int) int {
	return x*cw.Cols + y
}

// Get возвращает значение из матрицы по координатам
func (cw *CaveWrapper) Get(x, y int) int {
	return cw.Data[cw.Index(x, y)]
}

// Set устанавливает значение в матрице по координатам
func (cw *CaveWrapper) Set(x, y, value int) {
	cw.Data[cw.Index(x, y)] = value
}

// PathFinder представляет алгоритм поиска пути
type PathFinder struct {
	LengthMap  CaveWrapper
	Wave       []Point
	OldWave    []Point
	WaveStep   int
	EmptyValue int
}

// NewPathFinder создает новый PathFinder
func NewPathFinder(maze MazeWrapper) *PathFinder {
	return &PathFinder{
		LengthMap:  NewCaveWrapper(maze.Rows, maze.Cols, math.MaxInt),
		Wave:       []Point{},
		OldWave:    []Point{},
		WaveStep:   0,
		EmptyValue: math.MaxInt,
	}
}

// Solve ищет путь от начальной точки до конечной в лабиринте
func (pf *PathFinder) Solve(maze MazeWrapper, from, to Point) []Point {
	if !pf.IsValid(maze, from, to) {
		return nil
	}

	pf.InitializeStartState(maze, from)
	for len(pf.OldWave) > 0 {
		// fmt.Println("Debag01")
		if pf.StepWave(maze, to) {
			// fmt.Println("Debag02")
			break
		}
	}
	// fmt.Println("Debag03", pf.OldWave)
	fmt.Println("Debag04", pf.Wave)
	return pf.MakePath(maze, to)
}

// InitializeStartState инициализирует начальное состояние для поиска пути
func (pf *PathFinder) InitializeStartState(maze MazeWrapper, from Point) {
	pf.Wave = nil
	pf.WaveStep = 0
	pf.OldWave = []Point{from}
	pf.LengthMap = NewCaveWrapper(maze.Rows, maze.Cols, pf.EmptyValue)
	pf.LengthMap.Set(from.X, from.Y, pf.WaveStep)
}

// IsValid проверяет, находятся ли точки внутри границ лабиринта
func (pf *PathFinder) IsValid(maze MazeWrapper, from, to Point) bool {
	if !maze.IsGood() || from.X >= maze.Rows || from.Y >= maze.Cols || to.X >= maze.Rows || to.Y >= maze.Cols {
		return false
	}
	return true
}

// StepWave выполняет один шаг алгоритма BFS и возвращает true, если достигнута конечная точка
func (pf *PathFinder) StepWave(maze MazeWrapper, to Point) bool {
	pf.WaveStep++
	for _, p := range pf.OldWave {
		neighbors := []struct {
			x, y, value int
			// }{
			// 	{p.X + 1, p.Y, 0}, // здесь убираем вызов At, значение будем получать позднее
			// 	{p.X - 1, p.Y, 0},
			// 	{p.X, p.Y + 1, 0},
			// 	{p.X, p.Y - 1, 0},
			// }
		}{
			{p.X + 1, p.Y, maze.At(p.X, p.Y, false)},
			{p.X - 1, p.Y, maze.At(p.X-1, p.Y, false)},
			{p.X, p.Y + 1, maze.At(p.X, p.Y, true)},
			{p.X, p.Y - 1, maze.At(p.X, p.Y-1, true)},
		}
		// fmt.Println("neighbors", neighbors)
		for ind, n := range neighbors {
			// Проверка на допустимость координат
			// if n.x >= 0 && n.x < maze.Rows && n.y >= 0 && n.y < maze.Cols {
			if n.x > 0 && n.x < maze.Rows && n.y > 0 && n.y < maze.Cols {
				// Только теперь вызываем At для получения значения
				// n.value = maze.At(n.x, n.y, n.x == p.X)
				// n.value = maze.At(n.x-1, n.y-1, n.x != p.X)
				// fmt.Println()
				_ = ind
				// fmt.Println("val:", n.value, "| x, y:", n.x, n.y, "|", "vert:", n.x != p.X, "|", pf.LengthMap.Get(n.x, n.y) == pf.EmptyValue, "| ind", ind)
				if n.value == 0 && pf.LengthMap.Get(n.x, n.y) == pf.EmptyValue {
					fmt.Println("Debag2.2", pf.WaveStep, Point{n.x, n.y})
					pf.Wave = append(pf.Wave, Point{n.x, n.y})
					pf.LengthMap.Set(n.x, n.y, pf.WaveStep)
					if n.x == to.X && n.y == to.Y {
						fmt.Println("Debag3")
						return true
					}
				}
			}
		}
	}
	pf.OldWave = pf.Wave
	pf.Wave = nil
	return false
}

// MakePath восстанавливает путь из конечной точки в начальную
func (pf *PathFinder) MakePath(maze MazeWrapper, to Point) []Point {
	path := []Point{to}
	row, col := to.X, to.Y
	// fmt.Println("pf.path1", path, to.X, to.Y)
	for pf.LengthMap.Get(row, col) != 0 {
		// fmt.Println("for 1", pf.LengthMap.Get(row, col))
		if col > 0 && pf.LengthMap.Get(row, col-1)+1 == pf.LengthMap.Get(row, col) && maze.At(row, col-1, true) == 0 {
			col--
		} else if col+1 < maze.Cols && pf.LengthMap.Get(row, col+1)+1 == pf.LengthMap.Get(row, col) && maze.At(row, col+1, true) == 0 {
			col++
		} else if row > 0 && pf.LengthMap.Get(row-1, col)+1 == pf.LengthMap.Get(row, col) && maze.At(row-1, col, false) == 0 {
			row--
		} else if row+1 < maze.Rows && pf.LengthMap.Get(row+1, col)+1 == pf.LengthMap.Get(row, col) && maze.At(row+1, col, false) == 0 {
			row++
		} else {
			return nil
		}
		path = append(path, Point{row, col})
	}
	fmt.Println("pf.path2", path)
	reversePath(path)
	fmt.Println("pf.path3", path)
	return path
}

// reversePath переворачивает путь
func reversePath(path []Point) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}

func main() {
	// Пример использования
	maze := MazeWrapper{
		Vertical:   []int{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1}, // Инициализация вертикальных стен
		Horizontal: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1}, // Инициализация горизонтальных стен
		Rows:       5,
		Cols:       5,
	}
	// fmt.Println("maze2.At", maze.At(2, 1, false))
	// maze2 := MazeWrapper{
	// 	Vertical:   []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35},                          // Инициализация вертикальных стен
	// 	Horizontal: []int{111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135}, // Инициализация горизонтальных стен
	// 	Rows:       5,
	// 	Cols:       5,
	// }
	// fmt.Println("maze2.At", maze2.At(1, 1, false))

	// maze := MazeWrapper{
	// 	Vertical:   make([]int, 25), // Инициализация вертикальных стен
	// 	Horizontal: make([]int, 25), // Инициализация горизонтальных стен
	// 	Rows:       5,
	// 	Cols:       5,
	// }
	from := Point{1, 1}
	// to := Point{3, 1}
	to := Point{4, 3}

	pf := NewPathFinder(maze)
	path := pf.Solve(maze, from, to)

	fmt.Println("Path:", path)
}
