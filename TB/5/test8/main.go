package main

import (
	"fmt"
	"sort"
	"time"
)

type Request struct {
	TeamName        string
	Time            time.Time
	ServerID        string
	Result          string
	HackathonIsOver bool
}

type TeamResult struct {
	TeamName      string                // имя команды
	Position      int                   // место команды
	HackedServers map[string]ServerInfo // взломанные сервера (чтобы избежать дубликатов, потом вывводить len)
	NumberPoints  int                   // количество очков
	NumHackServ   int                   // количество взломанных серверов
	LastActivity  time.Time             // время последнего взлома
}

type ServerInfo struct {
	ServerIsHacked      bool
	FailedAttemptsCount int
}

type TimeInfo struct {
	PrevDiff  int
	HasNewDay bool
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("15:04:05", timeStr)
	return t
}

func checkAccessed(reqResult string) bool {
	return reqResult == "ACCESSED" || reqResult == "ACCESED" || reqResult == "ACESSED" || reqResult == "ACESED"
}

func checkForbidden(reqResult string) bool {
	return reqResult == "FORBIDEN" || reqResult == "FORBIDDEN" || reqResult == "DENIED"
}

func checkRequestsTimesOver(requests *[]Request, startTime time.Time) {
	startTimeIsMidnight := false
	var diff int
	// завести мапу с ключем название команды и значением prevDiff и hasNewDay
	// записывать и проверять prevDiff и hasNewDay только для своей команды
	if startTime.Format("15:04:05") == "00:00:00" {
		startTimeIsMidnight = true
	}

	timesInfo := make(map[string]TimeInfo)

	for i := range *requests {
		req := &(*requests)[i]
		intervalTime := req.Time

		timeInfo, exists := timesInfo[req.TeamName]
		if !exists {
			timeInfo = TimeInfo{}
		}

		if i == 0 && intervalTime.Before(startTime) && !startTimeIsMidnight {
			timeInfo.HasNewDay = true
		}

		diff = int(intervalTime.Sub(startTime).Seconds())

		if !timeInfo.HasNewDay && i != 0 && diff < timeInfo.PrevDiff && !startTimeIsMidnight {
			timeInfo.HasNewDay = true
		}

		if timeInfo.HasNewDay && !startTimeIsMidnight && intervalTime.After(startTime.Add(-1*time.Second)) {
			req.HackathonIsOver = true
		}

		if startTimeIsMidnight && diff < timeInfo.PrevDiff && int(timeInfo.PrevDiff) != 0 {
			req.HackathonIsOver = true
		}

		timeInfo.PrevDiff = diff
		timesInfo[req.TeamName] = timeInfo
	}
}

func calcDiffMinutes(date1, date2 time.Time) int {
	if date2.Before(date1) {
		date2 = date2.Add(24 * time.Hour)
	}
	// Вычисляем разницу между датами
	difference := date2.Sub(date1)

	// Получаем разницу в минутах
	return int(difference.Minutes())
}

// processTeamResults преобразует мапу в срез, сортирует и присваивает позиции
func processTeamResults(teamResults map[string]TeamResult) []TeamResult {
	// Преобразуем мапу в срез
	teams := make([]TeamResult, 0, len(teamResults))
	for _, teamResult := range teamResults {
		teams = append(teams, teamResult)
	}

	// Сортируем срез
	sort.Slice(teams, func(i, j int) bool {
		if len(teams[i].HackedServers) != len(teams[j].HackedServers) {
			return len(teams[i].HackedServers) < len(teams[j].HackedServers)
		}
		return teams[i].NumberPoints < teams[j].NumberPoints
	})

	// Присваиваем позицию
	var prevPosition int
	for i := range teams {
		if i != 0 && len(teams[i].HackedServers) == len(teams[i-1].HackedServers) && teams[i].NumberPoints == teams[i-1].NumberPoints {
			teams[i].Position = prevPosition
		} else {
			teams[i].Position = i + 1
			prevPosition = i + 1
		}
	}

	return teams
}

func main() {
	requests := []Request{
		{"VK", parseTime("00:10:21"), "A", "FORBIDDEN", false},
		{"T", parseTime("00:00:23"), "A", "DENIED", false},
		{"T", parseTime("00:20:23"), "A", "ACCESSED", false},
		{"VK", parseTime("00:30:23"), "A", "ACCESSED", false},
		{"YA", parseTime("00:40:23"), "B", "ACCESSED", false},
	}
	// requests := []Request{
	// 	{"T", parseTime("15:10:21"), "A", "FORBIDDEN", false},
	// 	{"T", parseTime("00:59:59"), "A", "DENIED", false},
	// 	{"YA", parseTime("01:00:00"), "A", "ACCESSED", false},
	// 	{"VK", parseTime("01:00:01"), "A", "ACCESSED", false},
	// 	{"YA", parseTime("00:40:23"), "A", "ACCESSED", false},
	// }

	startTimeStr := "00:00:00"
	startTime, _ := time.Parse("15:04:05", startTimeStr)

	checkRequestsTimesOver(&requests, startTime)
	// Print results
	// for _, req := range requests {
	// 	fmt.Printf("Request: %s at %s, HackathonIsOver: %v\n", req.TeamName, req.Time.Format("15:04:05"), req.HackathonIsOver)
	// }

	teamResults := make(map[string]TeamResult)
	for _, req := range requests {
		teamResult, exists := teamResults[req.TeamName]
		if !exists {
			teamResult = TeamResult{
				TeamName:      req.TeamName,
				HackedServers: make(map[string]ServerInfo),
				LastActivity:  startTime,
			}
		}

		if checkAccessed(req.Result) && !req.HackathonIsOver {
			// посчитать количество штрафных баллов
			// для этого нужна отдельная функция, которая считает разницу в минутах между двумя временами
			serverInfo := teamResult.HackedServers[req.ServerID]
			serverInfo.ServerIsHacked = true
			// fmt.Println(teamResult.LastActivity, req.Time, calcDiffMinutes(teamResult.LastActivity, req.Time))

			penaltyPoints := calcDiffMinutes(teamResult.LastActivity, req.Time)

			// прибавляем по 20 штрафных минут за каждую неудачную попытку входа, если сервер удалось взломать
			failedCount := teamResult.HackedServers[req.ServerID].FailedAttemptsCount

			teamResult.NumberPoints = penaltyPoints + failedCount*20
			teamResult.LastActivity = req.Time
			teamResult.NumHackServ++
			teamResult.HackedServers[req.ServerID] = serverInfo
		}
		if checkForbidden(req.Result) && !req.HackathonIsOver { //
			serverInfo := teamResult.HackedServers[req.ServerID]
			serverInfo.FailedAttemptsCount++
			teamResult.HackedServers[req.ServerID] = serverInfo
		}
		teamResults[req.TeamName] = teamResult
	}

	// for _, teamResult := range teamResults {
	// 	fmt.Printf("Команда: %s %v %v %v %v\n", teamResult.TeamName, teamResult.HackedServers, teamResult.LastActivity, teamResult.NumberPoints, teamResult.NumHackServ)
	// }

	// Обработка и вывод результатов
	teams := processTeamResults(teamResults)

	// Выводим результат
	for _, teamResult := range teams {
		fmt.Printf("%v \"%s\" %v %v\n", teamResult.Position, teamResult.TeamName, teamResult.NumHackServ, teamResult.NumberPoints)
	}
}

// краевой случай startTime == 23:55:00
// посчитать количество баллов
