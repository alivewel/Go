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
			// Поиск ближайшего числа, не равного -1
			left := i - 1
			right := i + 1

			// Сначала ищем ближайшее левое значение
			for left >= 0 && (*precipitation)[left] == -1 {
				left--
			}

			// Если не нашли подходящее левое значение, ищем правое
			if left < 0 {
				for right < len(*precipitation) && (*precipitation)[right] == -1 {
					right++
				}
				// Если нашли правое значение, используем его
				if right < len(*precipitation) {
					// (*precipitation)[i] = (*precipitation)[right] - currentDiff - 1
					(*precipitation)[i] = (*precipitation)[right] - currentDiff
				} else {
					// Если нет подходящих значений, можно задать значение по умолчанию
					fmt.Println("precDi!", (*precipitation)[i], currentDiff)
					(*precipitation)[i] = currentDiff // Пример: использовать текущую разницу как значение
				}
			} else {
				// Замена чисел, равных -1, на ближайшее левое значение
				(*precipitation)[i] = (*precipitation)[left] + currentDiff
			}

			// Увеличиваем текущую разницу для следующего шага
			currentDiff++
		}

		// проверка возрастания чисел
		if i > 0 {
			// fmt.Println(i, (*precipitation)[i], (*precipitation)[i]-(*precipitation)[i-1])
			// Записываем разницу
			currentDiff = (*precipitation)[i] - (*precipitation)[i-1]
			precDif = append(precDif, currentDiff)
			// fmt.Println("precDif", precDif[i], precDif[i-1])
			if len(precDif) > 1 && precDif[len(precDif)-1] <= precDif[len(precDif)-2] {
				return false, nil
			}
			// if precDif[i] <= precDif[i-1] {
			// 	return false, nil
			// }
		} else {
			// Для первого элемента просто добавляем его значение в разницы
			precDif = append(precDif, (*precipitation)[i])
		}
	}
	return true, precDif
}

func main() {
	countDays := 5
	// precipitation := []int{1, 3, -1, 10, -1}
	// precipitation := []int{1, 3, 6, 10, 15}
	precipitation := []int{-1, 3, 6, 10, 15}
	// precipitation := []int{-1, -1, 6, 10, 15}

	if countDays != len(precipitation) {
		fmt.Println("NO")
		return
	}

	isIncreasing, precDif := replaceAndValidation(&precipitation)
	if !isIncreasing {
		fmt.Println("NO")
		return
	} else {
		fmt.Println("YES")
		fmt.Println(precDif)
	}
	fmt.Println("precipitation", precipitation)
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
// Как еще можно использовать информацию о количестве дней? Она ведь зачем-то нужна.
// 7. Обработать случай, когда первые несколько чисел равны -1.
// Пример: {-1, 3, 6, 10, 15}, {-1, -1, 6, 10, 15}.
// 8. Обработать случай, когда несколько чисел равны -1 и находятся в середине.
// Пример: {1, 3, -1, -1, 15}.
// 8. Обработать случай, когда несколько чисел равны -1 и находятся в конце.
// Пример: {1, 3, 6, -1, -1}.
// 9. 

// 1. Переменная количества дней нигде не используется.
// Добавил проверку равенства количества дней и длины массива чисел количества выпавших осадков.
// Как еще можно использовать информацию о количестве дней? Она ведь зачем-то нужна.
// 2. Обработать случай, когда первые несколько чисел равны -1.
// Пример: {-1, 3, 6, 10, 15}, {-1, -1, 6, 10, 15}.
// 3. Обработать случай, когда несколько чисел равны -1 и находятся в середине.
// Пример: {1, 3, -1, -1, 15}.
// 4. Обработать случай, когда несколько чисел равны -1 и находятся в конце.
// Пример: {1, 3, 6, -1, -1}.