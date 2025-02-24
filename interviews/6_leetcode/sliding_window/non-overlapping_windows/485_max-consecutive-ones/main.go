package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	l, r := 0, 0
	countOnes := 0
	for r < len(nums) {
		for r+1 < len(nums) && nums[r] == nums[r+1] { 
			// Первое условие (r+1 < len(nums)) гарантирует, что r+1 не выйдет за границы массива, 
			// предотвращая обращение к несуществующему индексу nums[r+1].
			
			// Второе условие (nums[r] == nums[r+1]) проверяет, что текущий элемент равен следующему.
			// Это позволяет увеличивать r, пока идет последовательность одинаковых чисел.
			
			// Когда nums[r] != nums[r+1], цикл завершается, и r останется на последнем элементе группы.
			r++
		}
		if nums[l] == 1 {
			curCount := r - l + 1
			if curCount > countOnes {
				countOnes = curCount
			}
		}
		l = r + 1
		r = l
	}
    return countOnes 
}

func main() {
	nums := []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
