package main

// что выведет программа?

func main() {
	var m map[string]struct{}
	for k := range m {
		println(k)
	}
	println("end!")
}

// Что такое мапа в го?

// У нас есть хеш-функция, которая возвращает нам какой-то хеш.
// Хеш представляет из себя значение числа int64. Бакетов у нас меньше? 
// Как вычислить в какой бакет положить пару ключ-значение?
