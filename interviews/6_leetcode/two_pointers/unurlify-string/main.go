package main

import "fmt"

// ["a", "b", " ", "c", "#", "#"]
// ["a", "b", " ", "#", "#", "c"]
// ["a", "b", "%", "2", "0", "c"]
// У нас есть быстрый и медленный указатели
// Мы проходим по строке с конца

func unurlify(s []rune) []rune {
    fast, slow := 0, 0 
    for fast < len(s) { // fast >= 0
        if s[fast] != '%' {
            s[slow] = s[fast]
            slow++
        } else {
            s[slow] = ' '
            slow++
            fast += 2
        }
        fast++
    }
    for slow < len(s) {
        s[slow] = '#'
        slow++
    }
    return s
}

func main() {
	s := []rune{'h','e','l','l','o','%','2','0','w','o','r','l','d'} 
	fmt.Println(string(unurlify(s)))
}
