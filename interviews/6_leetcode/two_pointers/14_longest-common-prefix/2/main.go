
func longestCommonPrefix(strs []string) string {
    length := len(strs)
    index := 0
    var rune prevCh
    // for i, str := strs {
    for {
        for i := 0; i < length; i++ {
            if index < len(str) {
                break
            }
            if i == 0 {
                prevCh = strs[i][index]
            }
            if prevCh != curCh {
				break
			}
            if i == length - 1 {
                index++
            } 
        } 
    }    
}
