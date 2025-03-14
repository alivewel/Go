package main

import "fmt"

func mySqrt(x int) int {
    l, r := 0, x + 1 // l, r := 0, x
    res := 0
    // 1 2 3 4 5 6 7 8
    for r - l > 1 {
        // (r + l) / 2 = (8 + 0) / 2 = 4
        // (r + l) / 2 = (4 + 0) / 2 = 2
        // (r + l) / 2 = (4 + 2) / 2 = 3
        // (r + l) / 2 = (3 + 2) / 2 = 2
        middle := (r + l) / 2
        if middle * middle <= x {
            l = middle
            res = middle
        } else {
            r = middle
        }
    }
    return res
}

func main() {
	fmt.Println(mySqrt(8)) // Expected 2
}
