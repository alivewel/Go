package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched() // попробуйте закомментировать
	}
	printFinishWork(workerNum)
}

const goroutinesNum = 3

func main() {
	worketInput := make(chan string, 2)
	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, worketInput)
	}
	months := []string{"Январь", "Февраль", "Март",
		"Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь",
		"Октябрь", "Ноябрь", "Декабрь",
	}
	for _, monthName := range months {
		worketInput <- monthName
	}
	close(worketInput) // попробуйте закомментировать
	time.Sleep(time.Millisecond)
}

func formatWork(workerNum int, input string) string {
	return fmt.Sprintf("%s %s\n", strings.Repeat(" ", workerNum), input) // Используем Sprintf для форматирования
}

func printFinishWork(workerNum int) {
	fmt.Printf("Worker %d finished\n", workerNum) // Более информативное сообщение
}

// func formatWork(in int, input string) string {
// 	return fmt.Sprintln(strings.Repeat(" ", in), " ", input)
// }

// func printFinishWork(in int) {
// 	fmt.Printf("%d finish\n", in)
// }
