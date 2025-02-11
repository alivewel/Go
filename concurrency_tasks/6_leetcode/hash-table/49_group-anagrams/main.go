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
	countSortMap := make(map[[26]int][]string)
	for _, str := range strs {
		key := countSort(str)
		countSortMap[key] = append(countSortMap[key], str)
	}
	// можем заранее выделить память для массива зная размер мапы
	// make - len задаем 0, cap - размер мапы
	// иначе создадуться пустые ячейки, вот так:
	// [[] [] [] [eat tea ate] [tan nat] [bat]]
	result := make([][]string, 0, len(countSortMap))
	// значение мапы - []string, можем просто аппендить в [][]string
	for _, group := range countSortMap {
		result = append(result, group)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
