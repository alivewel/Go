```
func containsMap(m map[rune]int, ch rune) bool {
    _, found := m[ch]
    return found 
}

func compareMap(s, p map[rune]int) bool {
    for k, v1 := range p {
        if v2, found := s[k]; !found || v2 != v1 {
            return false
        }
    }
    return true 
}

func findAnagrams(s string, p string) []int {
    l, r := 0, -1
    pMap := make(map[rune]int)
    for _, ch := range p {
        pMap[ch]++
    }
    sMap := make(map[rune]int)
    res := make([]int, 0)
    // for r + 1 < len(s) { // 
    for l < len(s) {
        for r + 1 < len(s) && containsMap(pMap, rune(s[r + 1])) && r - l + 1 != len(p) {
            sMap[rune(s[r + 1])]++
            r++
        }
        if r - l + 1 == len(p) && compareMap(sMap, pMap) {
            res = append(res, l)
        }
        sMap[rune(s[l])]--
        if count, found := sMap[rune(s[l])]; count < 0 && found { // count == 0 && found 
            delete(sMap, rune(s[l]))
            r++
        }
        l = l + 1
    }
    return res
}
```

1) Написал почти идеальное решение. Долго дебажил на мок-собесе и не мог найти причину, почему не запускается.
Оказалось, что проблема в месте, когда я удаляю символ из мапы с накопенной подстрокой. Ранее я удалял символ, когда частота символов в мапе была равна 0 и менее, оказалось, что нужно удалять, когда количество символов равна -1 и менее.

Имеется:
if count, found := sMap[rune(s[l])]; count <= 0 && found 
Должно быть:
if count, found := sMap[rune(s[l])]; count < 0 && found 

2) 
Имеется:
for r + 1 < len(s)
Должно быть: 
for l < len(s) 