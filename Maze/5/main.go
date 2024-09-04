package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type MazeWrapper struct {
	Vertical   []int
	Horizontal []int
	Rows       int
	Cols       int
}

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

func main() {
	// Чтение содержимого файла
	content, err := ioutil.ReadFile("maze5x5.txt")
	// content, err := ioutil.ReadFile("maze4x4.txt")
	// content, err := ioutil.ReadFile("maze.txt")
	// _, err := ioutil.ReadFile("maze.txt")
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
	fmt.Println("Rows:", maze.Rows)
	fmt.Println("Cols:", maze.Cols)
	fmt.Println("Vertical:", maze.Vertical, len(maze.Vertical))
	fmt.Println("Horizontal:", maze.Horizontal, len(maze.Horizontal))

	// drawMaze(maze)

	// mazeGenerationSettings := MazeGenerationSettings{Rows: 3, Cols: 3}

	// maze := Generate(mazeGenerationSettings)

	PrintMaze(maze)
}
