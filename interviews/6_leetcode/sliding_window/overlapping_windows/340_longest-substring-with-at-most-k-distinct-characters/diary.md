1) 

Нужно начинать с r = -1. Понять почему так.

Левый указатель всегда двигается на 1 (l++). До этого было l = r + 1, с других задач взял этот способ движения указателя.



```
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
```