package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура для представления графа
type Graph struct {
	vertices   int
	graph      map[int]GraphInfo
	maxTime    int
	sortedList []int
}

type GraphInfo struct {
	dependencies []int
	time         int
	totalTime    int
}

// Функция для создания нового графа
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		graph:    make(map[int]GraphInfo),
	}
}

func (g *Graph) getGraph() map[int]GraphInfo {
	return g.graph
}

// Функция для добавления ребра в граф
func (g *Graph) addEdge(u, v int) {
	if u < g.vertices {
		// Получаем текущую информацию о графе для вершины u
		info := g.graph[u]
		// Добавляем зависимость
		info.dependencies = append(info.dependencies, v)
		// Сохраняем обновленную информацию обратно в граф
		g.graph[u] = info
	}
}

// Функция для добавления ребра в граф
func (g *Graph) addTime(u, time int) {
	// Получаем текущую информацию о графе для вершины u
	info := g.graph[u]
	// Добавляем время
	info.time = time
	// Сохраняем обновленную информацию обратно в граф
	g.graph[u] = info
}

// Рекурсивная функция, используемая в topologicalSort
func (g *Graph) topologicalSortUtil(v int, visited []bool, stack *[]int) {
	// Помечаем текущий узел как посещенный
	visited[v] = true

	// Рекурсивно вызываем функцию для всех смежных вершин
	for _, i := range g.graph[v].dependencies {
		if !visited[i] {
			g.topologicalSortUtil(i, visited, stack)
		}
	}

	*stack = append(*stack, v)
}

// Функция для поиска топологической сортировки
func (g *Graph) topologicalSort() {
	// Помечаем все вершины как непосещенные
	visited := make([]bool, g.vertices+1)
	var stack []int

	// Вызываем рекурсивную вспомогательную функцию
	// для поиска топологической сортировки для каждой вершины
	for i := 1; i <= g.vertices; i++ {
		if !visited[i] {
			g.topologicalSortUtil(i, visited, &stack)
		}
	}

	// Выводим содержимое стека
	g.sortedList = stack
}

// Функция для получения времени для заданной вершины
func (g *Graph) getTime(vertex int) int {
	if info, exists := g.graph[vertex]; exists {
		return info.time
	}
	return -1 // Возвращаем -1, если вершина не существует
}

func (g *Graph) getTotalTime(vertex int) int {
	if info, exists := g.graph[vertex]; exists {
		return info.totalTime
	}
	return -1 // Возвращаем -1, если вершина не существует
}

func (g *Graph) getDependencies(vertex int) []int {
	if info, exists := g.graph[vertex]; exists {
		return info.dependencies
	}
	return []int{} // Возвращаем пустой слайс, если зависимостей нет
}

// Функция для установки totalTime для заданной вершины
func (g *Graph) setTotalTime(u int, totalTime int) {
	if u < g.vertices+1 {
		// Получаем текущую информацию о графе для вершины u
		info := g.graph[u]
		// Устанавливаем новое значение totalTime
		info.totalTime = totalTime
		// Сохраняем обновленную информацию обратно в граф
		g.graph[u] = info
	}
}

func (g *Graph) calcTotalTime() int {
	maxValue, totalTime := 0, 0
	for _, i := range g.sortedList {
		i := i
		dependencies := g.getDependencies(i)
		totalTime = g.getTime(i)
		for i, j := range dependencies {
			currentTotalTime := g.getTotalTime(j)
			if i == 0 {
				maxValue = currentTotalTime
			} else {
				if currentTotalTime > maxValue {
					maxValue = currentTotalTime
				}
			}
		}
		if len(dependencies) > 0 {
			totalTime += maxValue
		}
		g.setTotalTime(i, totalTime)

		if totalTime > maxValue {
			maxValue = totalTime
		}
	}
	g.maxTime = maxValue
	return maxValue
}

func (g *Graph) printTotalTime() {
	for i := 1; i <= g.vertices; i++ {
		fmt.Println(g.getTime(i))
	}
	fmt.Println("maxTime", g.maxTime)
}

func scanNums(n int) [][]int {

	arrays := make([][]int, n) // Создаем срез для массивов

	scanner := bufio.NewScanner(os.Stdin) // Создаем новый сканер для ввода

	for i := 0; i < n; i++ {
		scanner.Scan() // Считываем строку

		line := scanner.Text()          // Получаем текст строки
		numbers := strings.Fields(line) // Разбиваем строку на элементы

		for _, num := range numbers {
			value, err := strconv.Atoi(num) // Преобразуем строку в число
			if err == nil {
				arrays[i] = append(arrays[i], value) // Добавляем элемент в массив
			}
		}
	}
	return arrays
}

func main() {
	var n int
	fmt.Scan(&n) // количество массивов

	g := NewGraph(n)
	nums := scanNums(n) // массивы

	for i, num := range nums {
		for j, array := range num {
			if j == 0 {
				g.addTime(i+1, array)
			} else {
				g.addEdge(i+1, array)
			}
		}
	}
	g.topologicalSort()

	fmt.Println(g.calcTotalTime())
}
