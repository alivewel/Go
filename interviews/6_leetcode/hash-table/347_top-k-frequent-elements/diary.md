1) Ниже в цикле я обращаюсь по индексу. Ловил панику. Неправильное создание слайса.
Было: 
freqArr := make([][]int, 0, len(nums))
Стало:
freqArr := make([][]int, len(nums))

2) 
Тестовый кейс с nums = [1] не проходил. Ловил панику. Выход за предела массива.
Было: 
freqArr := make([][]int, len(nums)) 
Стало:
freqArr := make([][]int, len(nums)+1) 

3) Забыл, что у нас двумерный массив и нужен вложенный цикл.

Было: 
for i := len(freqArr) - 1; i >= 0; i-- {
    res = append(res, num) // res := append(res, num)
    if len(res) >= k {
        return res // break - забыл, что вложенный цикл
    }
}
	
Стало:
for i := len(freqArr) - 1; i >= 0; i-- {
    for _, num := range freqArr[i] { 
        res = append(res, num) // res := append(res, num)
        if len(res) >= k {
            return res // break - забыл, что вложенный цикл
        }
    }
}

4) Тестовый кейс с nums = [-1,-1] не проходил. Некорректная проверка на наличие элемента в массиве. for range достает, только заполенные элемента массива. Проверка не нужна.

Было: 
for _, num := range freqArr[i] { // !
    // if num > 0 {
        res = append(res, num) // res := append(res, num)
        if len(res) >= k {
            return res // break - забыл, что вложенный цикл
        }
    // }
}

Стало:
for _, num := range freqArr[i] { // !
    res = append(res, num) // res := append(res, num)
    if len(res) >= k {
        return res // break - забыл, что вложенный цикл
    }
}