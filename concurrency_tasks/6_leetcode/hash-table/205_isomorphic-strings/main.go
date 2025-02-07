package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sMap := make(map[rune]rune)
	tMap := make(map[rune]rune)
	for i := range s {
		sRune := rune(s[i])
		tRune := rune(t[i])

		k1, found := sMap[sRune]
		if found && k1 != tRune {
			return false
		}
		k2, found := tMap[tRune]
		if found && k2 != sRune {
			return false
		}

		sMap[sRune] = tRune
		sMap[tRune] = sRune
	}

	return true
}

func main() {
	s := "ada"
	t := "bcb"
	// s := "egg"
	// t := "add"

	fmt.Println(isIsomorphic(s, t))
}
