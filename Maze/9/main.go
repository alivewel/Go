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
	if vertical {
		return m.Vertical[x*m.Cols+y]
	}
	return m.Horizontal[x*m.Cols+y]
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
		if pf.StepWave(maze, to) {
			break
		}
	}
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
		}{
			{p.X + 1, p.Y, 0}, // здесь убираем вызов At, значение будем получать позднее
			{p.X - 1, p.Y, 0},
			{p.X, p.Y + 1, 0},
			{p.X, p.Y - 1, 0},
		}
		fmt.Println("pf.WaveStep", pf.WaveStep)
		for _, n := range neighbors {
			// Проверка на допустимость координат
			if n.x >= 0 && n.x < maze.Rows && n.y >= 0 && n.y < maze.Cols {
				// Только теперь вызываем At для получения значения
				n.value = maze.At(n.x, n.y, n.x != p.X)
				fmt.Println("Debag1")
				if n.value == 0 && pf.LengthMap.Get(n.x, n.y) == pf.EmptyValue {
					fmt.Println("Debag2")
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
	fmt.Println("pf.path1", path, to.X, to.Y)
	for pf.LengthMap.Get(row, col) != 0 {
		fmt.Println("for 1", pf.LengthMap.Get(row, col))
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
	// maze := MazeWrapper{
	// 	Vertical:   make([]int, 25), // Инициализация вертикальных стен
	// 	Horizontal: make([]int, 25), // Инициализация горизонтальных стен
	// 	Rows:       5,
	// 	Cols:       5,
	// }
	from := Point{1, 1}
	to := Point{3, 4}

	pf := NewPathFinder(maze)
	path := pf.Solve(maze, from, to)

	fmt.Println("Path:", path)
}
