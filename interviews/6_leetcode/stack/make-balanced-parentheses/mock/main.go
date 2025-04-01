package main
//
// "((H)i)()))"

// stack = 
// removeStack = [5, 6, 8, 10, 11]
// result = ((H))
//             |
// "((H))))i)()))"

// stack = 0 1 
// removeStack = []

// "(((H)"

// stack = // отткрывающиеся
// removeStack = [2 3] // закрыва
//    |
// "(H))) )"

// s = ")Alg(o)Code("
// stack = 11
// removeStack = 0
// if len(stack) > 0  && i == stack[0]  {
//   stack := stack[1:]
// }

// stack = 1 2 3 
// removeStack = 0
// if len(stack) > 0 && i == stack[0]  {
//   stack := stack[1:] / 2 3 
// }

// time O(N)
// mem O(N)

// при первом проходе отмечать открывающиеся скобки
// при втором проходе удалять лишние открывающиеся скобки

// "(((H)"
func makeBalanced(s string) string {
    stack := make([]int, 0)
    removeStack := make([]int, 0)
    for i := range s {
        if s[i] == '(' {
            stack = append(stack, i)
        }
        if s[i] == ')' {
            if len(stack) == 0 {
                removeStack = append(removeStack, i)
            } else {
                stack = stack[:len(stack)-1]
            }
        }
    }
    result := make([]rune, 0, len(s)) // result := make([]rune, len(s))
    for i := range s {
        if len(stack) > 0 && i == stack[0] {
            stack = stack[1:]
            continue
        }
        if len(removeStack) > 0 && i == removeStack[0] {
            removeStack = removeStack[1:] 
            continue
        }
        result = append(result, rune(s[i]))
    }
    return string(result)
}













