package main

import "fmt"

type NumArray struct {
	arr []int
	px  []int
}

func Constructor(nums []int) NumArray {
	px := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			px[i] = num
			continue
		}
		px[i] += px[i-1] + num
	}
	return NumArray{
		arr: nums,
		px:  px,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return 0
}

func main() {
	arr := []int{1, 2, 3, 4}
	px := Constructor(arr)
	fmt.Println(px.px)
}
