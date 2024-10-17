package main

func main() {
	str := "Привет!"
	charCount := 0

	// посчитать кол-во символов к строке
	for range str {
		charCount++
	}

	if charCount == 7 {
		println("Success!")
	}
}

// самый простой способ пройтись в строке и инкрементить счетчик с каждой итерации
