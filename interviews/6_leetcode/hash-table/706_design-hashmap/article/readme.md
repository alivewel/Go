### Простая реализация мапы

Сходил я тут на собеседование в Авито.
1 этап. Алгоритмическая секция. Решил обе задачи. Довольный собой жду фидбека от рекрутера. На следующий день получаю вот такое сообщение:

Текст специально изменен. Личную переписку не хорошо палить.

Моя первая реакция была такая:

![wtf](images/what.png)

Отказ, после двух решенных задачек??? У вас там все нормально???

Чуть позже с холодной головой я провел рефлексию. В чем же оказалась проблема на самом деле?

После решения сдачи первой задачи интервьюрер задал мне ряд вопросов:

И:
Почему я использовал для эту задачу мапу? Почему нельзя использовать массив или другую структура данных?
Мой ответ:
Мапа нам дает следующие свойства:
- Для группировки по номерам бегунов
- Уникальные номера бегунов
- Константное время получения значения по ключу

И: За счет чего достигается 3-е свойство мапы?
Я: За счет использования бакетов и хеш-функции.

И: У нас есть идеальная хэш-функция. Она классно работает и никогда не повторяется. У нас нет времени делать бакеты, нужно сделать максимально простую реализацию мапы. Нужно быстро и дешево написать нашу хэш-мапу, чтобы она просто работала без использования дополнительных библиотек.
Я: ...

Признаюсь здесь я немного офигел от такого запроса. Чуть позже мы накидали следующую структуру:

``` go
type OurSuperMap struct {
    key []any 
    value []any
}
func (m OurSuperMap) OurSuperHashFunc(value any) int64 {...}

// Реализовать функцию вставки в мапу (можно просто словами).
func (m OurSuperMap) Add(key any, value any) {

}
```

На собесе реализовать я не смог. Чуть позже проводя разбор собеса я нашел похожую задачу на литкоде:
706. Design HashMap - https://leetcode.com/problems/design-hashmap/
Похоже что-то подобное ожидал от меня услышать интервьюрер.

По факту задача несложная. Посидев пару часов я с ней разобрался. Вот мое решение этой задачи:

``` Go
type MyHashMap struct {
    buckets [][]Pair
}

type Pair struct {
    key, value int
}

func Constructor() MyHashMap {
    size := 8
	return MyHashMap{
		buckets: make([][]Pair, size),
	}
}

func (this *MyHashMap) Put(key int, value int) {
    // определяем номер бакета
	index := key % len(this.buckets)
	// определяем бакет
	bucket := this.buckets[index]
	// пытаемся найти ключ в бакете
	for i, pair := range bucket {
		if pair.key == key {
			this.buckets[index][i].value = value
			return
		}
	}
	this.buckets[index] = append(this.buckets[index], Pair{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
    index := key % len(this.buckets)
	// определяем бакет
	bucket := this.buckets[index]
	// пытаемся найти ключ в бакете
	for _, pair := range bucket {
		if pair.key == key {
			return pair.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    newBucket := []Pair{}
    index := key % len(this.buckets) // определяем бакет
	bucket := this.buckets[index] // пытаемся найти ключ в бакете
	for _, p := range bucket {
		if p.key != key {
			newBucket = append(newBucket, p)
		}
	}
	this.buckets[index] = newBucket
}
```
Недостатки этой реализации:
1) В этой задаче от нас не требуется увеличение количества бакетов по ходу заполнения мапы. Мы один раз создали мапу с фиксированным количеством бакетов в конструкторе и далее с ней работаем. 
2) Здесь мы никак не боремся с коллизиями. 
3) Мы не учитываем дублирующиеся ключи, только в пределах одного бакета.
4) Нет ограничения по размеру бакета. Это может нам снизить алгоритмическую сложность получения элемента в мапе с О(1) до О(К), где К - количество элементов в бакете.
5) Нет равномерного распределения ключей по мапе. Мы просто берем ключ и вычисляем от него остаток от деления от количества бакетов. В реальной мапе вычисляется хэш (какое-то большое число int) и от него берется остаток от деления и количества бакетов.

Данное решение с бакетами, а интервьюрер хотел увидеть решение без бакетов. Вот как бы я сейчас ответил на этот вопрос:

Реализация с бакетами:
``` Go
type Pair struct {
    key, value int
}

type OurSuperMap struct {
    // buckets [][]Pair c бакетами
    elements []Pair // без бакетами
}

func (m OurSuperMap) OurSuperHashFunc(value any) int64 {...}

// Реализовать функцию вставки в мапу (можно просто словами).
func (m OurSuperMap) Add(key any, value any) {
    index := m.OurSuperHashFunc(key)
    bucket := m.buckets[index]

    // сначала пытаемся найти ключ в бакете и изменить значение по ключу
    for i, pair := range bucket {
        if pair.key == key {
            m.buckets[index][i].value = value
            return
        }
    }

    // если ключ не найден — добавляем новую пару
    m.buckets[index] = append(m.buckets[index], Pair{key: key.(int), value: value.(int)})
}
```

Реализация без бакетов:
``` Go
type Pair struct {
    key, value int
}

type OurSuperMap struct {
    elements []Pair // без бакетов
}

func (m OurSuperMap) OurSuperHashFunc(value any) int64 {...}

// Реализовать функцию вставки в мапу (можно просто словами).
func (m OurSuperMap) Add(key any, value any) {
    // сначала пытаемся найти ключ в бакете и изменить значение по ключу
    for i, pair := range m.elements {
        if pair.key == key {
            m.elements[i].value = value
            return
        }
    }

    // если ключ не найден — добавляем новую пару
    m.elements = append(m.elements, Pair{key: key.(int), value: value.(int)})
}
```
