package main

import "fmt"

func countSort(input string) [26]int {
	arr := [26]int{}
	for _, ch := range input {
		arr[ch-'a']++
	}
	return arr
}

func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[[26]int][]string) // забыл добавить ключевое слово map
    for _, str := range strs {
        key := [26]int{} // key := [26]int
        for _, ch := range str {
            key[ch-'a']++ // panic: key[ch-'0']++
        }
        anagrams[key] = append(anagrams[key], str)
    }
    res := make([][]string, 0, len(anagrams)) // res := make([][]string, len(anagrams))
    for _, anagram := range anagrams {
        res = append(res, anagram)
    }
    return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
