package main

import "fmt"

func isIsomorphic(s string, t string) bool {
    map1 := make(map[byte]byte)
    map2 := make(map[byte]byte)
    // p1, p2 := 0, 0
    // p := 0
    for i := 0; i < len(s); i++ {
        // проверяем соответствие в 1 строке
        if val, found := map1[s[i]]; found && val != t[i] {
            return false
        }
        // проверяем соответствие в 2 строке
        if val, found := map2[t[i]]; found && val != s[i] {
            return false
        }
        map1[s[i]] = t[i]
        map2[t[i]] = s[i]
        //s[i] byte получаем
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
