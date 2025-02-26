1) Геморрой работы с рунами.

Не правильно:
res = append(res, rune(countChars))
Правильно:
res = append(res, rune('0' + countChars))

2) С конца обрабатывал двузначные числа, а нужно с начала.

Ожидаемый результат

['a', 'b', '1', '2']
Результат исполнения

['a', 'b', '2', '1']

3) Для перевода int в массив рун можно преобразовать int в строку, используя strconv.Itoa(), затем перевести строку в []rune, а затем пройтись по каждой рун через for range.

```
    strCountChars := strconv.Itoa(countChars)
    arrRuneStrCountChars := []rune(strCountChars)
    for _, ch := range arrRuneStrCountChars {
        res = append(res, ch)
    }
```