package main

import (
	"fmt"
	"math/rand"
	"time"
)

func initializeFirstRow(width int) []int {
	sets := make([]int, width)
	for i := range sets {
		sets[i] = i
	}
	return sets
}

func connectRight(sets []int, width int) []bool {
	rightConnections := make([]bool, width-1)

	for i := 0; i < width-1; i++ {
		if sets[i] != sets[i+1] && rand.Intn(2) == 0 {
			rightConnections[i] = true
			sets[i+1] = sets[i]
		}
	}

	return rightConnections
}

func connectDown(sets []int, width int) []bool {
	downConnections := make([]bool, width)
	nextSets := make([]int, width)

	for i := 0; i < width; i++ {
		if rand.Intn(2) == 0 || i == 0 || sets[i] != sets[i-1] {
			downConnections[i] = true
			nextSets[i] = sets[i]
		} else {
			nextSets[i] = -1
		}
	}

	for i := 0; i < width; i++ {
		sets[i] = nextSets[i]
	}

	return downConnections
}

func ensureVerticalConnections(sets []int, downConnections []bool, width int) {
	uniqueSets := make(map[int]bool)

	for i := 0; i < width; i++ {
		if sets[i] != -1 {
			uniqueSets[sets[i]] = true
		}
	}

	for uniqueSet := range uniqueSets {
		inSet := []int{}
		for i := 0; i < width; i++ {
			if sets[i] == uniqueSet && downConnections[i] == false {
				inSet = append(inSet, i)
			}
		}

		if len(inSet) > 0 {
			downConnections[inSet[rand.Intn(len(inSet))]] = true
		}
	}
}

func finalizeLastRow(sets []int, width int) {
	for i := 0; i < width-1; i++ {
		if sets[i] != sets[i+1] {
			sets[i+1] = sets[i]
		}
	}
}

func generateMaze(width, height int) ([][]bool, [][]bool) {
	rand.Seed(time.Now().UnixNano())
	sets := initializeFirstRow(width)

	horizontalConnections := make([][]bool, height)
	verticalConnections := make([][]bool, height)

	for y := 0; y < height; y++ {
		horizontalConnections[y] = connectRight(sets, width)

		if y < height-1 {
			verticalConnections[y] = connectDown(sets, width)
			ensureVerticalConnections(sets, verticalConnections[y], width)
		} else {
			finalizeLastRow(sets, width)
		}
	}

	return horizontalConnections, verticalConnections
}

func drawMaze(horizontalConnections, verticalConnections [][]bool, width, height int) {
	for y := 0; y < height; y++ {
		// Верхняя граница каждой строки
		for x := 0; x < width; x++ {
			fmt.Print("+")
			if y == height-1 || !verticalConnections[y][x] {
				fmt.Print("---")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("+")

		// Левая граница и горизонтальные соединения
		for x := 0; x < width; x++ {
			if x == 0 || !horizontalConnections[y][x-1] {
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
		fmt.Print("----")
	}
	fmt.Println("+")
}

func main() {
	width, height := 3, 3
	horizontalConnections, verticalConnections := generateMaze(width, height)
	drawMaze(horizontalConnections, verticalConnections, width, height)
}
