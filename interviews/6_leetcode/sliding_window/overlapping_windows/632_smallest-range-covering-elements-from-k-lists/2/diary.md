1) В этом фрагменте:
Имеется:
``` go
for _, curNums := range nums {
	for i, num := range curNums {
		elements = append(elements, Element{Idx: i, Val: num})
	}
}
```

Должно быть:
``` go
for i, curNums := range nums {
	for _, num := range curNums {
		elements = append(elements, Element{Idx: i, Val: num})
	}
}
```
Ты используешь i как индекс группы (Idx), но на самом деле i — это индекс внутри подмассива, а не номер группы!

В результате у тебя структура elements получает неправильные значения Idx, что нарушает логику подсчёта количества покрытых групп в count.
