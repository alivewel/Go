package main

import "fmt"

func checkEqualMap(mapVirus, window map[byte]int) bool {
    if len(window) != len(virus) {
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
    if val, found := mapVirus[ch]; found {
        return true 
    }
    return false
}

func checkForVirus(gene, virus string) bool {
    mapVirus := make(map[byte]int)
    for _, val := range virus {
        mapVirus[val]++
    }
    l, r := 0, -1
    window := make(map[byte]int)
    for l < len(gene) {
        // при формировании окна смотреть есть ли элемент в mapVirus
        for r + 1 < len(gene) && r - l + 1 != len(virus) && containsMap(mapVirus, ch) {
            r++ 
            window[gene[r]]++
        }
        sizeWindow := r - l + 1
        if sizeWindow == len(virus) && checkEqualMap(mapVirus, window) {
            return true
        }
        window[gene[l]]--
		// если элементов window[gene[l]] == 0, тоже нужно удалять элементы из окна скорее всего
        if val, _ := window[gene[l]]; val == -1 {
            delete(window, gene[l])
            r++
        }
        l++
    }
    return false
}

func main() {
	gene := "cdeebba"
	virus := "abb"
	fmt.Println(checkForVirus(gene, virus))
}
