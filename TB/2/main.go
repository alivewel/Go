package main

import (
	"fmt"
)

func replaceAndValidation(precipitation *[]int) (bool, []int) {
	precDif := make([]int, 0)
	// precDif[0] = (*precipitation)[0]
	currentDiff := 1
	for i := 0; i < len(*precipitation); i++ {
		
		if (*precipitation)[i] == -1 {
			// поиск ближайшего числа, не равного -1
			left := i - 1
			for left >= 0 && (*precipitation)[left] == -1 {
				left--
			}
			// замена чисел, равных -1
			if left >= 0 {
				(*precipitation)[i] = (*precipitation)[left] + currentDiff
			}
		}
		// проверка возрастания чисел
		if i > 0 {
			if (*precipitation)[i] <= (*precipitation)[i-1] {
				return false, nil
			}
			fmt.Println(i, (*precipitation)[i], (*precipitation)[i]-(*precipitation)[i-1])
			// Записываем разницу
			currentDiff = (*precipitation)[i]-(*precipitation)[i-1]
			precDif = append(precDif, currentDiff)
			// Увеличиваем текущую разницу для следующего шага
			currentDiff++
		} else {
			// Для первого элемента просто добавляем его значение в разницы
			precDif = append(precDif, (*precipitation)[i])
		}
	}
	return true, precDif
}

func main() {
	countDays := 5
	precipitation := []int{1, 3, -1, 10, -1}

	if countDays != len(precipitation) {
		fmt.Println("NO")
		return
	}

	isIncreasing, precDif := replaceAndValidation(&precipitation)

	fmt.Println("Обновленный массив:", precipitation, isIncreasing, precDif)
}

// 1. Принять число n - количество дней и n - целых чисел,
// массив чисел, разделенных пробелом.
// 2. Вместо чисел -1 подставить предыдущее значение массива +1.
// 3. Провалидировать массив чисел на предмет возрастания.
// Числа должны возрастать как минимум на 1.
// 4. Вывести NO, если валидация не прошла.
// 5. Вывести YES, если валидация прошла и вывести график возрастания количества осадков.
// 6. Переменная количества дней нигде не используется.
// Добавил проверку равенства количества дней и длины массива чисел количества выпавших осадков.
// 7. Провалидировать, когда первые несколько чисел равны -1.
