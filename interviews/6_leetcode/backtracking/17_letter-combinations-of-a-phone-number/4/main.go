package main

import "fmt"

func bruteForce(index int, s string, curComb []rune, allComb *[]string) {
    if index == len(s) {
        *allComb = append(*allComb, string(curComb))
        return
    }
    digits := map[int]string{
        2: "abc", 3: "def", 4: "ghi",
        5: "jkl", 6: "mno", 7: "pqrs",
        8: "tuv", 9: "wxyz",
    }
    digit := digits[int(s[index])-'0'] // digits[s[index]-'0'] 
    for _, ch := range digit {
        curComb = append(curComb, ch)
        bruteForce(index+1, s, curComb, allComb) // bruteForce(index+1, s, curComb, &res) // allComb
        curComb = curComb[:len(curComb)-1]
    }
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    res := make([]string, 0)
    bruteForce(0, digits, []rune{}, &res)
    return res
}


func main() {
    // Пример использования
    digits := "23"
    combinations := letterCombinations(digits)
    fmt.Println(combinations) // Вывод: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
}