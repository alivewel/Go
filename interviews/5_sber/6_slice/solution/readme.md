### Что такое слайс? Почему он так называется?

Слайс - динамически расширяемый массив, который может расширятся с помощью функции append.

#### Структура слайса

```go
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```


- **Pointer**: `*array` — базовый массив.
  - Ссылка указывает на первый элемент массива.
  - Последний используемый элемент определяется с помощью `Length`.

- **Length**: `int`
  - Это количество элементов, которые мы "отщипываем" от базового массива.

- **Capacity**: `int`
  - Это количество элементов, которые могут поместиться в базовый массив, учитывая, что первый элемент слайса не обязательно совпадает с первым элементом массива.
  - Например, если размер массива 10, но мы начинаем использовать его с 5-го элемента, то `capacity` будет 5.

### Как работает функция append?

1. **Для слайсов до 256 элементов:**
   - Емкость увеличивается в два раза. Это позволяет эффективно управлять памятью и минимизировать количество операций по перераспределению памяти.

2. **Для слайсов более 256 элементов:**
   - Емкость увеличивается по формуле: 
     ![increase_capasity](images/increase_capasity.png)
   - Здесь константа равна 256. Этот подход позволяет более плавно увеличивать емкость, что особенно важно для больших слайсов, чтобы избежать резкого увеличения использования памяти.


### Что выведет программа?

```go
func main() {
	slice := make([]int, 0, 1000)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)     // ???
	process(slice)
	fmt.Println(slice)     // ???
	fmt.Println(slice[:6]) // ???
}

func process(slice []int) {
	for index := range slice {
		slice[index] = 0
	}
}
```

### Вывод программы

1. **Первый вывод: `fmt.Println(slice)`**  
   - Вывод: `[1 2 3]`
   - Изначально в слайс добавляются элементы 1, 2 и 3.

2. **Второй вывод: `fmt.Println(slice)` после `process(slice)`**  
   - Вывод: `[0 0 0]`
   - Функция `process(slice)` заменяет каждый элемент слайса на 0, изменяя его на месте.

3. **Третий вывод: `fmt.Println(slice[:6])`**  
   - Вывод: `[0 0 0 0 0 0]`
   - Слайс изначально создан с длиной 0 и емкостью 1000, поэтому `slice[:6]` имеет доступ к первым шести элементам подлежащего массива. Так как `process(slice)` изменяет только первые три элемента, остальные элементы остаются нулевыми, и выводится `[0 0 0 0 0 0]`.
