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
	arr := []int{1, 2, 3, 4}
	px := Constructor(arr)
	sum := px.SumRange(0, 1)
	fmt.Println(px.px)
	fmt.Println(sum)
}
