// stack = [a, c, d]
//          |
// s = "cbacdcbc"

func removeDuplicateLetters(s string) string {
    countSort := [26]int{}
    for _, ch := range s {
        countSort[ch-'a']++
    }
    res := make([]rune, 0)
    containsRes := make(map[rune]struct{})
    for _, ch := range s {
        countSort[byte(ch)-'a']--
        if _, found := containsRes[ch]; found {
            continue
        }
        for len(res) != 0 && ch < res[len(res)-1] && countSort[res[len(res)-1]-'a'] > 0 {
            delete(containsRes, res[len(res)-1])
            res = res[:len(res)-1] // очень важна последовательность удаления, если поменять местами эту и предыдущую строку, будет паника
        }
        res = append(res, ch)
        containsRes[ch] = struct{}{} // забыл про это
    }
    return string(res)
}

func removeDuplicateLetters2(s string) string {
    countSort := [26]int{}
    for _, ch := range s {
        countSort[ch-'a']++
    }
    res := make([]rune, 0)
    containsRes := make(map[rune]struct{})
    for _, ch := range s {
        countSort[ch-'a']--
        if _, found := containsRes[ch]; found {
            continue
        }
        // Удаляем из стека символы, которые:
        // 1. Больше текущего (ch < last)
        // 2. Ещё встретятся в строке (countSort[last] > 0)
        for len(res) != 0 && ch < res[len(res)-1] && countSort[res[len(res)-1]-'a'] > 0 {
            last := res[len(res)-1]
            delete(containsRes, last)
            res = res[:len(res)-1]
        }
        res = append(res, ch)
        containsRes[ch] = struct{}{}
    }
    return string(res)
}

func containsMap(m map[byte]struct{}, ch byte) bool {
    _, found := m[ch]
    return found
}
