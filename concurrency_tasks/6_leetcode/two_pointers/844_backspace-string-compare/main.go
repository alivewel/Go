package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	p1, p2 := len(s) - 1, len(t) - 1
	countRemoveS, countRemoveT := 0, 0
	for p1 >= 0 && p2 >= 0 { // подумать
		if s[p1] == t[p2] && s[p1] != '#' && t[p2] != '#' && countRemoveS == 0 && countRemoveT == 0 {
			p1--
			p2--
		} else if s[p1] == '#' {
			countRemoveS++
			p1--
		} else if t[p2] == '#' {
			countRemoveT++
			p2--
		} else if countRemoveS > 0 {
			countRemoveS--
			p1--
			// p1--
		} else if countRemoveT > 0 {
			countRemoveT--
			p2--
			// p2--
		} else if s[p1] != t[p2] && countRemoveS == 0 && countRemoveT == 0 {
			fmt.Println(countRemoveS, countRemoveT)
			fmt.Println(s[p1], t[p2])
			return false
		} 
	}
	if p1 != -1 || p2 != -1 {
		fmt.Println(p1, p2)
		return false
	}
    return true
}

func main() {
	// s := "ab#c"
	// t := "ad#c"

	s := "ab##" 
	t := "c#d#"
	// t := "aab#c"
	fmt.Println(backspaceCompare(s, t))
}
