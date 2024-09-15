package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура для обозначения процесса
type Process struct {
	time         int
	dependencies []int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение количества процессов
	n, _ := strconv.Atoi(readLine(reader))

	// Массив для хранения процессов
	processes := make([]Process, n)
	indegree := make([]int, n)

	// Заполнение данных о процессах
	for i := 0; i < n; i++ {
		line := strings.Split(readLine(reader), " ")
		time, _ := strconv.Atoi(line[0])
		var dependencies []int
		for _, dep := range line[1:] {
			if dep != "" {
				depIdx, _ := strconv.Atoi(dep)
				dependencies = append(dependencies, depIdx-1)
				indegree[depIdx-1]++
			}
		}
		processes[i] = Process{time: time, dependencies: dependencies}
	}
	fmt.Println(processes)
	fmt.Println()
	fmt.Println(indegree)
	// Очередь для топологической сортировки
	queue := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Массив для хранения времени завершения каждого процесса
	finishTime := make([]int, n)

	// Топологическая сортировка и вычисление времени выполнения
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Обновление времени завершения для каждого зависимого процесса
		for _, dep := range processes[current].dependencies {
			if finishTime[dep] < finishTime[current]+processes[dep].time {
				finishTime[dep] = finishTime[current] + processes[dep].time
			}
			indegree[dep]--
			if indegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}
	fmt.Println(indegree)
	// Нахождение максимального времени завершения
	maxTime := 0
	for i := 0; i < n; i++ {
		if maxTime < finishTime[i]+processes[i].time {
			maxTime = finishTime[i] + processes[i].time
		}
	}

	fmt.Println(maxTime)
}

// Функция для чтения строки
func readLine(reader *bufio.Reader) string {
	line, _, _ := reader.ReadLine()
	return string(line)
}

// 5
// 10 2 3 5
// 5 4
// 0
// 4
// 15 3

// 30

// 6
// 2 2
// 2 3
// 15 4
// 1 5
// 2 6
// 0

// 32
