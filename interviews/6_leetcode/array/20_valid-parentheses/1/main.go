package main

// "[]" - true
// "([][])" - true
// "[({})[]]" - true
// "[({[()]})([])]" - true
// "[(" - false
// "}" - false
// "[(})]" - false
// "[}" - false
// "[{]}" - false

// stackOpen, stackClose
// stackOpen [ (
// stackClose }

// Даны строки содержащие в себе скобки 3-х видов, квадратные, круглые и фигурные
// Необходимо написать фукнцию, которая будет проверять на валидность парности скобок

func containsOpenBrackets(currentBracket rune) bool {
    openBrackets := []rune{
        '{', '(', '['
    }
    for _, bracket := range openBrackets {
        if currentBracket == bracket {
            return true
        }
    }
    return false
}

func containsCloseBrackets(currentBracket rune) bool {
   closeBrackets := []rune{
        '}', ')', ']'
    }
   for _, bracket := range closeBrackets {
        if currentBracket == bracket {
            return true
        }
    }
    return false
}

func validateBrackets(brackets) bool {
    curlyBrackets, squareBrackets, roundBrackets := 0, 0, 0
    for _, ch := range brackets {
       if ch == '{' {
           curlyBrackets++
       } else if ch == '}' {
           curlyBrackets--
       }
       if ch == '[' {
           squareBrackets++
       } else if ch == ']' {
           squareBrackets--
       }
       if ch == '(' {
           roundBrackets++
       } else if ch == ')' {
           roundBrackets--
       }
    }
    return curlyBrackets == 0 && squareBrackets == 0 && roundBrackets == 0
}

// {
  "{": "}",
  "(": ")",
  "[": "]"
    
//}

func validateBrackets(brackets) {
    stackOpen := stack{}
    for _, ch := range brackets {
        if ch == '{' || ch == '(' || ch == '[' {
            stackOpen.Push(ch)
        } else {
            prevBrackets := stackOpen.Pop()
            if prevBrackets == '{' && ch != '}' {
                return false 
            } else if prevBrackets == '(' && ch != ')' {
                return false 
            } else if prevBrackets == '[' && ch != ']' {
                return false 
            }
        }
    }
 
    return stackOpen.Len() == 0
}
