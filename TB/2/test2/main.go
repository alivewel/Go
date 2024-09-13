package main

import (
	"fmt"
)

// Функция для замены -1 на ближайшие значения и проверки на возрастание разницы
func replaceAndValidation(precipitation *[]int) (bool, []int) {
	precDif := make([]int, 0)
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
				// Ищем ближайшее правое значение
				for right < len(*precipitation) && (*precipitation)[right] == -1 {
					right++
				}
				// Если нашли правое значение, используем его
				if right < len(*precipitation) {
					(*precipitation)[i] = (*precipitation)[right] - currentDiff
				} else {
					// Если нет подходящих значений, задаем значение по умолчанию
					(*precipitation)[i] = currentDiff // Пример: использовать текущую разницу как значение
				}
			} else {
				// Замена чисел, равных -1, на ближайшее левое значение
				(*precipitation)[i] = (*precipitation)[left] + currentDiff
			}

			// Увеличиваем текущую разницу для следующего шага
			currentDiff++
		}

		// Проверка возрастания чисел
		if i > 0 {
			currentDiff = (*precipitation)[i] - (*precipitation)[i-1]
			precDif = append(precDif, currentDiff)

			if len(precDif) > 1 && precDif[len(precDif)-1] <= precDif[len(precDif)-2] {
				return false, nil
			}
		} else {
			// Для первого элемента просто добавляем его значение в разницы
			precDif = append(precDif, (*precipitation)[i])
		}
	}
	return true, precDif
}

func main() {
	countDays := 5
	// precipitation := []int{-1, -1, 6, 10, 15}
	precipitation := []int{-1, 3, 6, 10, 15}

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
		fmt.Println("Обновленный массив:", precipitation)
		fmt.Println("Разницы:", precDif)
	}
}
