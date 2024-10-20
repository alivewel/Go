package main

import "fmt"

func twoSum(nums []int, target int) []int {
	checkSum := make(map[int]int) // ключ - первое слогаемое, значение - индекс в массиве
	for index1, value := range nums {
		if index2, found := checkSum[target-value]; found { // здесь value - второе слогаемое
			return []int{index1, index2} // первое слогаемое == сумма - второе слогаемое
		}
		checkSum[value] = index1 // здесь value - первое слогаемое
	}
	return nil
}

func main() {
	// nums := []int{2, 7, 11, 15}
	nums := []int{3, 3}
	res := twoSum(nums, 6)
	fmt.Println(res)
}

// мы создаем мапу в которой хранится первое слогаемое и его индекс.
// по проходимся по массиву чисел и постепенно добаляем эти числа
// проверяем есть ли в нашей мапе значение с ключем первого слагаемого.
// [target-value] - это комплементарное число, разница между суммой и вторым слогаемым равна первому слогаемому.
