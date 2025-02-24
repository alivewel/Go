package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sMap := make(map[rune]rune)
	tMap := make(map[rune]rune)
	for i := range s {
		fmt.Println(i)
		sRune := rune(s[i])
		tRune := rune(t[i])

		if k1, found := sMap[sRune]; found && k1 != tRune {
			return false
		}
		if k2, found := tMap[tRune]; found && k2 != sRune {
			return false
		}

		sMap[sRune] = tRune
		tMap[tRune] = sRune
	}

	return true
}

func main() {
	// s := "ada"
	// t := "bcb"
	// s := "egg"
	// t := "add"
	s := "badc"
	t := "baba"

	fmt.Println(isIsomorphic(s, t))
}
