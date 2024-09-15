package main

import (
	"fmt"
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

	// Добавляем текущую вершину в стек с результатом
	// if v != 0 {
	// *stack = append([]int{v}, *stack...)
	// }
	*stack = append(*stack, v)

	fmt.Println(*stack)
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
	fmt.Println("Following is a Topological Sort of the given graph")
	fmt.Println(stack)
}

// Функция для получения времени для заданной вершины
func (g *Graph) getTime(vertex int) int {
	if info, exists := g.graph[vertex]; exists {
		return info.time
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
	u = u + 1 // Если вы используете смещение на +1
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
	for i := 1; i <= g.vertices; i++ {
		dependencies := g.getDependencies(i)
		if len(dependencies) == 0 {
			totalTime = g.getTime(i)
			g.setTotalTime(i, totalTime)

		} else {
			for j := 1; j <= len(dependencies); j++ {
				totalTime += g.getTime(j)
			}
			g.setTotalTime(i, totalTime+g.getTime(i))
		}
		if totalTime > maxValue {
			maxValue = totalTime
		}
		fmt.Println(g.getTime(i))
	}
	g.maxTime = maxValue
	return maxValue
}

func main() {
	// g := NewGraph(6)
	// g.addEdge(5, 2)
	// g.addEdge(5, 0)
	// g.addEdge(4, 0)
	// g.addEdge(4, 1)
	// g.addEdge(2, 3)
	// g.addEdge(3, 1)

	// g := NewGraph(5)
	// g.addTime(1, 10)
	// g.addTime(2, 5)
	// g.addTime(3, 0)
	// g.addTime(4, 4)
	// g.addTime(5, 15)

	// g.addEdge(1, 2)
	// g.addEdge(1, 3)
	// g.addEdge(1, 5)
	// g.addEdge(2, 4)
	// g.addEdge(5, 3)

	g := NewGraph(6)
	g.addTime(1, 2)
	g.addTime(2, 2)
	g.addTime(3, 15)
	g.addTime(4, 1)
	g.addTime(5, 2)
	g.addTime(6, 0)

	g.addEdge(1, 2)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	g.addEdge(4, 5)
	g.addEdge(5, 6)

	fmt.Println(g.getGraph())

	// fmt.Println("Following is a Topological Sort of the given graph")
	g.topologicalSort()

	fmt.Println(g.calcTotalTime())
}

// добавить вместо graph    map[int][]int с map[int]struct с массивом и временем выполнения
