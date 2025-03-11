package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
    l, r := 0, len(s) - 1
    for l < r {
        lCh := unicode.ToLower(rune(s[l]))
        rCh := unicode.ToLower(rune(s[r]))
        if !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l]))  {
            l++
            continue
        }
        if !unicode.IsLetter(rune(s[r])) && !unicode.IsDigit(rune(s[r])) {
            r--
            continue
        }
        if lCh != rCh {
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
