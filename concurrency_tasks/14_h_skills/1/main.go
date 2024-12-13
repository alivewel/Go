package main

import "log"

func main() {
	m := map[interface{}]int{
		1:   1,
		"a": 2,
	}
	log.Println(m)

	m = map[interface{}]int{
		[3]int{1, 2, 3}: 1,
		// []int{1, 2, 3}:  1,
	}
	log.Println(m)
}
