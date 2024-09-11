package main

import (
	"fmt"
	"time"
)

// Структура запроса
type Request struct {
	TeamName string
	Time     time.Time
	ServerID string
	Result   string
}

// Структура результата команды
type TeamResult struct {
	TeamName       string
	Position       int
	HackedServers1 map[string]struct{}
	NumberPoints   string
}

func main() {
	// Пример данных
	requests := []Request{
		{"TeamA", time.Now(), "A", "ACCESSED"},
		{"TeamB", time.Now(), "B", "DENIED"},
		{"TeamA", time.Now(), "C", "ACCESSED"},
		{"TeamC", time.Now(), "A", "FORBIDDEN"},
	}

	// Карта для отслеживания уникальных команд
	uniqueTeams := make(map[string]TeamResult)

	// Проход по запросам
	for _, req := range requests {
		if _, exists := uniqueTeams[req.TeamName]; !exists {
			// Если команда еще не добавлена, добавляем ее
			uniqueTeams[req.TeamName] = TeamResult{
				TeamName:       req.TeamName,
				HackedServers1: make(map[string]struct{}),
			}
		}
	}

	// Пример вывода уникальных команд
	for _, teamResult := range uniqueTeams {
		fmt.Printf("Команда: %s\n", teamResult.TeamName)
	}
}
