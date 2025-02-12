package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	expectedSum := 0          // считаем сумму арифм. прогрессии вручную
	for i := 0; i <= n; i++ { // проходимся от 0 до n + 1
		expectedSum += i
	}
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums)) // 8
	nums = []int{3, 0, 1}
	fmt.Println(missingNumber(nums)) // 2
	nums = []int{0, 1}
	fmt.Println(missingNumber(nums)) // 2
}

// time: O(n)
// mem: O(1)

// несмотря на то, что мы считаем сумму арифм. прогрессии вручную
// пространственная сложность остаётся O(1)

// Выполнение цикла для нахождения суммы арифметической прогрессии 
// не увеличивает пространственную сложность, так как это просто 
// итерация с использованием одной переменной
