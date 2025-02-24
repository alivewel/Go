package main

import (
	"encoding/json"
	"fmt"
)

type Model struct {
	Name string `json:"test"`
	Age  int    `json:"age"`
}

func main() {
	data := `{"name":"Alex", "age":30}`
	var model Model
	err := json.Unmarshal([]byte(data), &model)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(model)
}

// Вывод:
// { 30}
