package main

import "fmt"

func calcFee(a, b, c, d int) int {
	if d-b > 0 {
		return a + c*(d-b)
	} else {
		return a
	}
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Println(calcFee(a, b, c, d))
}

// a = 100
// b = 10
// c = 12
// d = 1
