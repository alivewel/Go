// abc
// cbaebabacd
// p

// t O(n)
// m O(n+m)

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

// cbaebabacd"

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
        fmt.Println(rune(s[l]), r - l + 1, sMap, pMap)
        if r - l + 1 == len(p) && compareMap(sMap, pMap) {
            res = append(res, l)
        }
        sMap[rune(s[l])]--
        if count, found := sMap[rune(s[l])]; count <= 0 && found { // count == 0 && found 
            delete(sMap, rune(s[l]))
            r++
        }
        l = l + 1
    }
    return res
}














// func fillMapFirst(strMap map[rune]int, str string) {
// 	for _, ch := range str {
// 		strMap[ch]++
// 	}
// }

// // фунция для заполнения мапы
// func fillMap(strMap map[rune]int, addCh, removeCh rune) {
// 	strMap[addCh]++
// 	strMap[removeCh]--
// }

// // фунция для сравнения мап
// func compareMap(mapStrP, mapStrS map[rune]int) bool {
// 	for k1, v1 := range mapStrP {
// 		if v2, found := mapStrS[k1]; !found || v2 < v1 {
// 			return false
// 		} 
// 	}
// 	return true
// }

// func findAnagrams(s string, p string) []int {
// 	res := make([]int, 0)
// 	mapStrP := make(map[rune]int, len(p))
//     for _, ch := range p {
// 		mapStrP[ch]++
// 	}
// 	mapStrS := make(map[rune]int, len(p))
// 	for i := 0; i < len(s) - len(p)+1; i++ {
// 		if i == 0 {
// 			fillMapFirst(mapStrS, s[:len(p)])
// 		} else {
// 			fillMap(mapStrS, rune(s[i+len(p)-1]), rune(s[i-1]))
// 		}
// 		if compareMap(mapStrP, mapStrS) {
// 			res = append(res, i)
// 		}
// 	}
// 	return res
// }
