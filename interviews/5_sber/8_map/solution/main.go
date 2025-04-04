package main

// что выведет программа?

func main() {
	var m map[string]struct{}
	for k := range m {
		println(k)
	}
	println("end!")
}

// программа выведет end!
// мапа пустая, мы ничего туда не положили

// что такое мапа в го?
// мапа - это структура данных, которая содержит в себе пару ключ и значение.
// Под капотом в ней хеш-таблица, которая позволяет нам обращаться по ключу за константное время.
// Изначально создается мапа с 8 бакетами.
// Бакет - структура данных, в которой хранится ключи и значения в мапе. Данные внутри бакета хранятся в виде массива, в одном бакете может быть до 8 элементов. Если все аллоцированые бакеты в среднем заполнены на 6,5, то начинается эвакуация данных.
// Важное свойство для ключа - он должен быть comparable (сравнимый).
// Нельзя взять указатель на элемент мапы как раз по причине эвакуации данных.

// У нас есть хеш-функция, которая возвращает нам какой-то хеш.
// Хеш представляет из себя значение числа int64. Бакетов у нас меньше? 
// Как вычислить в какой бакет положить пару ключ-значение?

// Мы берем остаток от деления от количества от бакетов.
