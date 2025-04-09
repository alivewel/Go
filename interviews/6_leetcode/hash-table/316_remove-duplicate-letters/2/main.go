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

func containsMap(m map[byte]struct{}, ch byte) bool {
    _, found := m[ch]
    return found
}
