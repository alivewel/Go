1) 

```
func longestGeneSequence(gene string) int {
    l, r := 0, -1 // -1
	chars := make(map[byte]struct{})
	maxCountCh := 0
	for l < len(gene) {
		for r + 1 < len(gene) && !containsMap(chars, gene[r+1]) { // containsMap(chars, gene[r]) 
			chars[gene[r+1]] = struct{}{} // struct{}
			r++
		}
		delete(chars, gene[l])
		sizeWindow := r - l + 1
		if sizeWindow > maxCountCh {
			maxCountCh = sizeWindow
		}
		l++
	}
	return maxCountCh
}
```

2) Не правильно расчитывалось плавающие окно из-за этого. Не правильное булево значение из функции containsMap возвращал.
Имеется:
for r + 1 < len(gene) && containsMap(chars, gene[r+1])
Должно быть:
for r + 1 < len(gene) && !containsMap(chars, gene[r+1])

3) Сначала не понял зачем r создавать  с значением -1.

Имеется:
l, r := 0, 0
Должно быть:
l, r := 0, -1

4) Невнимательность.

Имеется:
chars[gene[r+1]] = struct{}
Должно быть:
chars[gene[r+1]] = struct{}{}

5) Сначала не понял зачем r создавать  с значением -1. Понял для чего.

Имеется:
chars[gene[r]] = struct{}{}
Должно быть:
chars[gene[r+1]] = struct{}{}

6) Сначала не понял зачем r создавать  с значением -1. Понял для чего.

Имеется:
for r  < len(gene)
Должно быть:
for r + 1 < len(gene)
