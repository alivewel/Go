package main

import (
	"fmt"
	"time"
)

func main() {
	// Время старта
	startTimeStr := "01:00:00"
	startTime, _ := time.Parse("15:04:05", startTimeStr)

	// Временные отрезки в формате time.Time
	timeIntervals := []time.Time{
		parseTime("01:10:21"),
		parseTime("02:10:21"),
		parseTime("06:10:21"),
		parseTime("15:10:21"),
		parseTime("00:59:59"),
		parseTime("01:00:00"),
		parseTime("01:00:01"),
		parseTime("01:10:21"),
		parseTime("00:10:21"),
	}

	checkTimeIntervals(startTime, timeIntervals)
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("15:04:05", timeStr)
	return t
}

func checkTimeIntervals(startTime time.Time, timeIntervals []time.Time) {
	hasNewDay, startTimeIsMidnight, hackathonIsOver := false, false, false
	var diff, prevDiff float64

	if startTime.Format("15:04:05") == "00:00:00" {
		startTimeIsMidnight = true
	}

	for i, intervalTime := range timeIntervals {
		if i == 0 && intervalTime.Before(startTime) && !startTimeIsMidnight {
			hasNewDay = true
		}

		diff = intervalTime.Sub(startTime).Seconds()

		if !hasNewDay && i != 0 && diff < prevDiff && !startTimeIsMidnight {
			hasNewDay = true
		}

		if hasNewDay && !startTimeIsMidnight && intervalTime.After(startTime.Add(-1*time.Second)) {
			hackathonIsOver = true
		}

		if startTimeIsMidnight && diff < prevDiff {
			hackathonIsOver = true
		}

		// Проверяем, превышает ли разница 24 часа
		if hackathonIsOver {
			fmt.Printf("Время %s выходит за 24-часовой интервал.\n", intervalTime.Format("15:04:05"))
		} else {
			fmt.Printf("Время %s в пределах 24-часового интервала.\n", intervalTime.Format("15:04:05"))
		}

		// Вычисляем разницу в секундах
		prevDiff = diff
	}
}
