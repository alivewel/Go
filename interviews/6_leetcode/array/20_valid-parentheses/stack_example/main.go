package main

import "fmt"

// Структура для стека
type Stack struct {
    elements []rune // Используем rune для хранения символов
}

// Добавление символа в стек
func (s *Stack) Push(value rune) {
    s.elements = append(s.elements, value)
}

// Удаление символа из стека
func (s *Stack) Pop() (rune, bool) {
    if len(s.elements) == 0 {
        return 0, false // Если стек пуст
    }
    lastIndex := len(s.elements) - 1
    value := s.elements[lastIndex]
    s.elements = s.elements[:lastIndex]
    return value, true
}

// Проверка, пуст ли стек
func (s *Stack) IsEmpty() bool {
    return len(s.elements) == 0
}

// Получение количества элементов в стеке
func (s *Stack) Size() int {
    return len(s.elements)
}

func main() {
    stack := Stack{}

    // Добавляем символы в стек
    stack.Push('A')
    stack.Push('B')
    stack.Push('C')

    // Выводим количество элементов
    fmt.Println("Количество элементов в стеке:", stack.Size())

    // Извлекаем символы
    for !stack.IsEmpty() {
        value, _ := stack.Pop()
        fmt.Printf("Извлечен символ: %c\n", value)
    }

    // Проверяем количество элементов после извлечения
    fmt.Println("Количество элементов в стеке:", stack.Size())
}