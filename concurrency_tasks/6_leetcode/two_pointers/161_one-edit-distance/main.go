package main

import "fmt"

func isOneEditDistance(s string, t string) bool {
	diffLen := abs(len(s) - len(t))
	if diffLen > 1 {
		return false
	}

	countDiff := 0
	largerStr := getLongerString(s, t)
	
	p1, p2 := 0, 0
	for p1 < len(s) && p2 < len(t) {
		if s[p1] != t[p2] {
			if largerStr == 1 {
				p1++
			} else if largerStr == 2 {
				p2++
			} else {
				p1++
				p2++
			}
			countDiff++
		} else {
			p1++
			p2++
		}
	}
	// обработать случай, когда не до конца какой-то строки дошли
	if p1 < len(s) || p2 < len(t) {
		countDiff++
	}
	return countDiff == 1
}

func getLongerString(s string, t string) int {
	if len(s) == len(t) {
		return 0
	} else if len(s) > len(t) {
		return 1
	}
	return 2
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

// func main() {
// 	// s := "ab"
// 	// t := "acb"

// 	s := "cab" 
// 	t := "ad"

// 	fmt.Println(isOneEditDistance(s, t))
// }

func main() {
    tests := []struct {
        s        string
        t        string
        expected bool
    }{
        // Тесты с одной операцией редактирования
        {"ab", "acb", true},   // Вставка символа
        {"cab", "ad", false},  // Более одной операции
        {"1203", "1213", true}, // Замена символа
        {"a", "", true},       // Удаление символа
        {"", "a", true},       // Вставка символа
        {"abc", "ab", true},   // Удаление символа
        {"ab", "abc", true},   // Вставка символа

        // Тесты без операций редактирования
        {"abc", "abc", false}, // Строки одинаковые
        {"", "", false},       // Пустые строки

        // Тесты с более чем одной операцией редактирования
        {"abc", "def", false}, // Замена более одного символа
        {"abc", "abcd", true}, // Вставка одного символа
        {"abc", "abde", false}, // Вставка и замена

        // Граничные случаи
        {"a", "b", true},      // Замена одного символа
        {"", "ab", false},     // Вставка двух символов
        {"ab", "", false},     // Удаление двух символов
        {"a", "a", false},     // Одинаковые строки длиной 1
        {"a", "aa", true},     // Вставка одного символа
        {"aa", "a", true},     // Удаление одного символа
    }

    for i, test := range tests {
        result := isOneEditDistance(test.s, test.t)
        fmt.Printf("Test %d: isOneEditDistance(%q, %q) = %v (expected %v)\n",
            i+1, test.s, test.t, result, test.expected)
    }
}
