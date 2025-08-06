package main

import "fmt"

func main(){
	s1 := []int{1, 2, 3}     // s1: [1 2 3]
	// s1 := make([]int, 3, 4) // cap > len
	s2 := s1[1:]             // s2: [2 3] (слайс ссылается на те же данные, что и s1)
	s2 = append(s2, 4)       // s2: [2 3 4] — возможно, выделена новая память!
	s2[0] = 10               // изменяется первый элемент s2 — но влияет ли это на s1?

    fmt.Println(s1) // [1 2 3]
    fmt.Println(s2) // [10 3 4]
	// fmt.Println(s1) // [0 10 0]
    // fmt.Println(s2) // [10 0 4]
}
