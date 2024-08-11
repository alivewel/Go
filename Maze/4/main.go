package main

import (
	"fmt"
)

// MazeWrapper представляет структуру данных для лабиринта
type MazeWrapper struct {
	Vertical   []int
	Horizontal []int
	Rows       int
	Cols       int
}

// drawMaze выводит сгенерированный лабиринт в консоль
func drawMaze(maze MazeWrapper) {
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

func main() {
	// Пример данных для отладки
	maze := MazeWrapper{
		Vertical: []int{1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1,
			1, 1, 1, 1, 1},
		Horizontal: []int{0, 0, 0, 0, 1,
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 1,
			0, 0, 0, 0, 1},
		Rows: 5,
		Cols: 5,
	}

	drawMaze(maze)
}
