package main

import (
	// "fmt"
	"strconv"
)

func compress(chars []rune) []rune {
	res := make([]rune, 0, len(chars))
	l, r := 0, 0
	for r < len(chars) {
		for r + 1 < len(chars) && chars[r] == chars[r+1] {
			r++
		}
		countChars := r - l + 1
		res = append(res, rune(chars[l]))
		if countChars > 1 && countChars < 10 {
			res = append(res, rune('0' + countChars))
		} else if countChars > 9 {
			strCountChars := strconv.Itoa(countChars)
			arrRuneStrCountChars := []rune(strCountChars)
			for _, ch := range arrRuneStrCountChars {
				res = append(res, ch)
			}

		}
		// если символ встречается 1 раз ничего не делаем
		l = r + 1
		r = r + 1
	}
    return res
}

func main() {
	// chars := []rune{'a','a','b','b','c','c','c'}
	chars := []rune{'y','y','x','x','x','x','x','x','x','x','x','x','x','x','y','y'}
	res := string(compress(chars))
	_ = res
	// fmt.Println(res)
}
