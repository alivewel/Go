package main

import "fmt"

// заводим 1 мапу, в которой 
// символы из строки s инкрементируем счетчик
// символы из строки t декрементируем счетчик
// не сразу вспомнил оптимальное решение, хотел через 2 мапы решить

// сначала создал map с ключом типа byte
// но оказывается for range по строке возвращает rune

// Интервьювер не примет ваше решение, так как оно не оптимально.
// 1 за время, так как оно оптимально. 0 за память, так как можно 
// использовать фиксированный размер массива для алфавита и 
// достичь O(1) пространственной сложности.

func isValidAnagram(s string, t string) bool {
    countLetters := make(map[rune]int)
    for _, ch := range s {
        countLetters[ch]++
    }
    for _, ch := range t {
        countLetters[ch]--
    }
    for _, count := range countLetters {
        if count != 0 {
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

	fmt.Println(isValidAnagram(s, t))
}
