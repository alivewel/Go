package main

import (
	"fmt"
)

// Структура для представления графа
type Graph struct {
	vertices int
	// graph    map[int][]int
	graph map[int]GraphInfo
}

type GraphInfo struct {
	dependencies []int
	time         int
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
	*stack = append([]int{v}, *stack...)

	fmt.Println(*stack)
}

// Функция для поиска топологической сортировки
func (g *Graph) topologicalSort() {
	// Помечаем все вершины как непосещенные
	visited := make([]bool, g.vertices)
	var stack []int

	// Вызываем рекурсивную вспомогательную функцию
	// для поиска топологической сортировки для каждой вершины
	for i := 0; i < g.vertices; i++ {
		if !visited[i] {
			g.topologicalSortUtil(i, visited, &stack)
		}
	}

	// Выводим содержимое стека
	fmt.Println("Following is a Topological Sort of the given graph")
	fmt.Println(stack)
}

func main() {
	// g := NewGraph(6)
	// g.addEdge(5, 2)
	// g.addEdge(5, 0)
	// g.addEdge(4, 0)
	// g.addEdge(4, 1)
	// g.addEdge(2, 3)
	// g.addEdge(3, 1)

	g := NewGraph(5)
	g.addEdge(5, 2)
	// g.addEdge(5, 2)

	fmt.Println(g.getGraph())

	// fmt.Println("Following is a Topological Sort of the given graph")
	g.topologicalSort()
}

// добавить вместо graph    map[int][]int с map[int]struct с массивом и временем выполнения
