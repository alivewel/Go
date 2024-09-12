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

	for _, intervalStr := range timeIntervals {
		intervalTime, _ := time.Parse("15:04:05", intervalStr)

		// Вычисляем разницу в секундах
		diff := intervalTime.Sub(startTime).Seconds()

		// Если разница отрицательная, добавляем 24 часа
		if diff < 0 {
			diff += dayInSeconds
		}

		// Проверяем, превышает ли разница 24 часа
		if diff >= dayInSeconds {
			fmt.Printf("Время %s выходит за 24-часовой интервал.\n", intervalStr)
		} else {
			fmt.Printf("Время %s в пределах 24-часового интервала.\n", intervalStr)
		}
	}
}
