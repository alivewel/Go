package main

import "fmt"

func longestStockGrowth(stock []int) int {
	l, r := 0, 0
	result := 0
	countZero, idxZero := 0, -1
	if stock[0] == 0 {
		countZero = 1
		idxZero = 0
	}
	// for r < len(stock) { // 1 косяк
	for l < len(stock) {
		// for r+1 < len(stock) && (stock[r] == stock[r+1] || countZero < 1) { // 2 косяк 
		for r+1 < len(stock) && (stock[r+1] == 1 || countZero < 1) { 
			if stock[r+1] == 0 {
				countZero++
				idxZero = r + 1
			}
			r++
		}
		curCountZero := r - l + 1
		result = max(result, curCountZero)
		countZero = 0
		// l = r + 1 // 3 косяк 
		// r = r + 1 // 4 косяк 
		if idxZero != -1 {
			// l = max(r + 1, idxZero) // 5 косяк 
			l = max(l + 1, idxZero+1)
		} else {
			// l = r + 1 // чтобы не зациклиться // 6 косяк 
			l = l + 1
		}
	}
    return result 
}

// func longestStockGrowth(stock []int) int {
//     l := 0 // Левый указатель окна
//     r := 0 // Правый указатель окна
//     result := 0 // Результирующая максимальная длина
//     zerosCount := 0 // Количество нулей в текущем окне
//     zeroIdx := -1 // Индекс последнего нуля в окне

//     if len(stock) > 0 && stock[0] == 0 {
//         zerosCount = 1
//         zeroIdx = 0
//     }

//     for l < len(stock) {
//         // Расширяем окно, пока следующий элемент 1 или количество нулей меньше 1
//         for r+1 < len(stock) && (stock[r+1] == 1 || zerosCount < 1) {
//             if stock[r+1] == 0 {
//                 zerosCount += 1
//                 zeroIdx = r + 1
//             }
//             r += 1
//         }

//         // обновляем ответ
//         windowSize := r - l + 1
//         if windowSize > result {
//             result = windowSize
//         }

//         // сдвигаем на zeroIdx когда у нас есть 0 в окне
//         // сдвигаем на l + 1 если нулей в окне нет (нужно чтобы не зациклиться)
//         zerosCount = 0
//         if zeroIdx != -1 {
//             l = max(zeroIdx+1, l+1)
//         } else {
//             l = l + 1
//         }
//     }

//     return result
// }


func main() {
	// stock := []int{1,1,0,1,1,1}
	stock := []int{0, 0, 0, 0}
	// stock := []int{0, 1, 1, 1, 0, 1, 1}
	// stock := []int{0, 1, 0, 1, 0, 1, 0}
	fmt.Println(longestStockGrowth(stock))
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