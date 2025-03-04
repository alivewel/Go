package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
    lettersDigits := map[int]string{
        2: "abc",
        3: "def",
        4: "ghi",
		5: "jkl",
        6: "mno",
        7: "pqrs",
        8: "tuv",
        9: "wxyz",
    }
    queue := []string{""}
    for _, val := range digits {
		letters := lettersDigits[int(val-'0')]
        newQueue := []string{}
        for _, combination := range queue {
            for _, letter := range letters {
                newQueue = append(newQueue, combination + string(letter))
            }
        }
        queue = newQueue
    }
    return queue
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
