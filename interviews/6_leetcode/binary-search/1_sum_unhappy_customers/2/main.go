package main

import (
    "fmt"
	"sort"
)

// **Первая:**

// На Авито размещено множество товаров, каждый из которых представлен числом.
// У каждого покупателя есть потребность в товаре, также выраженная числом.
// Если точного товара нет, покупатель выбирает ближайший по значению товар,
// что вызывает неудовлетворённость, равную разнице между его потребностью и купленным товаром.
// Количество каждого товара не ограничено, и один товар могут купить несколько покупателей.
// Рассчитайте суммарную неудовлетворённость всех покупателей.

// Нужно написать функцию, которая примет на вход два массива:
// массив товаров и массив потребностей покупателей, вычислит сумму неудовлетворённостей всех покупателей и вернет результат в виде числа.

// __________________________________
// ПРИМЕР
// Ввод
// goods = [8, 3, 5]
// buyerNeeds = [5, 14, 12, 44, 55] // 6 + 4 + 36 + 47 = 93

// Вывод
// res = 1 # первый покупатель покупает товар 5 и его неудовлетворённость = 0, второй также покупает товар 5 и его неудовлетворённость = 6-5 = 1

// needs - product = unhappy
// 0 + 6 + 4 + 36 + 47 = 93

// goods = [8, 3, 5]
// buyerNeeds = [5, 14, 12, 44, 55]

// goods = [3, 5, 8] target = 6
func binSearch(array []int, target int) int {
    if target <= array[0] {
        return array[0] - target
    }
    if target >= array[len(array)-1] {
        return target - array[len(array)-1]
    }
    l, r := 0, len(array)
    for r - l > 1 {
        m := (r + l) / 2
        if target <= array[m] {
            r = m
        } else {
            l = m
        }
    }
    leftDiff := target - array[l] // 2
    rightDiff := array[r] - target // 1
    if rightDiff >= leftDiff {
        return leftDiff
    }
    return rightDiff
}

func FindGood(goods []int, needs []int) int {
    sort.Ints(goods)
    unhappy := 0
    for _, need := range needs {
        unhappy += binSearch(goods, need)
    }
    return unhappy
}

func testFindGood() {
	type testCase struct {
		goods     []int
		needs     []int
		expected  int
	}

	tests := []testCase{
		{
			goods:    []int{8, 3, 5},
			needs:    []int{5, 14, 12, 44, 55},
			expected: 93, // 0 + 6 + 4 + 36 + 47
		},
		{
			goods:    []int{10},
			needs:    []int{10},
			expected: 0,
		},
		{
			goods:    []int{10},
			needs:    []int{8},
			expected: 2,
		},
		{
			goods:    []int{5, 10, 20},
			needs:    []int{1, 6, 15, 25},
			expected: 4 + 1 + 5 + 5, // 15
		},
		{
			goods:    []int{1, 2, 3},
			needs:    []int{10, 11, 12},
			expected: 7 + 8 + 9, // 24
            // 
		},
		{
			goods:    []int{3, 7, 15},
			needs:    []int{},
			expected: 0,
		},
	}

	for i, tc := range tests {
		result := FindGood(tc.goods, tc.needs)
		if result != tc.expected {
			fmt.Printf("❌ Test %d FAILED: expected %d, got %d\n", i+1, tc.expected, result)
		} else {
			fmt.Printf("✅ Test %d PASSED\n", i+1)
		}
	}
}

func main() {
	testFindGood()
}

