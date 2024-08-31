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

	for i := rows + 1; i <= 2*rows; i++ {
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

func main() {
	// Чтение содержимого файла
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
	fmt.Println("Rows:", maze.Rows)
	fmt.Println("Cols:", maze.Cols)
	fmt.Println("Vertical:", maze.Vertical)
	fmt.Println("Horizontal:", maze.Horizontal)
}
