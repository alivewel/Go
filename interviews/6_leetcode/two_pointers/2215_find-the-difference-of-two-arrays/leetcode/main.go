package main

import (
	"fmt"
	"sort"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1, set2 := make(map[int]struct{}), make(map[int]struct{})
	for _, val := range nums1 {
		set1[val] = struct{}{}
	}
	for _, val := range nums2 {
		set2[val] = struct{}{}
	}
	diff1 := []int{} // Элементы, которые есть в nums1, но не в nums2
    diff2 := []int{} 
	for key, _ := range set1 {
		if _, found := set2[key]; !found {
            diff1 = append(diff1, key)
		}
	}
	for key, _ := range set2 {
		if _, found := set1[key]; !found {
            diff2 = append(diff2, key)
		}
	}
	// Сортируем оба среза
    sort.Ints(diff1)
    sort.Ints(diff2)

    // Добавляем diff1 и diff2 в result
    result := make([][]int, 0)
    result = append(result, diff1)
    result = append(result, diff2)
	return result
}

func main() {
	nums1 := []int{1, 5, 7, 9}
	nums2 := []int{2, 3, 5, 6, 7, 8}
	fmt.Println(findDifference(nums1, nums2))
}