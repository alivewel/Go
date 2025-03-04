package main

func letterCombinations(digits string) []string {
    lettersDigits := map[int]string{
        2: "abc",
        3: "def",
        4: "ghi",
        // ...
    }
    queue := []string{}
    for i, digit := range digits {
        letters := lettersDigits[i]
        newQueue := []string{}
        for _, combination := range queue {
            for _, letter := range letters {
                newQueue = append(newQueue, combination+letter)
            }
        }
        queue = newQueue
    }
    return queue
}

func main() {

}
