// Простой калькулятор
// Дан массив s с положительными числами и знаками * и +. Нужно вернуть результат вычислений и важно, чтобы решение было за O(1) по дополнительной памяти.

// Пример 1:
// Ввод: s = ["2","*","3","*","1","+","2","+","2","+","2","*","0","*","13","+","1"]
// Вывод: 11

// Пример 2:
// Ввод: s = ["1","+","2","+","3"]
// Вывод: 6*-

// Ограничения:
// len(s) >= 1

// попробую сначала решить через очередь
package main

import (
    "fmt"
    "strconv"
)
                //  3
// s = ["1","+","2","+","3","*","4"]
// queueNums 3 12
// queueOper + 
// curRes
// s = ["3","*","4","+","5"]

func calculate(s []string) int {
    stackNums := Stack{}
    stackOper := Stack{}
    for _, str := range s {
        if isNum(str) {
            stackNums.Push(str)
        }
        if isOper(str) {
            if str == "+" || str == "*" {
              if stackOper.Size > 0 && stackNums.Size > 1 {
                  num1, num2 := stackNums.Pop(), stackNums.Pop()
                  oper := stackOper.Pop()
                  if oper == "+" {
                      res := strconv.Atoi(num1) + strconv.Atoi(num2)
                      stackNums.Push(strconv.Itoa(res))
                  } else if oper == "*" {
                      res := strconv.Atoi(num1) * strconv.Atoi(num2)
                      stackNums.Push(strconv.Itoa(res))
                  }
              }
              stackOper.Push(str)  
            }
        }
    }
    if stackOper.Size > 0 && stackNums.Size > 1 {
      num1, num2 := stackNums.Pop(), stackNums.Pop()
      oper := stackOper.Pop()
      if oper == "+" {
          res := strconv.Atoi(num1) + strconv.Atoi(num2)
          stackNums.Push(strconv.Itoa(res))
      } else if oper == "*" {
          res := strconv.Atoi(num1) * strconv.Atoi(num2)
          stackNums.Push(strconv.Itoa(res))
      }
    }
    return stackNums.Pop()
}

// Структура для стека
type Stack struct {
    elements []string // Используем string для хранения символов
}

// Добавление символа в стек
func (s *Stack) Push(value string) {
    s.elements = append(s.elements, value)
}

// Удаление символа из стека
func (s *Stack) Pop() (string, bool) {
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
    s := []string{"2","*","3","*","1","+","2","+","2","+","2","*","0","*","13","+","1"}
	fmt.Println(calculate(s))
}


// strconv.Itoa(num) - конвертация числа в строку
// strconv.Atoi(str) - конвертация строки в число