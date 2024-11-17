package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	jsonData := `{"name": "John", "age": 30}`
	reader := strings.NewReader(jsonData)

	var result map[string]interface{}
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
