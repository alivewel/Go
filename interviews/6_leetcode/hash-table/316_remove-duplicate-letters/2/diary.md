// stack = [a, c, d]
//          |
// s = "cbacdcbc"

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

1) 

Line 14: Char 36: cannot use ch (variable of type rune) as byte value in map index (solution.go)
Line 17: Char 35: invalid operation: ch < res[len(ch) - 1] (mismatched types rune and byte) (solution.go)
Line 17: Char 43: invalid argument: ch (variable of type rune) for built-in len (solution.go)
Line 17: Char 78: cannot use ch (variable of type rune) as byte value in argument to containsMap (solution.go)
Line 18: Char 28: invalid argument: ch (variable of type rune) for built-in len (solution.go)
Line 19: Char 33: cannot use ch (variable of type rune) as byte value in argument to delete (solution.go)
Line 21: Char 27: cannot use ch (variable of type rune) as byte value in argument to append (solution.go)