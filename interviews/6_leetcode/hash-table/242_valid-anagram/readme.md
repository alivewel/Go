Нашел 2 способа решения задач - с использованием мапы и массива.

1) Первый способ - мапа 
В первом случае в мапе в качестве ключа используется символ, в качестве значения количество раз, которое встречается данная буква в слове.

Проходимся по первой строке и с каждым разом увеличиваем счетчик на 1. (+1)
Проходимся по второй строке и с каждым разом уменьшаем счетчик на 1. (-1)
В конце у нас сумма для каждого сиимвола должна получиться равной нулю. Значит нам встретилась анаграмма.

2) Второй способ - массив
Сначала создаем слайс []int размером 26 (количество букв в английском алфавите).

Далее обращаемся по индексу от 0 до 25.
ch := 'a'
arr[ch-'a']
В данном случае индекс будет равен нулю. Если ch := 'b' индекс будет равен одному. И т.д.

time: O(n)
mem: O(1) т к размер массивов всегда 26
