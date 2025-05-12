// k = 2; map: A - 1, B - 2,
// s[l] == B
//       | |
// s = "ABABBA"

// k = 1; map: A - 1, B - 1,
// s[l] == B maxCount == 4 curCount == 1
//        |  |
// s = "AABABBA"

// k = 1; map: A - 11, B - 1
// s[l] == A maxCount == 12 curCount == 11
//         |        |
// s = "AABAAAAAAAAABBBB"
// 36 min

func traversalMap(countMap map[byte]int, ignorCh byte) int {
    res := 0
    for k, v := range countMap {
        if k != ignorCh {
            res += v
        }
    }
    return res
}

func characterReplacement(s string, k int) int {
    countMap := make(map[byte]int, 0)
    l, r := 0, -1
    curCount := 0
    res := 0
    prevLeft := byte(' ')
    for l < len(s) {
       for r + 1 < len(s) {
            if s[r+1] == s[l] && curCount < k {
                countMap[s[r+1]]++
                curCount++
                r++
            }
       }
       countMap[s[l]]--
       curCount--
       if countMap[s[l]] == 0 {
            delete(countMap, s[l])
       }
       if prevLeft != s[l] && l != 0 {
            curCount = traversalMap(countMap, s[l])
       }
       if r - l + 1  > res{
            res = r - l + 1
       }
       prevLeft = s[l]
       l++
    }
    return res
}
