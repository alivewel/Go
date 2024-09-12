package main

import (
	"fmt"
	"time"
)

func main() {
	// Парсим строки времени в объекты time.Time
	layout := "2006-01-02 15:04:05 -0700"
	date1, err1 := time.Parse(layout, "0000-01-01 23:30:23 +0000")
	date2, err2 := time.Parse(layout, "0000-01-01 00:30:23 +0000")

	fmt.Println(date1, date2)
	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка при парсинге даты:", err1, err2)
		return
	}

	fmt.Printf("Разница в минутах: %d\n", calcDiffMinutes(date1, date2))
}

func calcDiffMinutes(date1, date2 time.Time) int {
	if date2.Before(date1) {
		date2 = date2.Add(24 * time.Hour)
	}
	fmt.Println(date1, date2)
	// Вычисляем разницу между датами
	difference := date2.Sub(date1)

	// Получаем разницу в минутах
	return int(difference.Minutes())
}
