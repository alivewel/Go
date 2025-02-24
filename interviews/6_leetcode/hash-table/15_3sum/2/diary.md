1) Какую-то фигню не рабочую написал.
Вместо mapSum[complement] = struct{}{} нужно использовать mapSum[nums[j]] = struct{}{}.

Ошибка в индексе j
Начинаю внутренний цикл с j := i, но это приводит к тому, что я сравниваю элемент с самим собой (даже если пропускаю i == j). Это неэффективно и может привести к ошибкам.
Исправление: Нужно начинать внутренний цикл с j := i + 1, чтобы избежать проверки текущего элемента дважды.

Добавить обработку дублей.

```
// решение с помощью хэш-таблицы
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		// mapSum := make(map[[3]int]struct{}) 
		mapSum := make(map[int]struct{}) 
		for j := i; j < len(nums); j++ {
			if i == j {
				continue
			}
			complement := target - nums[j] // complement - дополнение
			if _, found := mapSum[complement]; found {
				// mapSum[[3]int{nums[i], complement, nums[j]}] = struct{}{}
				res = append(res, []int{nums[i], complement, nums[j]})
			}
			mapSum[complement] = struct{}{}
		}	
	}
	return res
}
```
