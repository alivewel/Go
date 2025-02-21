1) Была проблема с конвертацией int в byte

``` go
func compress(chars []byte) int {
	res := make([]byte, 0, len(chars))
	l, r := 0, 0
	for r < len(chars) {
		for r + 1 < len(chars) && chars[r] == chars[r+1] {
			r++
		}
		countChars := r - l + 1
		res = append(res, byte(chars[l]))
		if countChars > 1 && countChars < 10 {
			fmt.Println(countChars)
			// res = append(res, byte(countChars))
			res = append(res, strconv.Itoa(countChars)...)
		} else if countChars > 9 {
			for countChars != 0 {
				// res = append(res, byte(countChars % 10))
				res = append(res, strconv.Itoa(countChars % 10)...)
				countChars /= 10
			} 
		}
		// если символ встречается 1 раз ничего не делаем
		l = r + 1
		r = r + 1
	}
	chars = res // временное решение
	fmt.Println(string(res))
    return len(res)
}
```
2) 
Было:
res = append(res, strconv.Itoa(countChars))
Нужно:
res = append(res, strconv.Itoa(countChars)...)
Выходила ошибка:
cannot use strconv.Itoa(countChars) (value of type string) as byte value in argument to append

Ответ GPT:

Ошибка возникает из-за того, что `strconv.Itoa(countChars)` возвращает `string`, а `append` ожидает `byte`. Нужно добавить `...` (спред-оператор), чтобы развернуть строку в последовательность байтов.

Вот исправленный код:

```go
import "strconv"

res = append(res, []byte(strconv.Itoa(countChars))...)
```

### Исправленный вариант для `countChars > 9`:
В этом блоке кода также нужно преобразовывать число в строку, а не пытаться добавлять `byte` напрямую:

```go
for _, ch := range strconv.Itoa(countChars) {
    res = append(res, byte(ch)) // ch уже является руной (rune), можно привести к byte
}
```

### Итог:
- Используйте `strconv.Itoa(countChars)`, чтобы преобразовать число в строку.
- Разворачивайте строку в `[]byte` при `append` через `...`.
- В цикле `for countChars != 0`, вместо работы с `byte(countChars % 10)`, используйте `strconv.Itoa(countChars)`.

Это исправит вашу проблему с добавлением чисел в `res`.

----------

3) Внутри функции пытался chars = res // временное решение
Чтобы chars изменялся внутри функции, нужно делать: copy(chars, res), если res не длиннее chars.