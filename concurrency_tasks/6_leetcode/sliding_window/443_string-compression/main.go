package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	res := make([]byte, 0, len(chars))
	l, r := 0, 0
	for r < len(chars) {
		for r + 1 < len(chars) && chars[r] == chars[r+1] {
			r++
		}
		countChars := r - l + 1
		res = append(res, byte(chars[l]))
		if countChars > 1 && countChars < 10 {
			fmt.Println(countChars)
			// res = append(res, byte(countChars))
			res = append(res, strconv.Itoa(countChars)...)
		} else if countChars > 9 {
			for countChars != 0 {
				// res = append(res, byte(countChars % 10))
				res = append(res, strconv.Itoa(countChars % 10)...)
				countChars /= 10
			} 
		}
		// если символ встречается 1 раз ничего не делаем
		l = r + 1
		r = r + 1
	}
	// chars = res // временное решение
	copy(chars, res)
	fmt.Println(string(res))
    return len(res)
}

func main() {
	chars := []byte{'a','a','b','b','c','c','c'}
	fmt.Println(compress(chars), string(chars))
}
