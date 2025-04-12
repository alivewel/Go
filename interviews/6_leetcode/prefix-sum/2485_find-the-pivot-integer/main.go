// 1 2 3 4
// prefSum == allSum - allSum + i
// 1 + 2 == 10 - 3 + i
func pivotInteger(n int) int {
    allSum := 0
    for i := 1; i <= n; i++ { // i < n
        allSum += i
    }
    prefSum := 0
    for i := 1; i <= n; i++ { // i < n
        prefSum += i
        if prefSum == allSum - prefSum + i {
            return i
        }
    }
    return -1
}
