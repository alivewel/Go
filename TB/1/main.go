package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// input := "1-5,6,8-9,11"
	// input := "1-1,11,6,8-9"
	var input string
	fmt.Scan(&input)
	intervals := strings.Split(input, ",")
	result := []int{}

	for _, interval := range intervals {
		if strings.Contains(interval, "-") {
			bounds := strings.Split(interval, "-")
			start, _ := strconv.Atoi(bounds[0])
			end, _ := strconv.Atoi(bounds[1])
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			num, _ := strconv.Atoi(interval)
			result = append(result, num)
		}
	}
	sort.Ints(result)

	for i, num := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
		if i == len(result)-1 {
			fmt.Print("\n")
		}
	}
}

// 1. На вход нам всегда подается отсортированный диапозон чисел. О сортировке думать не нужно.
// 2. Можно использовать библиотеки типа strconv и strings.
// 3. Валидировать строку скорее всего не нужно.
// 4. Покрыть код тестами.
// 5. Принимать диапозон чисел через fmt.Scan.
// 6. Добавить сортировку в первом.
