package main

import "fmt"

func canPermutePalindrome(str string) bool {
	countSort := make([]int, 26)
	for _, s := range str {
		countSort[s-'a']++
	}
	cnt := 0
	for _, count := range countSort {
		if count % 2 != 0 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

func main() {
	str1 := "carerac" // true
	str2 := "code"    // false

	fmt.Println(canPermutePalindrome(str1))
	fmt.Println(canPermutePalindrome(str2))
}
