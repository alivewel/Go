package main

// stackOpen
// Если ) && len(stackOpen) != 0, то удаляем 
// Если len(stackOpen) == 0 ')' false
// В конце если len(stackOpen) != 0 false

// stackOpen (
// s = "(()()"

// s = "()"

func isBalanced(s string) bool {
    stackOpen := 0
    for _, ch := range s {
        if ch == '(' {
            stackOpen++
        } else {
            if stackOpen == 0 {
                return false
            } else {
                stackOpen--
            }
        }
    }
    if stackOpen != 0 {
        return false
    }
    return true
}