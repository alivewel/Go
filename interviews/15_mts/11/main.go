package main

import (
	"fmt"
	"unsafe"
)

func main() {
	nums := make([]int, 1, 2)
	fmt.Println(nums) // <- what's the output? []

	appendSlice(nums, 1024)
	fmt.Println(nums) // <- what's the output? []

	mutateSlice(nums, 1, 512)
	fmt.Println(nums)                // <- what's the output? []
	fmt.Println(unsafe.Sizeof(nums)) // 4 + 4 + 4 = 12
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
