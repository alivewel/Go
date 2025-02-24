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
				res = max(curRes, res)
			}
		}
		l = r + 1
		r = r + 1
	}
	return res
}

func main() {
	// spots := []int{1,0,0,0}
	// spots := []int{0,1}
	spots := []int{1,0,0,0,1,0,1} // res = 2
	fmt.Println(maxDistToClosest(seats))
}
