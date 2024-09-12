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

type TeamResult struct {
	TeamName      string              // имя команды
	Position      int                 // место команды
	HackedServers map[string]struct{} // взломанные сервера (чтобы избежать дубликатов, потом вывводить len)
	NumberPoints  string              // количество очков
}

// func parseTime(input string) (time.Time, error) {
// 	// Парсим время в формате hh:mm:ss
// 	return time.Parse("15:04:05", input)
// }

func parseTime2(timeStr string) time.Time {
	t, _ := time.Parse("15:04:05", timeStr)
	return t
}

func main() {
	/*
		// Чтение времени начала хакатона
		var startTimeStr string
		fmt.Scan(&startTimeStr)
		startTime, err := parseTime(startTimeStr)
		if err != nil {
			fmt.Println("Ошибка парсинга времени:", err)
			return
		}

		// Чтение количества запросов
		var n int
		fmt.Scan(&n)

		// Чтение запросов
		requests := make([]Request, n)
		for i := 0; i < n; i++ {
			var teamName, timeStr, serverID, result string
			fmt.Scan(&teamName, &timeStr, &serverID, &result)

			// Убираем кавычки из названия команды
			teamName = teamName[1 : len(teamName)-1]

			timeRequest, err := parseTime(timeStr)
			if err != nil {
				fmt.Println("Ошибка парсинга времени:", err)
				return
			}

			requests[i] = Request{
				TeamName: teamName,
				Time:     timeRequest,
				ServerID: serverID,
				Result:   result,
			}
		}

		fmt.Println(startTime)
	*/
	requests := []Request{
		{"T", parseTime2("00:10:21"), "A", "FORBIDDEN"},
		{"T", parseTime2("00:30:22"), "A", "DENIED"},
		{"YA", parseTime2("00:00:01"), "A", "ACCESSED"},
		{"VK", parseTime2("00:30:23"), "A", "ACCESSED"},
		{"YA", parseTime2("00:40:23"), "A", "ACCESSED"},
	}

	// Пример вывода парсенных данных
	for _, req := range requests {
		fmt.Printf("Команда: %s, Время: %s, Сервер: %s, Результат: %s\n",
			req.TeamName, req.Time.Format("15:04:05"), req.ServerID, req.Result)
	}

	// teamResult := make([]TeamResult, 0)
	teamResults := make(map[string]TeamResult)
	for _, req := range requests {
		teamResults[req.TeamName] = TeamResult{
			TeamName:      req.TeamName,
			HackedServers: make(map[string]struct{}),
		}
		if req.Result == "ACCESSED" {
			teamResults[req.TeamName].HackedServers[req.ServerID] = struct{}{}
		}
	}

	for _, teamResult := range teamResults {
		fmt.Printf("Команда: %s %s\n", teamResult.TeamName, teamResult.HackedServers)
	}
}

// зачем нужна информация о названии сервера?

// нужно как-то провалидировать время взлома, чтобы не выходить за пределы ограничения в 24 часа.
// если взлом сервера произошел после истечения 24 часов, решение не должно приниматься.
// но нам не поступает на вход дата, только время. по сути эти логи с отметкой времени.
// единственный сп
// у нас может быть ситуация при которой начало старта было не в 0:00:00, а со смещением, например 01:00:00 и даже 5:24:03.
// мы подразумеваем, что у нас события, которые поступают на вход, отосортированы по времени в формате неубывания.
// мы можем понять что у нас прос
// чтобы понять

// нужно взять время старта, отнять от нее 1 секунду - это будет время истечения.
// если время старта у нас равно 0:00:00, то нам не нужно отслеживать переход в другой день.
// если нет, нам нужно отследить переход и если он произошел, то сравнивать время с end_time
// так же предусмотреть случай, когда у нас первое время меньше start_date.
// это значит, что мы перешли в другой день. тогда нужно сразу сравнивать текущее время с end_time.
