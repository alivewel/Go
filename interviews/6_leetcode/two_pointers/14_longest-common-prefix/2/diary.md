1) Копипаст подвел. Повторное использование оператор моржа `:=`
Было:
flagExit := true
Стало:
flagExit = true

2) Перепутал последовательность prevCh и rune. Ошибка синтаксиса.
Было:
var rune prevCh
Стало:
var prevCh rune

3) Здесь нужно byte использовать.
curCh := strs[i][index] возвращает byte.
Было:
var prevCh rune
Стало:
var prevCh byte 
