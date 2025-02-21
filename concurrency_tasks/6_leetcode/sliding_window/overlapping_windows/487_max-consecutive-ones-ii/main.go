package main

import "fmt"

func longestStockGrowth(nums []int) int {
	l, r := 0, 0
	result := 0
	countZeros := 0
	for r < len(nums) {
		for r+1 < len(nums) && (nums[r] == nums[r+1] || countZeros < 1) { 
			if nums[r] == 0 {
				countZeros++
			}
			r++
		}
		fmt.Println(l, r)
		if nums[l] != 0 {
			curCountZero := r - l + 1
			result = max(result, curCountZero)
			countZeros = 0
		}
		l = r + 1
		r = r + 1
	}
    return result 
}

func main() {
	// nums := []int{1,1,0,1,1,1}
	nums := []int{0, 0, 0, 0}
	// nums := []int{0, 1, 1, 1, 0, 1, 1}
	// nums := []int{0, 1, 0, 1, 0, 1, 0}
	fmt.Println(longestStockGrowth(nums))
}

// func runTests() {
// 	tests := []struct {
// 		input    []int
// 		expected int
// 	}{
// 		{[]int{1, 1, 0, 1, 1, 1}, 6}, // Пример из задачи
// 		{[]int{0, 0, 0, 0}, 1},       // Все нули
// 		{[]int{1, 1, 1, 1}, 4},       // Все единицы
// 		{[]int{1, 0, 1, 0, 1}, 3},    // Чередующиеся единицы и нули
// 		{[]int{1, 1, 0, 0, 1, 1}, 4}, // Две группы единиц с нулями между ними
// 		{[]int{0, 1, 1, 1, 0, 1, 1}, 6}, // Нули в начале и конце
// 		{[]int{1, 0, 0, 1, 1, 1, 0, 1}, 5}, // Нули в середине
// 		{[]int{1, 1, 1, 0, 1, 1, 1, 1}, 8}, // Один ноль в середине
// 		{[]int{0, 1, 0, 1, 0, 1, 0}, 3}, // Максимум с одним нулем
// 		{[]int{1, 0, 1, 0, 1, 0, 1, 0}, 3}, // Чередующиеся с одним нулем
// 		{[]int{1, 1, 1, 0, 0, 1, 1, 1, 1}, 5}, // Две группы единиц с нулями
// 	}

// 	for _, test := range tests {
// 		result := longestStockGrowth(test.input)
// 		if result == test.expected {
// 			fmt.Printf("Test passed for input %v: got %d, expected %d\n", test.input, result, test.expected)
// 		} else {
// 			fmt.Printf("Test failed for input %v: got %d, expected %d\n", test.input, result, test.expected)
// 		}
// 	}
// }

// func main() {
// 	runTests()
// }