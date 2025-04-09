```
func removeDuplicateLetters(s string) string {
    countSort := [26]int{}
    for _, ch := range s {
        countSort[ch-'a']++
    }
    res := make([]byte, 0)
    for i, count := range countSort {
        if count != 0 {
            res = append(res, 'a' + i)
        }
    }
    return string(res)
}
```

1) 
res = append(res, 'a' + i)
Line 10: Char 31: cannot use 'a' + i (value of type int) as byte value in argument to append (solution.go)

2) Иправил ошибку. Но оказывается неправильно решил, решение не прошло. Я думал нужно в отсортированном порядке. По факту нет. Нужноучитывать порядок из входной строки.
Input s = "cbacdcbc"
Output "abcd"
Expected "acdb"

```
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
```