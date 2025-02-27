package main

import "fmt"

// func longestGeneSequence(gene string, k int) int {
// 	countMap := make(map[rune]int, k)
// 	l, r := 0, 0
// 	for l < len(gene) {
// 		// for r < len(gene) && gene[r] == gene[r+1] {
// 		// 	r++
// 		// }
// 		// смотрим наличие в мапе и размер мапы не больше 3
// 		// for len(countMap) < k || _, found := countMap[ch]; found {
// 		// 	r++
// 		// 	countMap[ch]++
// 		// }
// 		for len(countMap) < k || _, found := countMap[rune(gene[r])]; found {
//             countMap[rune(gene[r])]++
//             r++
//         }
// 		l = r + 1
// 		r = r + 1
// 	}
// 	// for _, ch := range gene {

// 	// }
// 	return 0
// }

func isMapContaine(m map[rune]int, key rune) bool {
    _, found := m[key]
    return found
}

func longestGeneSequence(gene string, k int) int {
    countMap := make(map[rune]int, k) // можно использовать byte
    l, r := 0, 0
    maxCount := 0
    for l < len(gene) {
        for r < len(gene) && (len(countMap) < k || isMapContaine(countMap, rune(gene[r]))) {
            countMap[rune(gene[r])]++ // тут нужен индекс r + 1, подумать почему
            r++
        }
        count := r - l + 1
        maxCount = max(count, maxCount)
        if countMap[rune(gene[l])] == 0 {
            delete(countMap, gene[l])
        }
        
        l = r + 1
    }
    return 0
}

func main() {
	gene := "YYxxXXXyyy"
	k := 3
	fmt.Println(longestGeneSequence(gene, k))
}
