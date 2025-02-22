package main

import "fmt"

func maxDistToClosest(seats []int) int {
	res := 0
	l, r := 0, 0
    for r < len(seats) {
		for r < len(seats) - 1 && seats[r] == seats[r+1] {
			r++
		}
		if seats[l] == 0 {
			if l == 0 || r == len(seats) - 1 { // если с левого или правого края свободные места 
				curRes = r - l + 1
				res = max(curRes, res)
			} else {
				curRes := r - l + 2 / 2 // идеальая формула
				// (4 - 0 + 2) / 2 == 3
				// 3 - 0 + 2 / 2 == 2
				// 2 - 0 + 2 / 2 == 2
				// 1 - 0 + 2 / 2 == 1
				// curRes := r - l + 1 / 2
				res = max(curRes, res)
			}
		}
		l = r + 1
		r = r + 1
	}
	return res
}

func main() {
	// seats := []int{1,0,0,0}
	// seats := []int{1,0,0,0,1,0,1}
	seats := []int{0,1}
	         //   [1,0,0,0,1,0,1]
	fmt.Println(maxDistToClosest(seats))
}

// нужно посчитать группы нулей
// может быть 3 случая:
// в самом начале - от левого бортика (l == 0),
// в середине, здесь нужно определить количество
// и посчитать по формуле для четных и нечетных чисел
// в самом конце - от правого бортика r == len(seats)
//
