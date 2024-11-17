package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Введите текст (Ctrl+D для завершения):")

	// Копируем данные из stdin в stdout
	n, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при копировании:", err)
		return
	}

	fmt.Printf("\nСкопировано %d байт\n", n)
}
