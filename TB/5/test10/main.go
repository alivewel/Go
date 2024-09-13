package main

import (
	"fmt"
	"sort"
	"time"
)

type ServerInfo struct {
	Active bool
	Count  int
}

type TeamResult struct {
	TeamName      string
	Position      int
	HackedServers map[string]ServerInfo
	NumberPoints  int
	NumHackServ   int // количество взломанных серверов
	LastActivity  time.Time
}

func main() {
	teamResults := make(map[string]TeamResult)

	// Пример данных
	teamResults["VK"] = TeamResult{
		TeamName: "VK",
		HackedServers: map[string]ServerInfo{
			"A": {true, 1},
		},
		LastActivity: time.Date(1, 1, 1, 0, 30, 23, 0, time.UTC),
		NumberPoints: 50,
		NumHackServ:  1,
	}

	teamResults["T"] = TeamResult{
		TeamName: "T",
		HackedServers: map[string]ServerInfo{
			"A": {true, 1},
		},
		LastActivity: time.Date(1, 1, 1, 0, 20, 23, 0, time.UTC),
		NumberPoints: 40,
		NumHackServ:  1,
	}

	teamResults["YA"] = TeamResult{
		TeamName: "YA",
		HackedServers: map[string]ServerInfo{
			"B": {true, 0},
		},
		LastActivity: time.Date(1, 1, 1, 0, 40, 23, 0, time.UTC),
		NumberPoints: 40,
		NumHackServ:  1,
	}

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
		if i != 0 && len(teams[i].HackedServers) == len(teams[i-1].HackedServers) && teams[i].NumberPoints == teams[i-i].NumberPoints {
			teams[i].Position = prevPosition
		} else {
			teams[i].Position = i + 1
			prevPosition = i + 1
		}
		// teams[i].Position = i + 1
	}

	// Выводим результат
	for _, teamResult := range teams {
		fmt.Printf("%v \"%s\" %v %v\n", teamResult.Position, teamResult.TeamName, teamResult.NumHackServ, teamResult.NumberPoints)
	}
}
