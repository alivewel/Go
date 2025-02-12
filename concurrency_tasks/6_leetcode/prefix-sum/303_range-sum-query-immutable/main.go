package main

import "fmt"

type NumArray struct {
	arr []int
	px  []int
}

func Constructor(nums []int) NumArray {
	px := make([]int, len(nums)+1)
	for i, num := range nums {
		px[i+1] += px[i] + num
	}
	return NumArray{
		arr: nums,
		px:  px,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.px[right+1] - this.px[left]
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{-5, 0, 6, 3, -2, 1}
	
	px1 := Constructor(arr1)
	px2 := Constructor(arr2)

	sum1 := px1.SumRange(0, 1)
	sum2 := px1.SumRange(2, 4)

	fmt.Println(px1.px)
	fmt.Println(px2.px)

	fmt.Println(sum1)
	fmt.Println(sum2)
}
