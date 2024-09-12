package main

import (
	"fmt"
	"time"
)

func main() {
	// Время старта
	startTimeStr := "00:00:00"
	startTime, _ := time.Parse("15:04:05", startTimeStr)

	// Временные отрезки
	timeIntervals := []string{"01:10:21", "02:10:21", "06:10:21", "15:10:21", "01:10:21", "00:10:21"}

	// Преобразуем строки в объекты time.Time
	var parsedTimes []time.Time
	for _, t := range timeIntervals {
		parsedTime, _ := time.Parse("15:04:05", t)
		parsedTimes = append(parsedTimes, parsedTime)
	}

	// Проверяем переход за 00:00:00
	for i := 1; i < len(parsedTimes); i++ {
		if parsedTimes[i].Before(parsedTimes[i-1]) {
			fmt.Printf("Переход за 00:00:00 обнаружен между %s и %s\n", parsedTimes[i-1].Format("15:04:05"), parsedTimes[i].Format("15:04:05"))
		}
	}
}
