package main

import "fmt"

// фунция для заполнения мапы
// для первого запуска по другому нужно заполнять
func fillMapFirst(strMap map[rune]int, str string) {
	for _, ch := range str {
		strMap[ch]++
	}
}

// фунция для заполнения мапы
func fillMap(strMap map[rune]int, addCh, removeCh rune) {
	strMap[addCh]++
	strMap[removeCh]--
}

// фунция для сравнения мап
func compareMap(mapStrP, mapStrS map[rune]int) bool {
	for k1, v1 := range mapStrP {
		if v2, found := mapStrS[k1]; !found || v2 < v1 {
			return false
		} 
	}
	return true
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	mapStrP := make(map[rune]int, len(p))
    for _, ch := range p {
		mapStrP[ch]++
	}
	mapStrS := make(map[rune]int, len(p))
	for i := 0; i < len(s) - len(p) + 1; i++ {
		if i == 0 {
			fillMapFirst(mapStrS, s[:len(p)])
		} else {
			fillMap(mapStrS, rune(s[i+len(p)-1]), rune(s[i-1]))
		}
		if compareMap(mapStrP, mapStrS) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	// s := "badc"
	// t := "baba"

	// s := "anagram"
	// t := "nagaram"

	s := "cbaebabacd"
	p := "abc"

	// s := "abab"
	// p := "ab"

	fmt.Println(findAnagrams(s, p))
}
