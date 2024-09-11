package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input := "1-5,6,8-9,11"
	input := "1-5,11,6,8-9"
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

	for i, num := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}

// посчитать количество составных чисел от l до r,
// количество делителей которых при этом является простым числом

// мне нужно посчитать простые числа используя решето эратосфена
// функция для расчета решета эратосфена принимает количество простых чиел
// а мне нужно посчитать указав правую границу диапозона, то есть максимальные число
// что мне задать в функции решета эратосфена?