package main

import "fmt"

func main(){
    a := make([]int, 0, 4)
    a = append(a, 1,2,3)
    appendSlice(a)
    fmt.Println(a[:]) // [1 2 3]
    fmt.Println(a[:4]) // [1 2 3 4]
}

func appendSlice(a []int) {
    a = append(a, 4)
}
