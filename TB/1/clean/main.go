package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
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
