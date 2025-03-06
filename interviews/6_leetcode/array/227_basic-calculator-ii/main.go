// s = "3+2*2"
// s = "+ * 3 2 2"
// curRes 
// queueHigh *
// queueLow +
// queueNums 3

// queueNums 3
// queueLow + 
// curRes 4
// s = "3+2*2"

// queueNums 2 3
// queueLow + +
// curRes 4
// s = "2+3+2*2"

// highRes 3
// lowRes 21
// prevOper := byte 
// s = "21+3+2*2+3"
// начинаем с highRes
//
// if highRes 
func calculate(s string) int {
    highRes := 0
    lowRes := 0
    num := ""
    for i := 0; i < len(s) - 1; i++ { 
        if isNum(s[i]) {
            for isNum(s[i + 1]) {
                num += s[i]
                i++
            }

            highRes = strconv.Itoa(num)
        }
        if isOperator(s[i]) {
            if s[i] == '+' || s[i] == '-' {
                lowRes += highRes
                highRes = 0
            } else if s[i] == '*' || s[i] == '/' {
                
            }
        }
        
    }
}