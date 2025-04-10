package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func isMapContaine(m map[rune]int, key rune) bool {
    _, found := m[key]
    return found
}

func longestGeneSequence(gene string, k int) int {
    countMap := make(map[rune]int, k) // можно использовать byte
    // l, r := 0, 0
    l, r := 0, -1
    maxCount := 0
    for l < len(gene) {
        for r + 1 < len(gene) && (len(countMap) < k || isMapContaine(countMap, rune(gene[r+1]))) {
            countMap[rune(gene[r+1])]++ // тут нужен индекс r + 1, подумать почему
            r++
        }
        count := r - l + 1
        maxCount = max(count, maxCount)

        countMap[rune(gene[l])]--
        // if countMap[rune(gene[l])] == 0 {
        if countMap[rune(gene[l])] <= 0 {
            delete(countMap, rune(gene[l]))
        }
        
        l++
    }
    return maxCount
}

func main() {
	gene := "YYxxXXXyyy"
	k := 3
	fmt.Println(longestGeneSequence(gene, k))
}
