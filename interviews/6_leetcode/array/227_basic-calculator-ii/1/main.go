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
                 3
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
            if str == "+" {
              if stackOper.Size > 0 && stackNums.Size > 1 {
                  num1, num2 := stackNums.Pop(), stackNums.Pop()
                  oper := stackOper.Pop()
                  if oper == "+" {
                      res := strconv.Atoi(num1) + strconv.Atoi(num2)
                      stackNums.Push(strconv.Itoa(res))
                  }
              }
              stackOper.Push(str) 
            }
            if str == "*" {
                
            }
        }
    }

}

// strconv.Itoa(num) - конвертация числа в строку
// strconv.Atoi(str) - конвертация строки в число