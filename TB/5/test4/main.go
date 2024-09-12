package main

import (
	"fmt"
	"time"
)

func main() {
	// Время старта
	startTimeStr := "01:00:00"
	startTime, _ := time.Parse("15:04:05", startTimeStr)

	// Временные отрезки
	timeIntervals := []string{"01:10:21", "02:10:21", "06:10:21", "15:10:21", "01:10:21", "00:10:21"}

	// Константа для 24 часов в секундах
	const dayInSeconds = 24 * 60 * 60
	hasNewDay := false

	var diff, prevDiff float64

	for i, intervalStr := range timeIntervals {
		intervalTime, _ := time.Parse("15:04:05", intervalStr)

		if i == 0 && intervalTime.Before(startTime) {
			hasNewDay = true
			fmt.Printf("Первое время с переходом!\n")
		}

		diff = intervalTime.Sub(startTime).Seconds()

		if !hasNewDay && i != 0 && diff < prevDiff {
			fmt.Printf("Переход на новый день!\n")
			hasNewDay = true
		}

		// Проверяем, превышает ли разница 24 часа
		if diff >= dayInSeconds {
			fmt.Printf("Время %s выходит за 24-часовой интервал. %v %v %v\n", intervalStr, diff, prevDiff, diff < prevDiff)
		} else {
			fmt.Printf("Время %s в пределах 24-часового интервала. %v %v %v\n", intervalStr, diff, prevDiff, diff < prevDiff)
		}

		// Вычисляем разницу в секундах
		prevDiff = diff
	}
}
