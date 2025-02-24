package main

import "fmt"

func main() {
	 sl := []*int{}
	for _, val := range []int{1, 2, 3, 4, 5} { 
		sl = append(sl, &val)
	}
	for _, val := range sl {
	fmt.Printf("%d ", *val)
	}
}