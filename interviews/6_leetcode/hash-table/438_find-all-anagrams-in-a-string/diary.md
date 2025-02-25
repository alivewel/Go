1) Второй тестовый случай не прошел

Input: s = "abab", p = "ab"
Output: [0,1]
Expected: [0,1,2]

Проблема была в условии выхода из цикла. Добавил +1 итерацию.
Имеется:
for i := 0; i < len(s) - len(p); i++ 
Нужно:
for i := 0; i < len(s) - len(p) + 1; i++ 

```
func fillMapFirst(strMap map[rune]int, str string) {
	for _, ch := range str {
		strMap[ch]++
	}
}

// фунция для заполнения мапы
func fillMap(strMap map[rune]int, addCh, removeCh rune) {
	strMap[addCh]++
	strMap[removeCh]--
}

// фунция для сравнения мап
func compareMap(mapStrP, mapStrS map[rune]int) bool {
	for k1, v1 := range mapStrP {
		if v2, found := mapStrS[k1]; !found || v2 < v1 {
			return false
		} 
	}
	return true
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	mapStrP := make(map[rune]int, len(p))
    for _, ch := range p {
		mapStrP[ch]++
	}
	mapStrS := make(map[rune]int, len(p))
	for i := 0; i < len(s) - len(p); i++ {
		if i == 0 {
			fillMapFirst(mapStrS, s[:len(p)])
		} else {
			fillMap(mapStrS, rune(s[i+len(p)-1]), rune(s[i-1]))
		}
		if compareMap(mapStrP, mapStrS) {
			res = append(res, i)
		}
	}
	return res
}
```