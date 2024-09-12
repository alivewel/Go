package main

import (
	"fmt"
	"time"
)

func main() {
	// // Время старта
	// startTimeStr := "00:00:00"
	// startTime, _ := time.Parse("15:04:05", startTimeStr)

	// // Временные отрезки
	// timeIntervals := []string{"01:10:21", "02:10:21", "06:10:21", "15:10:21", "23:59:59", "00:00:00", "00:00:01", "01:10:21", "00:10:21"}

	// Время старта
	startTimeStr := "01:00:00"
	startTime, _ := time.Parse("15:04:05", startTimeStr)

	// Временные отрезки
	timeIntervals := []string{"01:10:21", "02:10:21", "06:10:21", "15:10:21", "00:59:59", "01:00:00", "01:00:01", "01:10:21", "00:10:21"}

	hasNewDay, startTimeIsMidnight, hackathonIsOver := false, false, false

	var diff, prevDiff float64

	if startTimeStr == "00:00:00" {
		startTimeIsMidnight = true
	}

	for i, intervalStr := range timeIntervals {
		intervalTime, _ := time.Parse("15:04:05", intervalStr)

		if i == 0 && intervalTime.Before(startTime) && !startTimeIsMidnight {
			hasNewDay = true
			fmt.Printf("Первое время с переходом!\n")
		}

		diff = intervalTime.Sub(startTime).Seconds()

		if !hasNewDay && i != 0 && diff < prevDiff && !startTimeIsMidnight {
			fmt.Printf("Переход на новый день!\n")
			hasNewDay = true
		}
		// fmt.Println(intervalTime, hasNewDay, !startTimeIsMidnight && intervalTime.Before(startTime))
		if hasNewDay && !startTimeIsMidnight && intervalTime.After(startTime.Add(-1 * time.Second)) {
			hackathonIsOver = true
		}

		if startTimeIsMidnight && diff < prevDiff {
			hackathonIsOver = true
		}

		// Проверяем, превышает ли разница 24 часа
		if hackathonIsOver {
			fmt.Printf("Время %s выходит за 24-часовой интервал. %v %v %v\n", intervalStr, diff, prevDiff, diff < prevDiff)
		} else {
			fmt.Printf("Время %s в пределах 24-часового интервала. %v %v %v\n", intervalStr, diff, prevDiff, diff < prevDiff)
		}

		// Вычисляем разницу в секундах
		prevDiff = diff
	}
}

// в случае если у нас стартовое время 00:00:00
// то как только мы перешли за новый день, значит мы вышли за 24-часовой интервал.

// иначе
// как только мы перешли за новый день
// мы должны сравнивать текущую дату со (старт дат - 1 секунда)

// Миша участвует в специ