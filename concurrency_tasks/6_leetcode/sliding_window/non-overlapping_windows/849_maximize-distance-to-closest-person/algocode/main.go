package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func bestParkingSpot(spots []int) int {
	res := 0
	l, r := 0, 0
    for r < len(spots) {
		for r < len(spots) - 1 && spots[r] == spots[r+1] {
			r++
		}
		if spots[l] == 0 {
			if l == 0 || r == len(spots) - 1 { // если с левого или правого края свободные места 
				curRes := r - l + 1
				res = max(curRes, res)
			} else {
				fmt.Println(r, l)
				curRes := (r - l + 2) / 2 // идеальная формула
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
	
	fmt.Println(bestParkingSpot(spots))
}
