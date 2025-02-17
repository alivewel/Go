package main

import "fmt"

func intersect(A []int , B []int ) []int {
	res := make([]int, 0, 0)
	p1, p2 := 0, 0
	for p1 < len(A) && p2 < len(B) {
		if A[p1] == B[p2] {
			res = append(res, A[p1])
			p1++
			p2++
		} else if A[p1] > B[p2] {
			p2++
		} else {
			p1++
		}
	}
	return res
}

func main() {
	nums1 := []int{-3, 2, 2, 5, 8, 19, 31}
	nums2 := []int{1, 2, 2, 2, 6, 19, 52}
	fmt.Println(intersect(nums1, nums2))
}
