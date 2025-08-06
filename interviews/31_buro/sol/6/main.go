package main

import (
    "fmt"
)

// func main() {
//     a := []int{1, 2, 3}
//     b := []int{1, 2, 3}
//     fmt.Println(a == b)
//     // invalid operation: a == b (slice can only be compared to nil)
// }

func main() {
    a := [3]int{1, 2, 3}
    b := [3]int{1, 2, 3}
    fmt.Println(a == b) // true
}