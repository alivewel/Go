package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	index := 0
	prevCh := 'a'
	flagExit := false
	for {
		for i := range strs {
			if index > len(strs[i]) - 1 {
				flagExit = true
				break
			}
			curCh := rune(strs[i][index])
			if i == 0 {
				prevCh = rune(strs[i][index])
				fmt.Println(prevCh)
			}
			if prevCh != curCh {
				flagExit = true
				break
			}
		}
		if flagExit {
			fmt.Println("1")
			break
		}
		index++
	}
	return strs[0][0:index]
}

func main() {
	// strs := []string{"flower","flow","flight"}
	strs := []string{"a"}
	fmt.Println(longestCommonPrefix(strs)) // "fl"
}
