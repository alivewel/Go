package main

import (
	"fmt"
)

// func findSubstring(s, sub string) int {
// 	n, m := len(s), len(sub)
// 	for i := 0; i <= n-m; i++ { // Первый указатель двигается по строке
// 		match := true
// 		for j := 0; j < m; j++ { // Второй указатель проверяет совпадение символов
// 			if s[i+j] != sub[j] {
// 				match = false
// 				break
// 			}
// 		}
// 		if match {
// 			return i // Возвращаем индекс начала подстроки
// 		}
// 	}
// 	return -1 // Если подстрока не найдена
// }

// func findSubstring(s, sub string) int {
// 	n, m := len(s), len(sub)
// 	// result := make([]rune, 0, n)
// 	skipIndex := make(map[int]struct{}, m)
// 	for i := 0; i <= n-m; i++ { // Первый указатель двигается по строке
// 		match := true
// 		for j := 0; j < m; j++ { // Второй указатель проверяет совпадение символов
// 			if s[i+j] != sub[j] {
// 				match = false
// 				break
// 			}
// 		}
// 		if match {
// 			// return i // Возвращаем индекс начала подстроки
// 			// result = append(result, rune(s[i]))
// 			// result = append(result, rune(s[i+1]))
// 			skipIndex[i] = struct{}{}
// 			skipIndex[i+1] = struct{}{}
// 			skipIndex[i+2] = struct{}{}
// 		}
// 	}
// 	fmt.Println(skipIndex)
// 	return -1 // Если подстрока не найдена
// }


func removeOccurrences(s, part string) string {
	n, m := len(s), len(part)
	slow, fast := 0, 0
	result := make([]byte, n)

	for fast < n {
		result[slow] = s[fast]
		slow++
		fast++

		// Проверяем, совпадает ли конец текущего результата с подстрокой
		if slow >= m && string(result[slow-m:slow]) == part {
			slow -= m // Удаляем подстроку, сдвигая указатель slow
		}
	}

	return string(result[:slow])
}

func main() {
	string := "daabcbaabcbc"
	substring := "abc"
	fmt.Println(removeOccurrences(string, substring))
}
