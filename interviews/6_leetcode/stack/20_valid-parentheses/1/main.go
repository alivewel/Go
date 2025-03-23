package main

import (
	"fmt"
)

func isValid(brackets string) bool {
    stackOpen := Stack{}
    for _, ch := range brackets {
        if ch == '{' || ch == '(' || ch == '[' {
            stackOpen.Push(ch)
        } else {
            prevBrackets, _ := stackOpen.Pop()
            if prevBrackets == '{' && ch != '}' {
                return false 
            } else if prevBrackets == '(' && ch != ')' {
                return false 
            } else if prevBrackets == '[' && ch != ']' {
                return false 
            }
        }
    }
 
    return stackOpen.Size() == 0
}

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
	// s := "()[]{}"
	s := "]" // Expected false
	fmt.Println(isValid(s))
}

