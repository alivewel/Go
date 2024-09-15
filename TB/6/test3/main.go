package main

import (
	"fmt"
)

// Структура для представления графа
type Graph struct {
	vertices int
	graph    map[int]GraphInfo
}

// Структура для хранения информации о графе
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

// Функция для получения графа
func (g *Graph) getGraph() map[int]GraphInfo {
	return g.graph
}

// Функция для добавления ребра в граф
func (g *Graph) addEdge(u, v int) {
	// Получаем текущую информацию о графе для вершины u
	info := g.graph[u]
	// Добавляем зависимость
	info.dependencies = append(info.dependencies, v)
	// Сохраняем обновленную информацию обратно в граф
	g.graph[u] = info
}

// Функция для добавления времени выполнения
func (g *Graph) addTime(u, time int) {
	info := g.graph[u]
	// Добавляем время
	info.time = time
	// Сохраняем обновленную информацию обратно в граф
	g.graph[u] = info
}

// Рекурсивная функция для топологической сортировки
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
	*stack = append(*stack, v)
}

// Функция для поиска топологической сортировки
func (g *Graph) topologicalSort() {
	// Помечаем все вершины как непосещенные
	visited := make([]bool, g.vertices+1)
	var stack []int

	// Вызываем рекурсивную вспомогательную функцию
	for i := 1; i <= g.vertices; i++ {
		if !visited[i] {
			g.topologicalSortUtil(i, visited, &stack)
		}
	}

	// Выводим содержимое стека в обратном порядке
	fmt.Println("Следующая топологическая сортировка данного графа:")
	// for i := len(stack) - 1; i >= 0; i-- {
	// 	fmt.Print(stack[i], " ")
	// }
	// fmt.Println()
	
	fmt.Println(stack)
}

func main() {
	g := NewGraph(5)

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 5)
	g.addEdge(2, 4)
	g.addEdge(5, 3)

	fmt.Println(g.getGraph())

	g.topologicalSort()
}
