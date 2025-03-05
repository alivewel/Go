1) Забыд фигурные скобки при инициализации key := [26]int{}
2) panic выходило из-за этой ошибки/
Было:  key[ch-'0']++
Стало: key[ch-'a']++
3) 
Выходит такой ответ из-за этой ошибки: [[] [] [] [eat tea ate] [tan nat] [bat]]

Было:  res := make([][]string, len(anagrams))
Стало: res := make([][]string, 0, len(anagrams))