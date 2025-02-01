package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			res = append(res, nums1[p1]) // не путать последовательность
			p1++ // сначала добавить, потом сдвинуть указатели
			p2++
		}
	}
	return res
}

func main() {
	nums1 := []int{-3, 2, 2, 5, 8, 19, 31}
	nums2 := []int{1, 2, 2, 2, 6, 19, 52}
	fmt.Println(intersect(nums1, nums2))
}
