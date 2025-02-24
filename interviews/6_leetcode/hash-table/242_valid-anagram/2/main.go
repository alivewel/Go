package main

import "fmt"

func isAnagram(s string, t string) bool {
	countChar := make([]int, 26)
	for _, c := range s {
		countChar[c-'a']++
	}
	for _, c := range t {
		countChar[c-'a']--
	}
	for _, c := range countChar {
		if c != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "badc"
	t := "baba"

	// s := "anagram"
	// t := "nagaram"

	fmt.Println(isAnagram(s, t))
}
