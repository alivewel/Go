package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Запрос команды
type Request struct {
	TeamName string
	Time     time.Time
	ServerID string
	Result   string
}

func parseTime(input string) (time.Time, error) {
	// Парсим время в формате hh:mm:ss
	return time.Parse("15:04:05", input)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Чтение времени начала хакатона
	scanner.Scan()
	// startTime, err := parseTime(scanner.Text())
	// if err != nil {
	// 	fmt.Println("Ошибка парсинга времени:", err)
	// 	return
	// }

	// Чтение количества запросов
	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)

	// Чтение запросов
	requests := make([]Request, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()

		// Парсим строку запроса
		parts := strings.SplitN(line, " ", 4)
		if len(parts) != 4 {
			fmt.Println("Неверный формат строки:", line)
			return
		}

		teamName := strings.Trim(parts[0], "\"")
		timeRequest, err := parseTime(parts[1])
		if err != nil {
			fmt.Println("Ошибка парсинга времени:", err)
			return
		}
		serverID := parts[2]
		result := parts[3]

		requests[i] = Request{
			TeamName: teamName,
			Time:     timeRequest,
			ServerID: serverID,
			Result:   result,
		}
	}

	// Пример вывода парсенных данных
	for _, req := range requests {
		fmt.Printf("Команда: %s, Время: %s, Сервер: %s, Результат: %s\n",
			req.TeamName, req.Time.Format("15:04:05"), req.ServerID, req.Result)
	}
}

// 00:00:00
// 5
// "T" 00:10:21 A FORBIDDEN
// "T" 00:30:22 A DENIED
// "YA" 00:00:01 A ACCESSED
// "VK" 00:30:23 A ACCESSED
// "YA" 00:40:23 A ACCESSED

// 1 "T" 1 40
// 1 "YA" 1 40
// 3 "VK" 1 50
