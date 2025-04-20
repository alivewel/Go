package main

import "fmt"

func bruteForce(n, balance int, curComb []rune, allComb *[]string) {
    // fmt.Println(n, len(curComb))
    if balance < 0 || balance > n {
        return
    }
    if len(curComb) == n * 2 { // balance == 0 && 
        *allComb = append(*allComb, string(curComb))
        return
    }
    brackets := []struct{
		brack rune
        count int
	}{
		{
			brack: '(',
			count: 1,
		},
		{
			brack: ')',
			count: -1,
		},
	}
    for _, bracket := range brackets {
        curComb = append(curComb, bracket.brack)
        bruteForce(n, balance + bracket.count, curComb, allComb)
        curComb = curComb[:len(curComb)-1] // забыл про это
    }
}

func generateBrackets(n int) []string {
    res := make([]string, 0)
    bruteForce(n, 0, []rune{}, &res)
    return res
}

func main() {
	fmt.Println(generateBrackets(3))
}

// stack overflow
