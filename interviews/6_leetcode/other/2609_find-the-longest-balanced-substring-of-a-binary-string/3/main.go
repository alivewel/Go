// это решение намного проще остальных
func findTheLongestBalancedSubstring(s string) int {
    maxCount := 0
    countZeros, countOnes := 0, 0
    for i, num := range s {
        if num == '0' {
            countZeros++
        }
        if num == '1' {
            countOnes++
        }
        if i == len(s) - 1 || (s[i] == '1' && s[i+1] == '0') {
            sizeBalStr := min(countZeros, countOnes) * 2
            if sizeBalStr > maxCount {
                maxCount = sizeBalStr
            }
            countZeros = 0
            countOnes = 0
        }
    }

    return maxCount
}