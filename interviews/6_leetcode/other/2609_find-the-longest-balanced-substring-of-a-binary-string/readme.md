1) Здесь очень хочется применить паттерн пересекающегося окна. Но решение можно сделать гораздо проще.
Пересекающееся окно здесь ложиться, но усложняет решение и натыкаешься на подводные камни.
Мораль: нужно хорошенько думать о способах решения. Можно попасться на подобную ловушку.

Это задача из собеса в Тинькофф.

2) Изучить решение, которое предложил Алексей.
```
func findTheLongestBalancedSubstring(s string) int {
    countZeros, countOnes := 0, 0
    maxCount := 0
    for _, ch := range s {
        if ch == '0' {
            if countOnes > 0 {
               curCountPair := min(countZeros, countOnes) * 2
               maxCount = max(curCountPair, maxCount)
               countZeros = 0
               countOnes = 0
            }
            countZeros++
        } else {
            countOnes++
        }
    }
    curCountPair := min(countZeros, countOnes) * 2
    maxCount = max(curCountPair, maxCount) 
    return maxCount
}
```