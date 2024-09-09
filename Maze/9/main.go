package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	index := x*m.Cols + y

	if index < 0 || x < 0 || y < 0 || x >= m.Rows || y >= m.Cols {
		index = -1
		// fmt.Println("index", index, vertical)
		return 1 // Считаем, что за пределами лабиринта есть стена
	}
	// fmt.Println("index", index, vertical)
	if vertical {
		return m.Vertical[index]
	}
	return m.Horizontal[index]
}

// CaveWrapper представляет матрицу для хранения длины пути
type CaveWrapper struct {
	Data       [][]int
	Rows       int
	Cols       int
	EmptyValue int
}

// NewCaveWrapper создает новый CaveWrapper
func NewCaveWrapper(rows, cols, emptyValue int) CaveWrapper {
	// Создаем срез длиной rows*cols и заполняем его значениями emptyValue
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols) // Инициализируем каждый подмассив
		for j := range data[i] {
			data[i][j] = emptyValue // Заполняем значениями emptyValue
		}
	}

	return CaveWrapper{
		Data:       data,
		Rows:       rows,
		Cols:       cols,
		EmptyValue: emptyValue,
	}

}

// Get возвращает значение из матрицы по координатам
func (cw *CaveWrapper) Get(x, y int) int {
	// Проверяем, что индексы находятся в допустимых пределах
	if x < 1 || y < 1 || y > cw.Rows || x > cw.Cols {
		// Обработка ошибки
		// panic(fmt.Sprintf("Get: индекс вне диапазона! x: %d, y: %d", x, y))
		fmt.Printf("Get: индекс вне диапазона! x: %d, y: %d\n", x, y)
		return math.MaxInt
	}
	// fmt.Println("y, x", y, x)
	// fmt.Println("y, x", y, x, "|", cw.Data[y-1][x-1])
	return cw.Data[y-1][x-1] // сдвиг индексов на -1
}

// Set устанавливает значение в матрице по координатам
func (cw *CaveWrapper) Set(x, y, value int) {
	// Проверяем, что индексы находятся в допустимых пределах
	if x < 1 || y < 1 || y > cw.Rows || x > cw.Cols {
		// Обработка ошибки
		panic(fmt.Sprintf("Set: индекс вне диапазона! x: %d, y: %d", x, y))
	}
	cw.Data[y-1][x-1] = value // сдвиг индексов на -1
}

func (cw *CaveWrapper) Print() {
	fmt.Println("CaveWrapper visualization:", cw.Rows, cw.Cols)

	for i := 1; i <= cw.Rows; i++ { // Итерация по строкам
		for j := 1; j <= cw.Cols; j++ { // Итерация по столбцам
			ind := cw.Get(j, i) // Получаем значение с учетом смещения
			// ind := cw.Get(i, j) // Получаем значение с учетом смещения
			if ind == cw.EmptyValue {
				fmt.Printf("%4d ", 0) // Печать значения элемента
			} else {
				fmt.Printf("%4d ", ind) // Печать значения элемента
			}
		}
		fmt.Println() // Переход на новую строку после печати строки
	}
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
// Rows - строки, Cols - столбцы
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
	// pf.LengthMap.Print()
	for len(pf.OldWave) > 0 {
		if pf.StepWave(maze, to) {
			break
		}
	}
	fmt.Println()
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
	if !maze.IsGood() || from.X >= maze.Rows || from.Y >= maze.Cols || to.X >= maze.Rows+1 || to.Y >= maze.Cols+1 {
		return false
	}
	return true
}

// StepWave выполняет один шаг алгоритма BFS и возвращает true, если достигнута конечная точка
func (pf *PathFinder) StepWave(maze MazeWrapper, to Point) bool {
	pf.WaveStep++
	// pf.LengthMap.Print()
	for _, p := range pf.OldWave {
		neighbors := []struct {
			x, y, value int
			// }{
			// 	{p.X + 1, p.Y, maze.At(p.X, p.Y, false)},   // Проверяем стену снизу
			// 	{p.X - 1, p.Y, maze.At(p.X-1, p.Y, false)}, // Проверяем стену сверху
			// 	{p.X, p.Y + 1, maze.At(p.X, p.Y, true)},    // Проверяем стену справа
			// 	{p.X, p.Y - 1, maze.At(p.X, p.Y-1, true)},  // Проверяем стену слева
			// }
		}{
			{p.X + 1, p.Y, maze.At(p.X-1, p.Y-1, false)}, // Проверяем стену снизу
			{p.X - 1, p.Y, maze.At(p.X-2, p.Y-1, false)}, // Проверяем стену сверху
			{p.X, p.Y + 1, maze.At(p.X-1, p.Y-1, true)},  // Проверяем стену справа
			{p.X, p.Y - 1, maze.At(p.X-1, p.Y-2, true)},  // Проверяем стену слева
		}
		for _, n := range neighbors {
			// Проверка на допустимость координат
			// if p.X == 4 && p.Y == 4 {
			// 	fmt.Println("Debag2.1", p.X, p.Y, "|", Point{n.x, n.y}, "|", pf.WaveStep, "|", pf.LengthMap.Get(n.x, n.y), "|", n.value)
			// }
			// if n.x >= 0 && n.x < maze.Rows && n.y >= 0 && n.y < maze.Cols {
			if n.x > 0 && n.x <= maze.Rows && n.y > 0 && n.y <= maze.Cols {
				// fmt.Println("val:", n.value, "| x, y:", n.x, n.y, "|", pf.LengthMap.Get(n.x, n.y) == pf.EmptyValue, "|", "vert:", n.x != p.X, pf.WaveStep)
				if n.value == 0 && pf.LengthMap.Get(n.y, n.x) == pf.EmptyValue {
					// fmt.Println("Debag2.2", p.X, p.Y, "|", Point{n.x, n.y}, "|", pf.WaveStep)
					pf.Wave = append(pf.Wave, Point{n.x, n.y})
					// pf.LengthMap.Set(n.x, n.y, pf.WaveStep)
					pf.LengthMap.Set(n.y, n.x, pf.WaveStep)
					if n.x == to.X && n.y == to.Y { // if p.X == to.X && p.Y == to.Y {
						// if p.X == to.X && p.Y == to.Y {
						// fmt.Println("Debag3", pf.Wave, p.X, p.Y, n.x, n.y)
						return true
					}
				}
			}
		}
	}

	pf.OldWave = pf.Wave
	// pf.OldWave = append([]Point(nil), pf.Wave...) // Создаём копию pf.Wave
	pf.Wave = nil
	return false
}

// MakePath восстанавливает путь из конечной точки в начальную
func (pf *PathFinder) MakePath(maze MazeWrapper, to Point) []Point {
	path := []Point{to}
	row, col := to.X, to.Y
	pf.LengthMap.Print()
	for pf.LengthMap.Get(col, row) != 0 { // Используем (col, row)
		currentLen := pf.LengthMap.Get(col, row) // Используем (col, row)
		fmt.Println()
		fmt.Println("MakePath")
		fmt.Println("currentLen", currentLen-1)
		fmt.Println("row, col", row, col)

		fmt.Println(row < maze.Rows, pf.LengthMap.Get(col, row+1) == currentLen-1, maze.At(row-1, col-1, false) == 0, maze.At(row-1, col-1, false))

		// Проверяем движение влево
		if col > 1 && pf.LengthMap.Get(col-1, row) == currentLen-1 && maze.At(row-1, col-2, true) == 0 {
			fmt.Println("влево")
			col-- // Движение влево
		} else if col < maze.Cols && pf.LengthMap.Get(col+1, row) == currentLen-1 && maze.At(row-1, col-1, true) == 0 {
			fmt.Println("вправо")
			col++ // Движение вправо
		} else if row < maze.Rows && pf.LengthMap.Get(col, row+1) == currentLen-1 && maze.At(row-1, col-1, false) == 0 {
			fmt.Println("вниз")
			row++ // Движение вниз
		} else if row > 1 && pf.LengthMap.Get(col, row-1) == currentLen-1 && maze.At(row-2, col-1, false) == 0 {
			fmt.Println("вверх", pf.LengthMap.Get(col, row-1))
			row-- // Движение вверх
		} else {
			fmt.Println("nil", row, col)
			return nil // Если путь не найден, возвращаем nil
		}

		path = append(path, Point{row, col})
		fmt.Println("pf.path3", path)
	}

	reversePath(path)
	return path
}

// reversePath переворачивает путь
func reversePath(path []Point) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}

func printMaze(maze MazeWrapper) {
	fmt.Println("Maze visualization:")
	for i := 0; i < maze.Rows; i++ {
		// Верхние стены
		for j := 0; j < maze.Cols; j++ {
			fmt.Print("+")
			if i == 0 {
				fmt.Print("-----") // Верхняя граница
			} else {
				if maze.At(i-1, j, false) == 1 {
					fmt.Print("-----") // Горизонтальная стена
				} else {
					fmt.Print("     ")
				}
			}
		}
		fmt.Println("+")

		// Вертикальные стены и пространство между ними
		for j := 0; j < maze.Cols; j++ {
			if j == 0 || maze.At(i, j-1, true) == 1 {
				fmt.Print("|")
				// maze.At(i, j-1, true)
			} else {
				fmt.Print(" ")
			}
			fmt.Print("     ")
		}
		fmt.Println("|")
	}

	// Нижние стены
	for j := 0; j < maze.Cols; j++ {
		fmt.Print("+-----")
	}
	fmt.Println("+")
}

func printMaze2(maze MazeWrapper) {
	fmt.Println("Maze visualization:")
	for i := 0; i < maze.Rows; i++ {
		// Верхние стены
		for j := 0; j < maze.Cols; j++ {
			fmt.Print("+")
			if i == 0 {
				fmt.Print("-----") // Верхняя граница
			} else {
				if maze.At(i-1, j, false) == 1 {
					fmt.Print("-----") // Горизонтальная стена
				} else {
					fmt.Printf("  %d  ", maze.At(i-1, j, false)) // Печать значения ячейки вместо стены
				}
			}
		}
		fmt.Println("+")

		// Вертикальные стены и пространство между ними
		for j := 0; j <= maze.Cols; j++ {
			if j == 0 || maze.At(i, j-1, true) == 1 {
				if j == 0 {
					fmt.Print("|")
				} else {
					fmt.Printf("%d", maze.At(i, j-1, true)) // Печать значения ячейки вместо стены
				}
			} else {
				// fmt.Print(" ")
				fmt.Printf("%d", maze.At(i, j-1, true)) // Печать значения ячейки вместо стены
			}
			fmt.Print("     ")
		}
		fmt.Println()
		// fmt.Println("|")
		// fmt.Println("/")
	}

	// Нижние стены
	for j := 0; j < maze.Cols; j++ {
		fmt.Print("+")
		fmt.Printf("  %d  ", maze.At(maze.Rows-1, j, false)) // Печать значения ячейки вместо нижней стены
	}
	fmt.Println("+")
}

func printMaze3(maze MazeWrapper, path []Point) {
	fmt.Println("Maze visualization:")

	// Создаем карту для быстрого поиска, находится ли точка на пути
	pathMap := make(map[Point]bool)
	for _, p := range path {
		pathMap[p] = true
	}

	for i := 0; i < maze.Rows; i++ {
		// Верхние стены
		for j := 0; j < maze.Cols; j++ {
			fmt.Print("+")
			if i == 0 {
				fmt.Print("-----") // Верхняя граница
			} else {
				if maze.At(i-1, j, false) == 1 {
					fmt.Print("-----") // Горизонтальная стена
				} else {
					fmt.Print("     ")
				}
			}
		}
		fmt.Println("+")

		// Вертикальные стены и пространство между ними
		for j := 0; j < maze.Cols; j++ {
			if j == 0 || maze.At(i, j-1, true) == 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}

			// Проверяем, принадлежит ли текущая позиция пути
			if pathMap[Point{i + 1, j + 1}] {
				fmt.Print("  *  ") // Отображаем путь как '*'
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println("|")
	}

	// Нижние стены
	for j := 0; j < maze.Cols; j++ {
		fmt.Print("+-----")
	}
	fmt.Println("+")
}

func main() {

	// // // Создаем новый CaveWrapper с 5 строками и 5 столбцами, заполняем значением -1
	// cw := NewCaveWrapper(5, 5, -1)

	// // Заполняем CaveWrapper значениями
	// cw.Set(1, 1, 5)
	// cw.Set(1, 2, 15)
	// cw.Set(2, 2, 8)
	// cw.Set(3, 3, 9)
	// cw.Set(4, 4, 7)
	// cw.Set(5, 5, 6)

	// // Печатаем CaveWrapper
	// cw.Print()

	// cw2 := NewCaveWrapper(2, 2, -1)
	// cw2.Set(1, 1, 1)
	// // cw2.Set(1, 2, 2)
	// cw2.Print()
	maze1 := MazeWrapper{
		Vertical: []int{
			1, 1, 1, 1, 1,
			1, 1, 0, 1, 1,
			0, 1, 1, 1, 1,
			1, 0, 0, 1, 1,
			0, 0, 1, 0, 1}, // Инициализация вертикальных стен
		Horizontal: []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			1, 0, 0, 1, 0,
			0, 1, 0, 0, 0,
			1, 1, 1, 1, 1}, // Инициализация горизонтальных стен
		Rows: 5,
		Cols: 5,
	}

	maze := MazeWrapper{
		Vertical: []int{
			1, 2, 3, 4, 5,
			6, 7, 0, 8, 9,
			0, 10, 11, 12, 13,
			14, 0, 0, 15, 16,
			0, 0, 17, 0, 18}, // Инициализация вертикальных стен
		Horizontal: []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			19, 0, 0, 20, 0,
			0, 21, 0, 0, 0,
			22, 23, 24, 25, 26}, // Инициализация горизонтальных стен
		Rows: 5,
		Cols: 5,
	}

	from := Point{1, 1}
	to := Point{2, 1}

	pf := NewPathFinder(maze)
	path := pf.Solve(maze, from, to)

	fmt.Println("Path:", path)

	printMaze(maze1)
	printMaze3(maze1, path)
	// maze2 := MazeWrapper{
	// 	Vertical:   []int{1, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35},                         // Инициализация вертикальных стен
	// 	Horizontal: []int{2, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135}, // Инициализация горизонтальных стен
	// 	Rows:       5,
	// 	Cols:       5,
	// }

	mazeGenerationSettings := MazeGenerationSettings{Rows: 10, Cols: 10}

	maze3 := Generate(mazeGenerationSettings)

	printMaze(maze3)

	from2 := Point{1, 1}
	to2 := Point{10, 1}

	pf2 := NewPathFinder(maze3)
	path2 := pf2.Solve(maze3, from2, to2)

	printMaze3(maze3, path2)
}

// функция At - корректная
// maze.At(i-1, j, false) - для горизонтальных стен
// maze.At(i, j-1, true) - для вертикальных стен

// MazeGenerationSettings содержит настройки для генерации лабиринта
type MazeGenerationSettings struct {
	Rows int
	Cols int
}

// Generate генерирует лабиринт с использованием алгоритма Эллера
func Generate(s MazeGenerationSettings) MazeWrapper {
	vertical := make([]int, s.Rows*s.Cols)
	horizontal := make([]int, s.Rows*s.Cols)

	if !isGood(vertical, s.Rows, s.Cols) || !isGood(horizontal, s.Rows, s.Cols) {
		return MazeWrapper{}
	}

	random := make([]int, s.Rows*s.Cols*3)
	rand.Seed(time.Now().UnixNano())
	for i := range random {
		random[i] = rand.Intn(2)
	}

	randIndex := 0
	counter := 1

	line := make([]int, s.Cols)

	merge := func(i int) {
		mergedItem := line[i+1]
		for col := range line {
			if line[col] == mergedItem {
				line[col] = line[i]
			}
		}
	}

	for row := 0; row < s.Rows-1; row++ {
		for i := range line {
			if line[i] == 0 {
				fmt.Println("line[i]", line[i])
				line[i] = counter
				counter++
				// fmt.Println("line1", line)
				// fmt.Println("line[i]", line[i])
			}
		}
		fmt.Println("---")
		fmt.Println("line", line)
		for col := 0; col < s.Cols-1; col++ {
			if line[col] == line[col+1] {
				vertical[row*s.Cols+col] = 1
			}
		}

		for col := 0; col < s.Cols-1; col++ {
			choice := random[randIndex]
			randIndex++
			if choice == 1 || line[col] == line[col+1] {
				vertical[row*s.Cols+col] = 1
			} else {
				merge(col)
			}
		}
		fmt.Println("row*s.Cols+s.Cols-1", row*s.Cols+s.Cols-1)
		vertical[row*s.Cols+s.Cols-1] = 1

		for col := 0; col < s.Cols; col++ {
			choice := random[randIndex]
			randIndex++
			if choice == 1 {
				count := 0
				for c := 0; c < s.Cols; c++ {
					if line[c] == line[col] && horizontal[row*s.Cols+c] == 0 {
						count++
					}
				}
				if count != 1 {
					horizontal[row*s.Cols+col] = 1
				}
			}
		}

		for col := range line {
			if horizontal[row*s.Cols+col] == 1 {
				line[col] = 0
			}
		}
	}

	for i := range line {
		if line[i] == 0 {
			line[i] = counter
			counter++
		}
	}

	for col := 0; col < s.Cols-1; col++ {
		choice := random[randIndex]
		randIndex++
		if choice == 1 || line[col] == line[col+1] {
			vertical[(s.Rows-1)*s.Cols+col] = 1
		} else {
			merge(col)
		}
	}
	vertical[(s.Rows-1)*s.Cols+s.Cols-1] = 1

	for col := 0; col < s.Cols-1; col++ {
		if line[col] != line[col+1] {
			vertical[(s.Rows-1)*s.Cols+col] = 0
			merge(col)
		}
		horizontal[(s.Rows-1)*s.Cols+col] = 1
	}
	vertical[(s.Rows-1)*s.Cols+s.Cols-1] = 1

	return MazeWrapper{
		Vertical:   vertical,
		Horizontal: horizontal,
		Rows:       s.Rows,
		Cols:       s.Cols,
	}
}

// isGood проверяет корректность данных для генерации лабиринта
func isGood(data []int, rows, cols int) bool {
	if rows == 0 || rows > 150 || cols == 0 || cols > 150 {
		return false
	}
	if rows*cols != len(data) {
		return false
	}

	for _, val := range data {
		if val != 0 && val != 1 {
			return false
		}
	}
	return true
}

// drawMaze выводит сгенерированный лабиринт в консоль
func PrintMaze(maze MazeWrapper) {
	width, height := maze.Cols, maze.Rows

	for x := 0; x < width; x++ {
		fmt.Print("+---")
	}
	fmt.Println("+")

	for y := 0; y < height; y++ {
		// Печать левой границы и содержимого текущей строки
		for x := 0; x < width+1; x++ {
			if x == 0 || maze.Horizontal[y*width+x-1] == 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("   ")
		}
		fmt.Println()

		// Печать нижней границы текущей строки
		for x := 0; x < width; x++ {
			fmt.Print("+")
			if maze.Vertical[y*width+x] == 1 {
				fmt.Print("---")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("+")
	}
}
