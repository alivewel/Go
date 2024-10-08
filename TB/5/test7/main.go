package main

import (
	"fmt"
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
	NumberPoints  string                // количество очков
}

type ServerInfo struct {
	ServerIsHacked      bool
	FailedAttemptsCount int
}

func main() {
	requests := []Request{
		{"VK", parseTime("00:10:21"), "A", "FORBIDDEN", false},
		{"T", parseTime("00:00:23"), "A", "DENIED", false},
		{"T", parseTime("00:23:23"), "A", "ACCESSED", false},
		{"VK", parseTime("00:30:23"), "A", "ACCESSED", false},
		{"YA", parseTime("00:40:23"), "B", "ACCESSED", false},
	}

	startTimeStr := "01:00:00"
	startTime, _ := time.Parse("15:04:05", startTimeStr)

	checkRequestsTimesOver(&requests, startTime)

	teamResults := make(map[string]TeamResult)
	for _, req := range requests {
		// Проверяем, существует ли уже команда
		teamResult, exists := teamResults[req.TeamName]
		if !exists {
			teamResult = TeamResult{
				TeamName:      req.TeamName,
				HackedServers: make(map[string]ServerInfo),
			}
		}

		if checkAccessed(req.Result) && !req.HackathonIsOver {
			serverInfo := teamResult.HackedServers[req.ServerID]
			serverInfo.ServerIsHacked = true
			teamResult.HackedServers[req.ServerID] = serverInfo
		}

		if checkForbidden(req.Result) && !req.HackathonIsOver {
			serverInfo := teamResult.HackedServers[req.ServerID]
			serverInfo.FailedAttemptsCount++
			teamResult.HackedServers[req.ServerID] = serverInfo
		}

		// Обновляем карту с измененной информацией о команде
		teamResults[req.TeamName] = teamResult
	}

	for _, teamResult := range teamResults {
		fmt.Printf("Команда: %s %v\n", teamResult.TeamName, teamResult.HackedServers)
	}
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("15:04:05", timeStr)
	return t
}

func checkAccessed(reqResult string) bool {
	return reqResult == "ACCESSED" || reqResult == "ACCESED" || reqResult == "ACESSED" || reqResult == "ACESED"
}

func checkForbidden(reqResult string) bool {
	return reqResult == "FORBIDEN" || reqResult == "FORBIDDEN"
}

func checkRequestsTimesOver(requests *[]Request, startTime time.Time) {
	hasNewDay, startTimeIsMidnight := false, false
	var diff, prevDiff float64

	if startTime.Format("15:04:05") == "00:00:00" {
		startTimeIsMidnight = true
	}

	for i := range *requests {
		req := &(*requests)[i]
		intervalTime := req.Time

		if i == 0 && intervalTime.Before(startTime) && !startTimeIsMidnight {
			hasNewDay = true
		}

		diff = intervalTime.Sub(startTime).Seconds()

		if !hasNewDay && i != 0 && diff < prevDiff && !startTimeIsMidnight {
			hasNewDay = true
		}

		if hasNewDay && !startTimeIsMidnight && intervalTime.After(startTime.Add(-1*time.Second)) {
			req.HackathonIsOver = true
		}

		if startTimeIsMidnight && diff < prevDiff {
			req.HackathonIsOver = true
		}

		prevDiff = diff
	}
}
