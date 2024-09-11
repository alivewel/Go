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

func main() {
	// Создаем массив запросов
	requests := []Request{
		{"T", parseTime("00:10:21"), "A", "FORBIDDEN"},
		{"T", parseTime("00:30:22"), "A", "DENIED"},
		{"YA", parseTime("00:00:01"), "A", "ACCESSED"},
		{"VK", parseTime("00:30:23"), "A", "ACCESSED"},
		{"YA", parseTime("00:40:23"), "A", "ACCESSED"},
	}

	// Пример вывода данных
	for _, req := range requests {
		fmt.Printf("Команда: %s, Время: %s, Сервер: %s, Результат: %s\n",
			req.TeamName, req.Time.Format("15:04:05"), req.ServerID, req.Result)
	}
}

// Функция для парсинга времени
func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("15:04:05", timeStr)
	return t
}
