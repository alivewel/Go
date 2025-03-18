//        |
// "1111101000001110"
// пропускаем первые единицы
// ищем первый нуль
// считаем количество нулей
// учитываем момент перехода от нулей к единицам
// считаем количество единиц
// как только встречаем новый нуль, выходим из окна.
// плавающее окно непересек.
// if nums[l] == 0
// for nums[l] == 0

//     |  
// "01000111"
// нужно ли занулять countZeros countOnes

// "01"
// "00"

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




// func findTheLongestBalancedSubstring(s string) int {
//     l, r := 0, 0
//     maxCount := 0
//     if len(s) < 2 {
//         return 0
//     }
//     for r < len(s) {
//         if s[l] == '0' {
//             countZeros, countOnes := 1, 0 // 1, 1
//             for r + 1 < len(s) && s[r+1] == '0' {
//                 r++
//                 countZeros++
//             }
//             // countZeros := r - l + 1
//             // l = r
//             for r + 1 < len(s) && s[r+1] == '1' {
//                 r++
//                 countOnes++
//             }
//             // countOnes := r - l + 1
//             // if countZeros > 0 && countOnes > 0 {
//                 curCountPair := min(countZeros, countOnes) * 2
//                 maxCount = max(curCountPair, maxCount)
//             // }
//         }
//         l = r + 1
//         r = r + 1
//     }
//     return maxCount
// }