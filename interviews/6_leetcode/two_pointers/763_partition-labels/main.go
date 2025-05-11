// проходимся по строке, заводим мапу ключ - символ, значение - индекс, когда последний раз встретили элемент
// заводим 2 переменные start - индекс первого элемента, end - индекс, когда последний раз встретили текущий элемент 
// проходимся по строке снова
// если для другого символа в подстроке встретили end, который больше чем текущий end, обновляем его
// как только дошли до индекса, равному end, обновляем старт, считаем разницу между start и end, и записываем ее в результат
func partitionLabels(s string) []int {
    if len(s) == 0 {
        return []int{}
    }
    lastIndChar := make(map[rune]int, 0)
    for i, ch := range s {
        lastIndChar[ch] = i
    }
    start, end := 0, lastIndChar[rune(s[0])]
    res := make([]int, 0)
    // a b c
    // 0 1 2
    for i, ch := range s {
        if lastIndChar[ch] > end {
            end = lastIndChar[ch]
        }
        if i == end {
            res = append(res, end - start + 1)
            start = i + 1
        } 
    }
    return res
}
