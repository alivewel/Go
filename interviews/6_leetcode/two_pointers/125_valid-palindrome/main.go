package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])) {
			l++
		}
		for l < r && !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	// s := "AB_CD_D_C_B_A"
	// s := "amanaplanacanalpanama"
	s := "A man, a plan, a canal: Panama"
	res := isPalindrome(s)
	fmt.Println(res)
}
