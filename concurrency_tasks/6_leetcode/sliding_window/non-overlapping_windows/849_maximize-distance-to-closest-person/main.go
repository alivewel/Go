package main

import "fmt"

func maxDistToClosest(seats []int) int {
	res := 0
	l, r := 0, 0
    for r < len(seats) {
		for r < len(seats) - 1 && seats[r] == seats[r+1] {
			r++
		}
		fmt.Println(l, r)
		if seats[l] == 0 {
			if l == 0 {
				res = r - l + 1
			} else if r == len(seats) - 1 {
				// curRes := 0
				// diff := r - l + 1
				// if diff % 2 != 0 {
				// 	curRes = (diff + 1) / 2
				// } else {
				// 	curRes = diff / 2 // при diff == 2 не сходится 
				// }
				
				// 3 - 0 + 1 / 2 == 2
				// 2 - 0 + 1 / 2 == 1
				// 1 - 0 + 1 / 2 == 1
				

				curRes := r - l + 1 // неправильно определил при первом запуске
				res = max(curRes, res)
			} else {
				curRes := r - l + 1 / 2
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
