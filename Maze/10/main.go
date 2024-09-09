package main

import (
	"fmt"
)

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

// Get возвращает значение из матрицы по координатам
func (cw *CaveWrapper) Get(x, y int) int {
	// Проверяем, что индексы находятся в допустимых пределах
	if x < 1 || y < 1 || y > cw.Rows || x > cw.Cols {
		// Обработка ошибки
		panic(fmt.Sprintf("Get: индекс вне диапазона! x: %d, y: %d", x, y))
	}
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

func main() {
	// fmt.Println("Path:")
	cw2 := NewCaveWrapper(1, 2, -1)
	cw2.Set(1, 1, 1)
	cw2.Set(2, 1, 2)
	cw2.Print()
}
