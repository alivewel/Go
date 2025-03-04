1) В цикле было условие i == hIndex, а нужно было i <= hIndex

```
if i == hIndex { // i <= hIndex 
    return i
}
```

2) В цикле возвращал hIndex а нужно было i.
```
if i == hIndex {
    return hIndex // return i
}
```

3) Второй цикл скопировал из первого и забыл, что нужно с конца проходиться. То, что с конца нужно проходиться помнил и говорил об этом, когда проговаривал решение, из-за копипаста забыл.

```
for i := len(countCit)-1; i >= 0; i-- 
```

Изначально создал 
```
i := len(countCit)
```
а нужно было 
```
i := len(countCit)-1
```

4) Компилятор ругался на мелкую ошибку.
Имеется
countCit := make([]int, len(citations+1))
Должно быть
countCit := make([]int, len(citations)+1)
