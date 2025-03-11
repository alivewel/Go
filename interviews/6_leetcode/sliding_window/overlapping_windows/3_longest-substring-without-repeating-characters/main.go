package main

import "fmt"

func containsMap(chars map[byte]struct{}, ch byte) bool {
	if _, found := chars[ch]; found {
		return true
	}
	return false
}

func longestGeneSequence(gene string) int {
    l, r := 0, -1 // -1
	chars := make(map[byte]struct{})
	maxCountCh := 0
	for l < len(gene) {
		for r + 1 < len(gene) && !containsMap(chars, gene[r+1]) { // containsMap(chars, gene[r]) 
			chars[gene[r+1]] = struct{}{} // struct{}
			r++
		}
		delete(chars, gene[l])
		fmt.Println(l, r)
		sizeWindow := r - l + 1
		if sizeWindow > maxCountCh {
			maxCountCh = sizeWindow
		}
		l++
	}
	return maxCountCh
}

func main() {
	gene := "yxyabcxyx" // Вывод: 5
	fmt.Println(longestGeneSequence(gene))
}
