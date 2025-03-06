package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    index := 0
    var prevCh byte // var prevCh rune // var rune prevCh
    flagExit := false
    for {
        for i := 0; i < len(strs); i++ {
            if index >= len(strs[i]) {
                flagExit = true // flagExit := true
                break
            }
            curCh := strs[i][index]
            if i == 0 {
                prevCh = curCh
            }
            if prevCh != curCh {
                flagExit = true // flagExit := true
				break
			} 
        }
        if flagExit {
            break
        }
        index++
    }
    return strs[0][:index]
}

func main() {
    strs := []string{"flower","flow","flight"}
    fmt.Println(longestCommonPrefix(strs))
}