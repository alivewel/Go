package main

import "fmt"

// фунция для заполнеия мапы
func fillMap(countSort map[rune]int, str string) {
	
}

// фунция для сравнения мап

func findAnagrams(s string, p string) []int {
	mapStrP := make(map[rune]int, len(p))
    for ch := range p {
		mapStrP[ch]++
	}
	mapStrS := make(map[rune]int, len(p))
	for ch := range s {
		mapStrP[ch]++
	} 
}

func main() {
	// s := "badc"
	// t := "baba"

	// s := "anagram"
	// t := "nagaram"

	s := "cbaebabacd"
	p := "abc"

	fmt.Println(findAnagrams(s, t))
}
