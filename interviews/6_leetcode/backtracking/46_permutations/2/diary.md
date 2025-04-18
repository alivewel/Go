1) Из-за этого пустой массив возвращалась
Имеется:
copyComb := make([]int, 0, len(curComb))
Должно быть:
copyComb := make([]int, len(curComb))

2) Забыл про эту строку
curComb = curComb[:len(curComb)-1]

3) Перепутал аргументы в сигнатуре функции. Сначала dist, потом source.
Имеется:
copy(curComb, copyComb)
Должно быть:
copy(copyComb, curComb)
