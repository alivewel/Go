
func removeDuplicateLetters(s string) string {
    countSort := [26]int{}
    for _, ch := range s {
        countSort[ch-'a']++
    }
    res := make([]byte, 0)
    for i, count := range countSort {
        if count != 0 {
            res = append(res, 'a' + byte(i))
        }
    }
    return string(res)
}
