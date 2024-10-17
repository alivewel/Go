package main

import (
	"fmt"
)

// какой будет порядок элементов в первой и втором принте
func main() {
	groceries := map[string]string{
		"x": "10",
		"b": "20",
		"c": "30",
	}
	for idx, item := range groceries {
		fmt.Println(idx, item)
	}
	fmt.Println(groceries)
}
