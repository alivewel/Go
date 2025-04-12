1) Много ошибок при работе с byte, for range работает при работе со строкой возвращает rune, а не byte.

Line 14: Char 36: cannot use ch (variable of type rune) as byte value in map index (solution.go)
Line 17: Char 35: invalid operation: ch < res[len(ch) - 1] (mismatched types rune and byte) (solution.go)
Line 17: Char 43: invalid argument: ch (variable of type rune) for built-in len (solution.go)
Line 17: Char 78: cannot use ch (variable of type rune) as byte value in argument to containsMap (solution.go)
Line 18: Char 28: invalid argument: ch (variable of type rune) for built-in len (solution.go)
Line 19: Char 33: cannot use ch (variable of type rune) as byte value in argument to delete (solution.go)
Line 21: Char 27: cannot use ch (variable of type rune) as byte value in argument to append (solution.go)


func removeDuplicateLetters(s string) string {
    countSort := [26]int{}
    for _, ch := range s {
        countSort[ch-'a']++
    }
    res := make([]byte, 0) // стек для хранения результата
    containsRes := make(map[byte]struct{})
    for _, ch := range s {
        countSort[ch-'a']--
        if _, found := containsRes[bytech]; found {
            continue
        }
        for len(res) != 0 && ch < res[len(ch)-1] && containsMap(containsRes, ch) && countSort[ch-'a'] > 0 {
            res = res[:len(ch)-1]
            delete(containsRes, ch)
        }
        res = append(res, ch)
    }
    return string(res)
}

func containsMap(m map[byte]struct{}, ch byte) bool {
    _, found := m[ch]
    return found
}



2) Много логических ошибок.

Имеется:
delete(containsRes, ch)  // Удаляем текущий символ ch, а нужно удалить последний из стека!
Должно быть:
delete(containsRes, res[len(res)-1])

Имеется:
for len(res) != 0 && ch < res[len(res)-1] && countSort[byte(ch)-'a'] > 0
// Здесь я проверяю, что countSort[ch] > 0, но нужно проверять оставшееся количество последнего символа в стеке (countSort[last]), а не текущего ch.
Должно быть:
last := res[len(res)-1]
for len(res) != 0 && ch < last && countSort[last-'a'] > 0 

delete(containsRes, res[len(res)-1])
res = res[:len(res)-1] // очень важна последовательность удаления, если поменять местами эту и предыдущую строку, будет паника

Имеется:
countSort[byte(ch)-'a'] > 0 
Должно быть:
countSort[res[len(res)-1]-'a'] > 0

func removeDuplicateLetters(s string) string {
    countSort := [26]int{}
    for _, ch := range s {
        countSort[ch-'a']++
    }
    res := make([]byte, 0) // стек для хранения результата
    containsRes := make(map[byte]struct{})
    for _, ch := range s {
        countSort[ch-'a']--
        if _, found := containsRes[byte(ch)]; found {
            continue
        }
        for len(res) != 0 && byte(ch) < res[len(res)-1] && containsMap(containsRes, byte(ch)) && countSort[byte(ch)-'a'] > 0 {
            res = res[:len(res)-1]
            delete(containsRes, byte(ch))
        }
        res = append(res, byte(ch))
    }
    return string(res)
}
