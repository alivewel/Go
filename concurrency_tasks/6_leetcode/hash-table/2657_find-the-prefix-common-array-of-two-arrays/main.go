package main

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	frequency := make([]int, len(A)+1)
	result := make([]int, len(A))
	counter := 0
	for i := range A {
		frequency[A[i]]++
		frequency[B[i]]++
		if A[i] == B[i] {
			counter++
			result[i] = counter
			continue
		}
		if frequency[A[i]] == 2 {
			counter++
		}
		if frequency[B[i]] == 2 {
			counter++
		}
		result[i] = counter
	}
	return result
}

func main() {
	A := []int{1,3,2,4}
	B := []int{3,1,2,4}
	fmt.Println(findThePrefixCommonArray(A, B))
}
