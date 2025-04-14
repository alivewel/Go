func filterSortedList(a, b []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(a) {
		// Пропускаем все элементы в b, которые меньше текущего a[i]
		for j < len(b) && b[j] < a[i] {
			j++
		}
		// Если равны — элемент есть в b, пропускаем
		if j < len(b) && a[i] == b[j] {
			i++
		} else {
			result = append(result, a[i])
			i++
		}
	}
	return result
}
