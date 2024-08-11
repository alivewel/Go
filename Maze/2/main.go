package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MazeWrapper представляет сгенерированный лабиринт
type MazeWrapper struct {
	Vertical   []int
	Horizontal []int
	Rows       int
	Cols       int
}

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
				line[i] = counter
				counter++
			}
		}

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

// // PrintMaze выводит сгенерированный лабиринт в консоль
// func PrintMaze(m MazeWrapper) {
// 	for row := 0; row < m.Rows; row++ {
// 		// Печатаем горизонтальные стены
// 		for col := 0; col < m.Cols; col++ {
// 			if row == 0 { // Верхняя граница лабиринта
// 				// fmt.Print("1")
// 				continue
// 			} else if m.Horizontal[(row-1)*m.Cols+col] == 1 {
// 				fmt.Print("---")
// 			} else {
// 				// fmt.Print("  ")
// 				fmt.Print("   ")
// 			}
// 		}
// 		// fmt.Println()
// 		fmt.Println("+")

// 		// Печатаем вертикальные стены и пространство для ячеек
// 		for col := 0; col < m.Cols; col++ {
// 			if m.Vertical[row*m.Cols+col] == 1 {
// 				fmt.Print("|")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 			// fmt.Print("_")
// 			fmt.Print("   ")
// 		}
// 		// fmt.Println("|") // Правая граница лабиринта
// 		fmt.Println("+")
// 	}

// 	// Нижняя граница лабиринта
// 	for x := 0; x < m.Cols; x++ {
// 		fmt.Print("----")
// 	}
// 	fmt.Println("+")
// }

// drawMaze выводит сгенерированный лабиринт в консоль
func PrintMaze(maze MazeWrapper) {
	width, height := maze.Cols, maze.Rows

	for y := 0; y < height; y++ {
		// Верхняя граница каждой строки
		for x := 0; x < width; x++ {
			fmt.Print("+")
			if y == height-1 || maze.Vertical[y*width+x] == 1 {
				fmt.Print("---")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("+")

		// Левая граница и горизонтальные соединения
		for x := 0; x < width; x++ {
			if x == 0 || maze.Horizontal[y*width+x-1] == 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("   ")
		}
		fmt.Println("|")
	}

	// Нижняя граница лабиринта
	for x := 0; x < width; x++ {
		fmt.Print("+---")
	}
	fmt.Println("+")
}

func main() {
	// width, height := 3, 3
	// horizontalConnections, verticalConnections := generateMaze(width, height)
	// PrintMaze(horizontalConnections, verticalConnections, width, height)

	mazeGenerationSettings := MazeGenerationSettings{Rows: 3, Cols: 3}
	maze := Generate(mazeGenerationSettings)
	PrintMaze(maze)
}

// // MazeGenerationSettings содержит настройки для генерации лабиринта
// type MazeGenerationSettings struct {
// 	Rows int
// 	Cols int
// }

// // Generate генерирует лабиринт с использованием алгоритма Эллера
// func Generate(s MazeGenerationSettings) MazeWrapper {
