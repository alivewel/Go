package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// Map of digits to corresponding letters
	lettersDigits := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	// Backtracking function
	var backtrack func(index int, combination string)
	backtrack = func(index int, combination string) {
		// Base case: if the combination length matches the input digits
		if index == len(digits) {
			result = append(result, combination)
			return
		}

		// Get the letters corresponding to the current digit
		letters := lettersDigits[digits[index]]

		// Explore each letter for the current digit
		for _, letter := range letters {
			backtrack(index+1, combination+string(letter))
		}
	}

	// Start backtracking from the first digit
	backtrack(0, "")

	return result
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
