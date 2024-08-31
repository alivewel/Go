package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type MazeWrapper struct {
	Vertical   []int
	Horizontal []int
	Rows       int
	Cols       int
}

func parseFileContent(content string) (MazeWrapper, error) {
	lines := strings.Split(content, "\n")

	// Первая строка - размеры лабиринта
	dimensions := strings.Fields(lines[0])
	if len(dimensions) != 2 {
		return MazeWrapper{}, fmt.Errorf("неверный формат размеров")
	}

	rows, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return MazeWrapper{}, err
	}

	cols, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return MazeWrapper{}, err
	}

	vertical := make([]int, 0, rows*cols)
	horizontal := make([]int, 0, rows*cols)

	// Остальные строки - данные вертикальных и горизонтальных линий
	for i := 1; i <= rows; i++ {
		values := strings.Fields(lines[i])
		if len(values) != cols {
			return MazeWrapper{}, fmt.Errorf("1неверное количество значений в строке %d", i)
		}

		for _, v := range values {
			num, err := strconv.Atoi(v)
			if err != nil {
				return MazeWrapper{}, err
			}
			vertical = append(vertical, num)
		}
	}

	for i := rows + 2; i <= 2*rows+1; i++ {
		values := strings.Fields(lines[i])
		if len(values) != cols {
			return MazeWrapper{}, fmt.Errorf("2неверное количество значений в строке %d", i)
		}

		for _, v := range values {
			num, err := strconv.Atoi(v)
			if err != nil {
				return MazeWrapper{}, err
			}
			horizontal = append(horizontal, num)
		}
	}

	return MazeWrapper{
		Vertical:   vertical,
		Horizontal: horizontal,
		Rows:       rows,
		Cols:       cols,
	}, nil
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

// func PrintMaze(maze MazeWrapper) {
// 	width, height := maze.Cols, maze.Rows

// 	for x := 0; x < width; x++ {
// 		fmt.Print("+---")
// 	}
// 	fmt.Println("+")

// 	for y := 0; y < height; y++ {
// 		// Печать левой границы и содержимого текущей строки
// 		for x := 0; x < width+1; x++ {
// 			if x == 0 || maze.Horizontal[y*width+x-1] == 1 {
// 				fmt.Print("|")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 			fmt.Print("   ")
// 		}
// 		fmt.Println()

// 		// Печать нижней границы текущей строки
// 		for x := 0; x < width; x++ {
// 			fmt.Print("+")
// 			if maze.Vertical[y*width+x] == 1 {
// 				fmt.Print("---")
// 			} else {
// 				fmt.Print("   ")
// 			}
// 		}
// 		fmt.Println("+")
// 	}
// }

func PrintMaze(maze MazeWrapper) {
	width, height := maze.Cols, maze.Rows

	// Печать верхней границы лабиринта
	for x := 0; x < width; x++ {
		fmt.Print("+---")
	}
	fmt.Println("+")

	for y := 0; y < height; y++ {
		// Печать левой границы и содержимого текущей строки
		for x := 0; x < width; x++ {
			if x == 0 || maze.Vertical[y*width+x-1] == 1 { // Используем вертикальный массив
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("   ")
		}
		// Печать правой границы строки
		fmt.Println("|")

		// Печать нижней границы текущей строки
		for x := 0; x < width; x++ {
			fmt.Print("+")
			if maze.Horizontal[y*width+x] == 1 {
				fmt.Print("---")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("+")
	}
}

func main() {
	// Чтение содержимого файла
	// content, err := ioutil.ReadFile("maze5x5.txt")
	// content, err := ioutil.ReadFile("maze4x4.txt")
	content, err := ioutil.ReadFile("maze.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Парсинг содержимого в структуру
	maze, err := parseFileContent(string(content))
	if err != nil {
		fmt.Println("Ошибка при парсинге файла:", err)
		return
	}

	// Вывод результата
	// fmt.Println("Rows:", maze.Rows)
	// fmt.Println("Cols:", maze.Cols)
	// fmt.Println("Vertical:", maze.Vertical, len(maze.Vertical))
	// fmt.Println("Horizontal:", maze.Horizontal, len(maze.Horizontal))

	// drawMaze(maze)
	PrintMaze(maze)
}
