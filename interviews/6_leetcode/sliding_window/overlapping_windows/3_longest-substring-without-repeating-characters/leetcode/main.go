package main

import "fmt"

func containsMap(chars map[byte]struct{}, ch byte) bool {
	if _, found := chars[ch]; found {
		return true
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
    l, r := 0, -1 // -1
	chars := make(map[byte]struct{})
	maxCountCh := 0
	for l < len(s) {
		for r + 1 < len(s) && !containsMap(chars, s[r+1]) { // containsMap(chars, s[r]) 
			chars[s[r+1]] = struct{}{} // struct{}
			r++
		}
		delete(chars, s[l])
		sizeWindow := r - l + 1
		if sizeWindow > maxCountCh {
			maxCountCh = sizeWindow
		}
		l++
	}
	return maxCountCh
}

func main() {
	s := "yxyabcxyx" // Вывод: 5
	fmt.Println(lengthOfLongestSubstring(s))
}
