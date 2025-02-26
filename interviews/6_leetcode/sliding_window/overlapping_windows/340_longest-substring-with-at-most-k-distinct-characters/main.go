package main

import "fmt"

func longestGeneSequence(gene string, k int) int {
	countMap := make(map[rune]int, k)
	l, r := 0, 0
	for l < len(gene) {
		for r < len(gene) && gene[r] == gene[r+1] {
			r++
		}
		l = r + 1
		r = r + 1
	}
	// for _, ch := range gene {

	// }
	return 0
}

func main() {
	gene := "YYxxXXXyyy"
	k := 3
	fmt.Println(longestGeneSequence(gene, k))
}
