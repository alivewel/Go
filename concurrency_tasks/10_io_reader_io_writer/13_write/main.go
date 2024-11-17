package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Println("Введите текст:")

	// Читаем строку из stdin
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении:", err)
		return
	}

	// Записываем введенный текст в stdout
	_, err = writer.WriteString("Вы ввели: " + input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при записи:", err)
		return
	}

	// Сбрасываем буфер, чтобы данные были записаны в stdout
	err = writer.Flush()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при сбросе буфера:", err)
	}
}
