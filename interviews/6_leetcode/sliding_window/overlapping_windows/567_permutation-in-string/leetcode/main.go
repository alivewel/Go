package main

import "fmt"

func checkEqualMap(mapVirus, window map[byte]int) bool {
    if len(window) != len(mapVirus) {
        return false
    }
    for k, v := range mapVirus {
        if val, found := window[k]; found {
            if v != val {
                return false
            }
        }
    }
    return true
}

func containsMap(mapVirus map[byte]int, ch byte) bool {
    if _, found := mapVirus[ch]; found {
        return true 
    }
    return false
}

func checkInclusion(s1, s2 string) bool {
    mapVirus := make(map[byte]int)
    for i := range s2 {
        mapVirus[s2[i]]++
    }
    l, r := 0, -1
    window := make(map[byte]int)
    for l < len(s1) {
        // при формировании окна смотреть есть ли элемент в mapVirus
        for r + 1 < len(s1) && r - l + 1 != len(s2) && containsMap(mapVirus, s1[r+1]) {
            r++ 
            window[s1[r]]++
        }
        sizeWindow := r - l + 1
        if sizeWindow == len(s2) && checkEqualMap(mapVirus, window) {
            return true
        }
        window[s1[l]]--
		// если элементов window[s1[l]] == 0, тоже нужно удалять элементы из окна скорее всего
        if val, _ := window[s1[l]]; val == -1 {
            delete(window, s1[l])
            r++
        }
        l++
    }
    return false
}

func main() {
	s1 := "cdeebba"
	s2 := "abb"
	fmt.Println(checkInclusion(s1, s2))
}
